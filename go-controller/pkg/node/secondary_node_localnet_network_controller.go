package node

import (
	"context"

	"github.com/ovn-org/ovn-kubernetes/go-controller/pkg/factory"
	"github.com/ovn-org/ovn-kubernetes/go-controller/pkg/node/iprulemanager"
	"github.com/ovn-org/ovn-kubernetes/go-controller/pkg/node/vrfmanager"
	"github.com/ovn-org/ovn-kubernetes/go-controller/pkg/util"
	kapi "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

// SecondaryNodeLocalNetNetworkController structure is the object which holds the controls for starting
// and reacting upon the watched resources, node wise, (e.g. pods) for secondary localnet network
type SecondaryNodeLocalNetNetworkController struct {
	*SecondaryNodeNetworkController
	// pod events factory handler
	podHandlerLocalnet *factory.Handler
	// BaseNodeNetworkController
	// // pod events factory handler
	// podHandler *factory.Handler
	// // stores the networkID of this network
	// networkID *int
	// // responsible for programing gateway elements for this network
	// gateway *UserDefinedNetworkGateway
}

// NewSecondaryNodeLocalNetNetworkController creates a new node controller for reacting upon selected events
// of localnet topology type.
func NewSecondaryNodeLocalNetNetworkController(
	cnnci *CommonNodeNetworkControllerInfo,
	netInfo util.NetInfo,
	vrfManager *vrfmanager.Controller,
	ruleManager *iprulemanager.Controller,
	defaultNetworkGateway Gateway,
) (*SecondaryNodeLocalNetNetworkController, error) {
	controller, err := NewSecondaryNodeNetworkController(cnnci, netInfo, vrfManager, ruleManager, defaultNetworkGateway)
	if err != nil {
		return nil, err
	}

	return &SecondaryNodeLocalNetNetworkController{
		SecondaryNodeNetworkController: controller,
	}, nil
}

// // NewSecondaryNodeNetworkController creates a new OVN controller for creating logical network
// // infrastructure and policy for the given secondary network. It supports layer3, layer2 and
// // localnet topology types.
// func NewSecondaryNodeNetworkController(
// 	cnnci *CommonNodeNetworkControllerInfo,
// 	netInfo util.NetInfo,
// 	vrfManager *vrfmanager.Controller,
// 	ruleManager *iprulemanager.Controller,
// 	defaultNetworkGateway Gateway,
// ) (*SecondaryNodeNetworkController, error) {

// 	snnc := &SecondaryNodeNetworkController{
// 		BaseNodeNetworkController: BaseNodeNetworkController{
// 			CommonNodeNetworkControllerInfo: *cnnci,
// 			ReconcilableNetInfo:             util.NewReconcilableNetInfo(netInfo),
// 			stopChan:                        make(chan struct{}),
// 			wg:                              &sync.WaitGroup{},
// 		},
// 	}
// 	if util.IsNetworkSegmentationSupportEnabled() && snnc.IsPrimaryNetwork() {
// 		node, err := snnc.watchFactory.GetNode(snnc.name)
// 		if err != nil {
// 			return nil, fmt.Errorf("error retrieving node %s while creating node network controller for network %s: %v",
// 				snnc.name, netInfo.GetNetworkName(), err)
// 		}
// 		networkID, err := snnc.getNetworkID()
// 		if err != nil {
// 			return nil, fmt.Errorf("error retrieving network id for network %s: %v", netInfo.GetNetworkName(), err)
// 		}

// 		snnc.gateway, err = NewUserDefinedNetworkGateway(snnc.GetNetInfo(), networkID, node,
// 			snnc.watchFactory.NodeCoreInformer().Lister(), snnc.Kube, vrfManager, ruleManager, defaultNetworkGateway)
// 		if err != nil {
// 			return nil, fmt.Errorf("error creating UDN gateway for network %s: %v", netInfo.GetNetworkName(), err)
// 		}
// 	}
// 	return snnc, nil
// }

// // Start starts the default controller; handles all events and creates all needed logical entities
func (nc *SecondaryNodeLocalNetNetworkController) Start(ctx context.Context) error {
	klog.Infof("Start secondary node network localnet controller of network %s", nc.GetNetworkName())

	err := nc.SecondaryNodeNetworkController.Start(ctx)
	if err != nil {
		return err
	}

	handler, err := nc.watchPods()
	if err != nil {
		return err
	}
	nc.podHandlerLocalnet = handler

	return nil
}

// Stop gracefully stops the controller
func (nc *SecondaryNodeLocalNetNetworkController) Stop() {
	klog.Infof("Stop secondary node localnet network controller of network %s", nc.GetNetworkName())

	if nc.podHandlerLocalnet != nil {
		nc.watchFactory.RemovePodHandler(nc.podHandlerLocalnet)
	}

	nc.SecondaryNodeNetworkController.Stop()
}

// watchPods watch add / updates for pods
func (nc *SecondaryNodeLocalNetNetworkController) watchPods() (*factory.Handler, error) {
	//clientSet := cni.NewClientSet(nc.client, corev1listers.NewPodLister(nc.watchFactory.LocalPodInformer().GetIndexer()))

	//netName := nc.GetNetworkName()
	return nc.watchFactory.AddPodHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(*kapi.Pod)
			if !util.PodScheduled(pod) {
				return
			}

			if nc.name != pod.Spec.NodeName {
				return
			}

			// TODO
			// given the pod / pod NSE, need to determine if the pod use a network related to the netName this controller belongs to
			// maybe for that we need nadInformer ? read the nads of the NSE, to get their netName
			// can NSE dont have namespace, if so it takes ns from the pod ?

			allNetworks, err := util.GetK8sPodAllNetworkSelections(pod)
			if err != nil {
				return
			}

			nInfo := nc.GetNetInfo()
			nadFound := false
			for _, network := range allNetworks {
				nadName := util.GetNADName(network.Namespace, network.Name)
				if nInfo.HasNAD(nadName) {
					nadFound = true
					break
				}
			}

			if !nadFound {
				return
			}

			//klog.V(5).Infof("Add for Pod: %s/%s for network %s", pod.Namespace, pod.Name, netName)
			klog.Infof("DBG at SecondaryNodeLocalNetNetworkController AddFunc, pod %s", pod.Name)

			// if util.PodWantsHostNetwork(pod) {
			// 	return
			// }

			// // add all the Pod's NADs into Pod's nadToDPUCDMap
			// // For default network, NAD name is DefaultNetworkName.
			// nadToDPUCDMap := map[string]*util.DPUConnectionDetails{}
			// if bnnc.IsSecondary() {
			// 	on, networkMap, err := util.GetPodNADToNetworkMapping(pod, bnnc.GetNetInfo())
			// 	if err != nil || !on {
			// 		if err != nil {
			// 			// configuration error, no need to retry, do not return error
			// 			klog.Errorf("Error getting network-attachment for pod %s/%s network %s: %v",
			// 				pod.Namespace, pod.Name, bnnc.GetNetworkName(), err)
			// 		} else {
			// 			klog.V(5).Infof("Skipping Pod %s/%s as it is not attached to network: %s",
			// 				pod.Namespace, pod.Name, netName)
			// 		}
			// 		return
			// 	}
			// 	for nadName := range networkMap {
			// 		nadToDPUCDMap[nadName] = nil
			// 	}
			// } else {
			// 	nadToDPUCDMap[types.DefaultNetworkName] = nil
			// }

			// for nadName := range nadToDPUCDMap {
			// 	dpuCD := bnnc.podReadyToAddDPU(pod, nadName)
			// 	if dpuCD != nil {
			// 		err := bnnc.addDPUPodForNAD(pod, dpuCD, netName, nadName, clientSet)
			// 		if err != nil {
			// 			klog.Errorf(err.Error())
			// 		} else {
			// 			nadToDPUCDMap[nadName] = dpuCD
			// 		}
			// 	}
			// }
			// bnnc.podNADToDPUCDMap.Store(pod.UID, nadToDPUCDMap)
		},
		UpdateFunc: func(old, newer interface{}) {
			// TODO need to do same sanity check of Add
			klog.Infof("DBG at SecondaryNodeLocalNetNetworkController UpdateFunc")

			// oldPod := old.(*kapi.Pod)
			// newPod := newer.(*kapi.Pod)
			// klog.V(5).Infof("Update for Pod: %s/%s for network %s", newPod.Namespace, newPod.Name, netName)
			// v, ok := bnnc.podNADToDPUCDMap.Load(newPod.UID)
			// if !ok {
			// 	klog.V(5).Infof("Skipping update for Pod %s/%s as it is not attached to network: %s",
			// 		newPod.Namespace, newPod.Name, netName)
			// 	return
			// }
			// nadToDPUCDMap := v.(map[string]*util.DPUConnectionDetails)
			// for nadName := range nadToDPUCDMap {
			// 	oldDPUCD := nadToDPUCDMap[nadName]
			// 	newDPUCD := bnnc.podReadyToAddDPU(newPod, nadName)
			// 	if !dpuConnectionDetailChanged(oldDPUCD, newDPUCD) {
			// 		continue
			// 	}
			// 	if oldDPUCD != nil {
			// 		// VF already added, but new Pod has changed, we'd need to delete the old VF
			// 		klog.Infof("Deleting the old VF since either kubelet issued cmdDEL or assigned a new VF or "+
			// 			"the sandbox id itself changed. Old connection details (%v), New connection details (%v)",
			// 			oldDPUCD, newDPUCD)
			// 		err := bnnc.delDPUPodForNAD(oldPod, oldDPUCD, nadName, false)
			// 		if err != nil {
			// 			klog.Errorf(err.Error())
			// 		}
			// 		nadToDPUCDMap[nadName] = nil
			// 	}
			// 	if newDPUCD != nil {
			// 		klog.Infof("Adding VF during update because either during Pod Add we failed to add VF or "+
			// 			"connection details weren't present or the VF ID has changed. Old connection details (%v), "+
			// 			"New connection details (%v)", oldDPUCD, newDPUCD)
			// 		err := bnnc.addDPUPodForNAD(newPod, newDPUCD, netName, nadName, clientSet)
			// 		if err != nil {
			// 			klog.Errorf(err.Error())
			// 		} else {
			// 			nadToDPUCDMap[nadName] = newDPUCD
			// 		}
			// 	}
			// }
			// bnnc.podNADToDPUCDMap.Store(newPod.UID, nadToDPUCDMap)
		},
	}, nil)
}

