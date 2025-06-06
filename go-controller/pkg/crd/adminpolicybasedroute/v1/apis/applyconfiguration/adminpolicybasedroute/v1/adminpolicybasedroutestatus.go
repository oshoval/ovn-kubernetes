/*


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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	adminpolicybasedroutev1 "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/adminpolicybasedroute/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AdminPolicyBasedRouteStatusApplyConfiguration represents a declarative configuration of the AdminPolicyBasedRouteStatus type for use
// with apply.
type AdminPolicyBasedRouteStatusApplyConfiguration struct {
	LastTransitionTime *metav1.Time                        `json:"lastTransitionTime,omitempty"`
	Messages           []string                            `json:"messages,omitempty"`
	Status             *adminpolicybasedroutev1.StatusType `json:"status,omitempty"`
}

// AdminPolicyBasedRouteStatusApplyConfiguration constructs a declarative configuration of the AdminPolicyBasedRouteStatus type for use with
// apply.
func AdminPolicyBasedRouteStatus() *AdminPolicyBasedRouteStatusApplyConfiguration {
	return &AdminPolicyBasedRouteStatusApplyConfiguration{}
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *AdminPolicyBasedRouteStatusApplyConfiguration) WithLastTransitionTime(value metav1.Time) *AdminPolicyBasedRouteStatusApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithMessages adds the given value to the Messages field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Messages field.
func (b *AdminPolicyBasedRouteStatusApplyConfiguration) WithMessages(values ...string) *AdminPolicyBasedRouteStatusApplyConfiguration {
	for i := range values {
		b.Messages = append(b.Messages, values[i])
	}
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *AdminPolicyBasedRouteStatusApplyConfiguration) WithStatus(value adminpolicybasedroutev1.StatusType) *AdminPolicyBasedRouteStatusApplyConfiguration {
	b.Status = &value
	return b
}
