package driver

import (
	cpiconfig "bosh-vmrun-cpi/config"
	"time"
)

type ConfigImpl struct {
	cpiConfig cpiconfig.Config
}

func NewConfig(cpiConfig cpiconfig.Config) Config {
	return &ConfigImpl{cpiConfig: cpiConfig}
}

func (c ConfigImpl) OvftoolPath() string {
	return c.cpiConfig.Cloud.Properties.Vmrun.Ovftool_Bin_Path
}

func (c ConfigImpl) VmrunPath() string {
	return c.cpiConfig.Cloud.Properties.Vmrun.Vmrun_Bin_Path
}

func (c ConfigImpl) VdiskmanagerPath() string {
	return c.cpiConfig.Cloud.Properties.Vmrun.Vdiskmanager_Bin_Path
}

func (c ConfigImpl) VmPath() string {
	return c.cpiConfig.Cloud.Properties.Vmrun.Vm_Store_Path
}

func (c ConfigImpl) VmStartMaxWait() time.Duration {
	return c.cpiConfig.Cloud.Properties.Vmrun.Vm_Start_Max_Wait
}

func (c ConfigImpl) VmSoftShutdownMaxWait() time.Duration {
	return c.cpiConfig.Cloud.Properties.Vmrun.Vm_Soft_Shutdown_Max_Wait
}
