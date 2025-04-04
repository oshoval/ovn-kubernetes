// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	config "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/config"
	mock "github.com/stretchr/testify/mock"

	net "net"

	util "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/util"
)

// NetInfo is an autogenerated mock type for the NetInfo type
type NetInfo struct {
	mock.Mock
}

// AllowsPersistentIPs provides a mock function with given fields:
func (_m *NetInfo) AllowsPersistentIPs() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AllowsPersistentIPs")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// EqualNADs provides a mock function with given fields: nads
func (_m *NetInfo) EqualNADs(nads ...string) bool {
	_va := make([]interface{}, len(nads))
	for _i := range nads {
		_va[_i] = nads[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for EqualNADs")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(...string) bool); ok {
		r0 = rf(nads...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ExcludeSubnets provides a mock function with given fields:
func (_m *NetInfo) ExcludeSubnets() []*net.IPNet {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ExcludeSubnets")
	}

	var r0 []*net.IPNet
	if rf, ok := ret.Get(0).(func() []*net.IPNet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*net.IPNet)
		}
	}

	return r0
}

// GetEgressIPAdvertisedNodes provides a mock function with given fields:
func (_m *NetInfo) GetEgressIPAdvertisedNodes() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetEgressIPAdvertisedNodes")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetEgressIPAdvertisedOnNodeVRFs provides a mock function with given fields: node
func (_m *NetInfo) GetEgressIPAdvertisedOnNodeVRFs(node string) []string {
	ret := _m.Called(node)

	if len(ret) == 0 {
		panic("no return value specified for GetEgressIPAdvertisedOnNodeVRFs")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(node)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetEgressIPAdvertisedVRFs provides a mock function with given fields:
func (_m *NetInfo) GetEgressIPAdvertisedVRFs() map[string][]string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetEgressIPAdvertisedVRFs")
	}

	var r0 map[string][]string
	if rf, ok := ret.Get(0).(func() map[string][]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]string)
		}
	}

	return r0
}

// GetNADNamespaces provides a mock function with given fields:
func (_m *NetInfo) GetNADNamespaces() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNADNamespaces")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetNADs provides a mock function with given fields:
func (_m *NetInfo) GetNADs() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNADs")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetNetInfo provides a mock function with given fields:
func (_m *NetInfo) GetNetInfo() util.NetInfo {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNetInfo")
	}

	var r0 util.NetInfo
	if rf, ok := ret.Get(0).(func() util.NetInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(util.NetInfo)
		}
	}

	return r0
}

