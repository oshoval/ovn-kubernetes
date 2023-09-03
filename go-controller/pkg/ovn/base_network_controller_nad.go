package ovn

import (
	"fmt"

	nadapi "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	"github.com/ovn-org/ovn-kubernetes/go-controller/pkg/types"
	ovntypes "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/types"
	"github.com/ovn-org/ovn-kubernetes/go-controller/pkg/util"
	"k8s.io/klog/v2"
)

func (bsnc *BaseSecondaryNetworkController) syncNetworkAttachments(networkAttachments []interface{}) error {
	for _, nadInterface := range networkAttachments {
		nad, ok := nadInterface.(*nadapi.NetworkAttachmentDefinition)
		if !ok {
			return fmt.Errorf("spurious object in syncNetworkAttachments: %v", nadInterface)
		}
		klog.Infof("DBG syncNetworkAttachments %s/%s", nad.Namespace, nad.Name)
		bsnc.addNetworkAttachmentDefinition(nad)
	}

	return nil
}

func (bsnc *BaseSecondaryNetworkController) addNetworkAttachmentDefinition(nad *nadapi.NetworkAttachmentDefinition) {
	if validateNetworkAttachmentDefinition(nad) != nil {
		return
	}

	klog.Infof("DBG addNetworkAttachmentDefinition %s/%s", nad.Namespace, nad.Name)
	nadName := fmt.Sprintf("%s/%s", nad.Namespace, nad.Name)
	bsnc.AddNAD(nadName)
	bsnc.NetInfo.AddNAD(nadName)
}

func (bsnc *BaseSecondaryNetworkController) deleteNetworkAttachmentDefinition(nad *nadapi.NetworkAttachmentDefinition) {
	if validateNetworkAttachmentDefinition(nad) != nil {
		return
	}

	klog.Infof("DBG deleteNetworkAttachmentDefinition %s/%s", nad.Namespace, nad.Name)
	nadName := fmt.Sprintf("%s/%s", nad.Namespace, nad.Name)
	bsnc.DeleteNAD(nadName)
	bsnc.NetInfo.DeleteNAD(nadName)
}

func validateNetworkAttachmentDefinition(nad *nadapi.NetworkAttachmentDefinition) error {
	nInfo, err := util.ParseNADInfo(nad)
	if err != nil {
		return err
	}

	if nInfo.GetNetworkName() == types.DefaultNetworkName {
		return fmt.Errorf("NAD for default network, skip it")
	}

	topoType := nInfo.TopologyType()
	switch topoType {
	case ovntypes.Layer3Topology:
	case ovntypes.Layer2Topology:
	case ovntypes.LocalnetTopology:
		break
	default:
		return fmt.Errorf("topology type %s not supported", topoType)
	}

	return nil
}
