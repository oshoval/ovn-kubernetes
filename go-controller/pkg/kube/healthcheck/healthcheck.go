/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package healthcheck

import (
	"fmt"
	"net"
	"net/http"
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"

	utilerrors "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/util/errors"
)

// Server serves HTTP endpoints for each service name, with results
// based on the endpoints.  If there are 0 endpoints for a service, it returns a
// 503 "Service Unavailable" error (telling LBs not to use this node).  If there
// are 1 or more endpoints, it returns a 200 "OK".
type Server interface {
	// Make the new set of services be active.  Services that were open before
	// will be closed.  Services that are new will be opened.  Service that
	// existed and are in the new set will be left alone.  The value of the map
	// is the healthcheck-port to listen on.
	SyncServices(newServices map[types.NamespacedName]uint16) error
	// Make the new set of endpoints be active.  Endpoints for services that do
	// not exist will be dropped.  The value of the map is the number of
	// endpoints the service has on this node.
	SyncEndpoints(newEndpoints map[types.NamespacedName]int) error
}

// Listener allows for testing of Server.  If the Listener argument
// to NewServer() is nil, the real net.Listen function will be used.
type Listener interface {
	// Listen is very much like net.Listen, except the first arg (network) is
	// fixed to be "tcp".
	Listen(addr string) (net.Listener, error)
}

// HTTPServerFactory allows for testing of Server.  If the
// HTTPServerFactory argument to NewServer() is nil, the real
// http.Server type will be used.
type HTTPServerFactory interface {
	// New creates an instance of a type satisfying HTTPServer.  This is
	// designed to include http.Server.
	New(addr string, handler http.Handler) HTTPServer
}

// HTTPServer allows for testing of Server.
type HTTPServer interface {
	// Server is designed so that http.Server satisfies this interface,
	Serve(listener net.Listener) error
}

// NewServer allocates a new healthcheck server manager.  If either
// of the injected arguments are nil, defaults will be used.
func NewServer(hostname string, recorder record.EventRecorder, listener Listener, httpServerFactory HTTPServerFactory) Server {
	if listener == nil {
		listener = stdNetListener{}
	}
	if httpServerFactory == nil {
		httpServerFactory = stdHTTPServerFactory{}
	}
	return &server{
		hostname:    hostname,
		recorder:    recorder,
		listener:    listener,
		httpFactory: httpServerFactory,
		services:    map[types.NamespacedName]*hcInstance{},
	}
}

// Implement Listener in terms of net.Listen.
type stdNetListener struct{}

func (stdNetListener) Listen(addr string) (net.Listener, error) {
	return net.Listen("tcp", addr)
}

var _ Listener = stdNetListener{}

// Implement HTTPServerFactory in terms of http.Server.
type stdHTTPServerFactory struct{}

func (stdHTTPServerFactory) New(addr string, handler http.Handler) HTTPServer {
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}

var _ HTTPServerFactory = stdHTTPServerFactory{}

type server struct {
	hostname    string
	recorder    record.EventRecorder // can be nil
	listener    Listener
	httpFactory HTTPServerFactory

	lock     sync.RWMutex
	services map[types.NamespacedName]*hcInstance
}

func (hcs *server) SyncServices(newServices map[types.NamespacedName]uint16) error {
	hcs.lock.Lock()
	defer hcs.lock.Unlock()
	var err error
	var errors []error
	// Remove any that are not needed any more.
	for nsn, svc := range hcs.services {
		if port, found := newServices[nsn]; !found || port != svc.port {
			klog.Infof("Closing healthcheck %q on port %d", nsn.String(), svc.port)
			if err = svc.listener.Close(); err != nil {
				errors = append(errors, fmt.Errorf("Close(%v): %v", svc.listener.Addr(), err))
				continue
			}
			delete(hcs.services, nsn)
		}
	}

	// Add any that are needed.
	for nsn, port := range newServices {
		if hcs.services[nsn] != nil {
			klog.V(5).Infof("Existing healthcheck %q on port %d", nsn.String(), port)
			continue
		}

		klog.Infof("Opening healthcheck %q on port %d", nsn.String(), port)
		svc := &hcInstance{port: port}
		addr := fmt.Sprintf(":%d", port)
		svc.server = hcs.httpFactory.New(addr, hcHandler{name: nsn, hcs: hcs})
		svc.listener, err = hcs.listener.Listen(addr)
		if err != nil {
			err := fmt.Errorf("node %s failed to start healthcheck %q on port %d: %v", hcs.hostname, nsn.String(), port, err)

			if hcs.recorder != nil {
				hcs.recorder.Eventf(
					&corev1.ObjectReference{
						Kind:      "Service",
						Namespace: nsn.Namespace,
						Name:      nsn.Name,
						UID:       types.UID(nsn.String()),
					}, corev1.EventTypeWarning, "FailedToStartServiceHealthcheck", err.Error())
			}
			errors = append(errors, err)
			continue
		}
		hcs.services[nsn] = svc

		go func(nsn types.NamespacedName, svc *hcInstance) {
			// Serve() will exit when the listener is closed.
			klog.V(5).Infof("Starting goroutine for healthcheck %q on port %d", nsn.String(), svc.port)
			if err := svc.server.Serve(svc.listener); err != nil {
				klog.V(5).Infof("Healthcheck %q closed: %v", nsn.String(), err)
				return
			}
			klog.V(5).Infof("Healthcheck %q closed", nsn.String())
		}(nsn, svc)
	}
	return utilerrors.Join(errors...)
}

type hcInstance struct {
	port      uint16
	listener  net.Listener
	server    HTTPServer
	endpoints int // number of local endpoints for a service
}

type hcHandler struct {
	name types.NamespacedName
	hcs  *server
}

var _ http.Handler = hcHandler{}

func (h hcHandler) ServeHTTP(resp http.ResponseWriter, _ *http.Request) {
	h.hcs.lock.RLock()
	svc, ok := h.hcs.services[h.name]
	if !ok || svc == nil {
		h.hcs.lock.RUnlock()
		klog.Errorf("Received request for closed healthcheck %q", h.name.String())
		return
	}
	count := svc.endpoints
	h.hcs.lock.RUnlock()

	resp.Header().Set("Content-Type", "application/json")
	if count == 0 {
		resp.WriteHeader(http.StatusServiceUnavailable)
	} else {
		resp.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(resp, `{ "service": { "namespace": %q, "name": %q }, "localEndpoints": %d }`,
		h.name.Namespace, h.name.Name, count)
}

func (hcs *server) SyncEndpoints(newEndpoints map[types.NamespacedName]int) error {
	hcs.lock.Lock()
	defer hcs.lock.Unlock()

	for nsn, count := range newEndpoints {
		if hcs.services[nsn] == nil {
			klog.V(5).Infof("Not saving endpoints for unknown healthcheck %q", nsn.String())
			continue
		}
		klog.V(5).Infof("Reporting %d endpoints for healthcheck %q", count, nsn.String())
		hcs.services[nsn].endpoints = count
	}
	for nsn, hci := range hcs.services {
		if _, found := newEndpoints[nsn]; !found {
			hci.endpoints = 0
		}
	}
	return nil
}
