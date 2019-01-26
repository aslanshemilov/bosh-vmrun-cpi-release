package action

import (
	"bosh-vmrun-cpi/driver"
	"bosh-vmrun-cpi/vm"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
)

type SetVMMetadataMethod struct {
	client driver.Client
	logger boshlog.Logger
}

func NewSetVMMetadataMethod(client driver.Client, logger boshlog.Logger) SetVMMetadataMethod {
	return SetVMMetadataMethod{client: client, logger: logger}
}

func (m SetVMMetadataMethod) SetVMMetadata(vmCid apiv1.VMCID, apiVMMeta apiv1.VMMeta) error {
	var err error
	vmId := "vm-" + vmCid.AsString()

	m.logger.DebugWithDetails("SetVMMetadata", "metadata:", apiVMMeta)
	vmMeta, err := vm.NewVMMetadata(apiVMMeta)
	if err != nil {
		return err
	}

	vmDisplayName := vmMeta.DisplayName()

	if vmDisplayName != "" {
		m.logger.Debug("SetVMMetadata", "Setting VM display name: ", vmDisplayName)
		err = m.client.StopVM(vmId)
		if err != nil {
			return err
		}

		err = m.client.SetVMDisplayName(vmId, vmDisplayName)
		if err != nil {
			return err
		}

		err = m.client.StartVM(vmId)
		if err != nil {
			return err
		}
	} else {
		m.logger.DebugWithDetails("SetVMMetadata", "Metadata does not contain enough data for display name", vmMeta)
	}

	return nil
}