func (nc *SecondaryNodeLocalNetNetworkController) getNetworkName() string {
	netInfo := nc.GetNetInfo()
	if netInfo.PhysicalNetworkName() != "" {
		return netInfo.PhysicalNetworkName()
	}
	return netInfo.GetNetworkName()
}

// 	klog.Infof("Start secondary node network controller of network %s", nc.GetNetworkName())

// 	// enable adding ovs ports for dpu pods in both primary and secondary user defined networks
// 	if (config.OVNKubernetesFeature.EnableMultiNetwork || util.IsNetworkSegmentationSupportEnabled()) && config.OvnKubeNode.Mode == types.NodeModeDPU {
// 		handler, err := nc.watchPodsDPU()
// 		if err != nil {
// 			return err
// 		}
// 		nc.podHandler = handler
// 	}
// 	if util.IsNetworkSegmentationSupportEnabled() && nc.IsPrimaryNetwork() {
// 		if err := nc.gateway.AddNetwork(); err != nil {
// 			return fmt.Errorf("failed to add network to node gateway for network %s at node %s: %w",
// 				nc.GetNetworkName(), nc.name, err)
// 		}
// 	}
// 	return nil
// }

// // Stop gracefully stops the controller
// func (nc *SecondaryNodeNetworkController) Stop() {
// 	klog.Infof("Stop secondary node network controller of network %s", nc.GetNetworkName())
// 	close(nc.stopChan)
// 	nc.wg.Wait()

