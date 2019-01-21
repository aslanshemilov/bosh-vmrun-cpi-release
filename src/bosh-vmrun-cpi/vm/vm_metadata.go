package vm

import (
	"encoding/json"

	"github.com/cppforlife/bosh-cpi-go/apiv1"
)

type VMMetadata struct {
	Director   string
	Deployment string
	Job        string
	Id         string
	Index      string
	Name       string
}

func NewVMMetadata(apiVMMeta apiv1.VMMeta) (*VMMetadata, error) {
	var err error
	data, err := apiVMMeta.MarshalJSON()
	if err != nil {
		return nil, err
	}

	vmMeta := &VMMetadata{}
	err = json.Unmarshal(data, vmMeta)
	if err != nil {
		return nil, err
	}

	return vmMeta, nil
}
