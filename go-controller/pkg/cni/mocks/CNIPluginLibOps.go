// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	net "net"

	mock "github.com/stretchr/testify/mock"

	netlink "github.com/vishvananda/netlink"

	ns "github.com/containernetworking/plugins/pkg/ns"
)

// CNIPluginLibOps is an autogenerated mock type for the CNIPluginLibOps type
type CNIPluginLibOps struct {
	mock.Mock
}

// AddRoute provides a mock function with given fields: ipn, gw, dev, mtu
func (_m *CNIPluginLibOps) AddRoute(ipn *net.IPNet, gw net.IP, dev netlink.Link, mtu int) error {
	ret := _m.Called(ipn, gw, dev, mtu)

	if len(ret) == 0 {
		panic("no return value specified for AddRoute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*net.IPNet, net.IP, netlink.Link, int) error); ok {
		r0 = rf(ipn, gw, dev, mtu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetupVeth provides a mock function with given fields: contVethName, hostVethName, mtu, contVethMac, hostNS
func (_m *CNIPluginLibOps) SetupVeth(contVethName string, hostVethName string, mtu int, contVethMac string, hostNS ns.NetNS) (net.Interface, net.Interface, error) {
	ret := _m.Called(contVethName, hostVethName, mtu, contVethMac, hostNS)

	if len(ret) == 0 {
		panic("no return value specified for SetupVeth")
	}

	var r0 net.Interface
	var r1 net.Interface
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string, int, string, ns.NetNS) (net.Interface, net.Interface, error)); ok {
		return rf(contVethName, hostVethName, mtu, contVethMac, hostNS)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, string, ns.NetNS) net.Interface); ok {
		r0 = rf(contVethName, hostVethName, mtu, contVethMac, hostNS)
	} else {
		r0 = ret.Get(0).(net.Interface)
	}

	if rf, ok := ret.Get(1).(func(string, string, int, string, ns.NetNS) net.Interface); ok {
		r1 = rf(contVethName, hostVethName, mtu, contVethMac, hostNS)
	} else {
		r1 = ret.Get(1).(net.Interface)
	}

	if rf, ok := ret.Get(2).(func(string, string, int, string, ns.NetNS) error); ok {
		r2 = rf(contVethName, hostVethName, mtu, contVethMac, hostNS)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewCNIPluginLibOps creates a new instance of CNIPluginLibOps. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCNIPluginLibOps(t interface {
	mock.TestingT
	Cleanup(func())
}) *CNIPluginLibOps {
	mock := &CNIPluginLibOps{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