// 	if nc.podHandler != nil {
// 		nc.watchFactory.RemovePodHandler(nc.podHandler)
// 	}
// }

// // Cleanup cleans up node entities for the given secondary network
// func (nc *SecondaryNodeNetworkController) Cleanup() error {
// 	if nc.gateway != nil {
// 		return nc.gateway.DelNetwork()
// 	}
// 	return nil
// }

// func (oc *SecondaryNodeNetworkController) getNetworkID() (int, error) {
// 	if oc.networkID == nil || *oc.networkID == util.InvalidID {
// 		oc.networkID = ptr.To(util.InvalidID)
// 		if netID := oc.GetNetworkID(); netID != util.InvalidID {
// 			*oc.networkID = netID
// 			return *oc.networkID, nil
// 		}

// 		nodes, err := oc.watchFactory.GetNodes()
// 		if err != nil {
// 			return util.InvalidID, err
// 		}
// 		*oc.networkID, err = util.GetNetworkID(nodes, oc.GetNetInfo())
// 		if err != nil {
// 			return util.InvalidID, err
// 		}
// 	}
// 	return *oc.networkID, nil
// }

// func (oc *SecondaryNodeNetworkController) Reconcile(netInfo util.NetInfo) error {
// 	// reconcile network information, point of no return
// 	err := util.ReconcileNetInfo(oc.ReconcilableNetInfo, netInfo)
// 	if err != nil {
// 		klog.Errorf("Failed to reconcile network %s: %v", oc.GetNetworkName(), err)
// 	}
// 	return nil
// }