// GetNetworkID provides a mock function with given fields:
func (_m *NetInfo) GetNetworkID() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkID")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetNetworkName provides a mock function with given fields:
func (_m *NetInfo) GetNetworkName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNetworkScopedClusterRouterName provides a mock function with given fields:
func (_m *NetInfo) GetNetworkScopedClusterRouterName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkScopedClusterRouterName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNetworkScopedClusterSubnetSNATMatch provides a mock function with given fields: nodeName
func (_m *NetInfo) GetNetworkScopedClusterSubnetSNATMatch(nodeName string) string {
	ret := _m.Called(nodeName)

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkScopedClusterSubnetSNATMatch")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(nodeName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNetworkScopedExtPortName provides a mock function with given fields: bridgeID, nodeName
func (_m *NetInfo) GetNetworkScopedExtPortName(bridgeID string, nodeName string) string {
	ret := _m.Called(bridgeID, nodeName)

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkScopedExtPortName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(bridgeID, nodeName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNetworkScopedExtSwitchName provides a mock function with given fields: nodeName
func (_m *NetInfo) GetNetworkScopedExtSwitchName(nodeName string) string {
	ret := _m.Called(nodeName)

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkScopedExtSwitchName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(nodeName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNetworkScopedGWRouterName provides a mock function with given fields: nodeName
func (_m *NetInfo) GetNetworkScopedGWRouterName(nodeName string) string {
	ret := _m.Called(nodeName)

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkScopedGWRouterName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(nodeName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNetworkScopedJoinSwitchName provides a mock function with given fields:
func (_m *NetInfo) GetNetworkScopedJoinSwitchName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkScopedJoinSwitchName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNetworkScopedK8sMgmtIntfName provides a mock function with given fields: nodeName
func (_m *NetInfo) GetNetworkScopedK8sMgmtIntfName(nodeName string) string {
	ret := _m.Called(nodeName)

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkScopedK8sMgmtIntfName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(nodeName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNetworkScopedLoadBalancerGroupName provides a mock function with given fields: lbGroupName
func (_m *NetInfo) GetNetworkScopedLoadBalancerGroupName(lbGroupName string) string {
	ret := _m.Called(lbGroupName)

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkScopedLoadBalancerGroupName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(lbGroupName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNetworkScopedLoadBalancerName provides a mock function with given fields: lbName
func (_m *NetInfo) GetNetworkScopedLoadBalancerName(lbName string) string {
	ret := _m.Called(lbName)

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkScopedLoadBalancerName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(lbName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNetworkScopedName provides a mock function with given fields: name
func (_m *NetInfo) GetNetworkScopedName(name string) string {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkScopedName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNetworkScopedPatchPortName provides a mock function with given fields: bridgeID, nodeName
func (_m *NetInfo) GetNetworkScopedPatchPortName(bridgeID string, nodeName string) string {
	ret := _m.Called(bridgeID, nodeName)

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkScopedPatchPortName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(bridgeID, nodeName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNetworkScopedSwitchName provides a mock function with given fields: nodeName
func (_m *NetInfo) GetNetworkScopedSwitchName(nodeName string) string {
	ret := _m.Called(nodeName)

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkScopedSwitchName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(nodeName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetPodNetworkAdvertisedOnNodeVRFs provides a mock function with given fields: node
func (_m *NetInfo) GetPodNetworkAdvertisedOnNodeVRFs(node string) []string {
	ret := _m.Called(node)

	if len(ret) == 0 {
		panic("no return value specified for GetPodNetworkAdvertisedOnNodeVRFs")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(node)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetPodNetworkAdvertisedVRFs provides a mock function with given fields:
func (_m *NetInfo) GetPodNetworkAdvertisedVRFs() map[string][]string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPodNetworkAdvertisedVRFs")
	}

	var r0 map[string][]string
	if rf, ok := ret.Get(0).(func() map[string][]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]string)
		}
	}

	return r0
}

// HasNAD provides a mock function with given fields: nadName
func (_m *NetInfo) HasNAD(nadName string) bool {
	ret := _m.Called(nadName)

	if len(ret) == 0 {
		panic("no return value specified for HasNAD")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(nadName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IPMode provides a mock function with given fields:
func (_m *NetInfo) IPMode() (bool, bool) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IPMode")
	}

	var r0 bool
	var r1 bool
	if rf, ok := ret.Get(0).(func() (bool, bool)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// IsDefault provides a mock function with given fields:
func (_m *NetInfo) IsDefault() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsDefault")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsPrimaryNetwork provides a mock function with given fields:
func (_m *NetInfo) IsPrimaryNetwork() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsPrimaryNetwork")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsSecondary provides a mock function with given fields:
func (_m *NetInfo) IsSecondary() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsSecondary")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// JoinSubnetV4 provides a mock function with given fields:
func (_m *NetInfo) JoinSubnetV4() *net.IPNet {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for JoinSubnetV4")
	}

	var r0 *net.IPNet
	if rf, ok := ret.Get(0).(func() *net.IPNet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*net.IPNet)
		}
	}

	return r0
}

// JoinSubnetV6 provides a mock function with given fields:
func (_m *NetInfo) JoinSubnetV6() *net.IPNet {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for JoinSubnetV6")
	}

	var r0 *net.IPNet
	if rf, ok := ret.Get(0).(func() *net.IPNet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*net.IPNet)
		}
	}

	return r0
}

// JoinSubnets provides a mock function with given fields:
func (_m *NetInfo) JoinSubnets() []*net.IPNet {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for JoinSubnets")
	}

	var r0 []*net.IPNet
	if rf, ok := ret.Get(0).(func() []*net.IPNet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*net.IPNet)
		}
	}

	return r0
}

// MTU provides a mock function with given fields:
func (_m *NetInfo) MTU() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MTU")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// PhysicalNetworkName provides a mock function with given fields:
func (_m *NetInfo) PhysicalNetworkName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PhysicalNetworkName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RemoveNetworkScopeFromName provides a mock function with given fields: name
func (_m *NetInfo) RemoveNetworkScopeFromName(name string) string {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for RemoveNetworkScopeFromName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Subnets provides a mock function with given fields:
func (_m *NetInfo) Subnets() []config.CIDRNetworkEntry {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Subnets")
	}

	var r0 []config.CIDRNetworkEntry
	if rf, ok := ret.Get(0).(func() []config.CIDRNetworkEntry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]config.CIDRNetworkEntry)
		}
	}

	return r0
}

// TopologyType provides a mock function with given fields:
func (_m *NetInfo) TopologyType() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TopologyType")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Vlan provides a mock function with given fields:
func (_m *NetInfo) Vlan() uint {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Vlan")
	}

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// NewNetInfo creates a new instance of NetInfo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNetInfo(t interface {
	mock.TestingT
	Cleanup(func())
}) *NetInfo {
	mock := &NetInfo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
