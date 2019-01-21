package vm_test

import (
	"encoding/json"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bosh-vmrun-cpi/vm"

	"github.com/cppforlife/bosh-cpi-go/apiv1"
)

var _ = Describe("VMMetadata", func() {
	Describe("NewVMMetadata", func() {
		It("sets the duration fields", func() {
			propsJson := []byte(`{}`)

			var cloudProps apiv1.CloudPropsImpl
			Expect(json.Unmarshal(propsJson, &cloudProps)).To(Succeed())

			vmProps, err := vm.NewVMProps(cloudProps)
			Expect(err).ToNot(HaveOccurred())

			Expect(vmProps.CPU).To(Equal(1))
			Expect(vmProps.RAM).To(Equal(1024))
			Expect(vmProps.Disk).To(Equal(0))
			Expect(vmProps.Bootstrap.Script_Content).To(Equal(""))
			Expect(vmProps.Bootstrap.Script_Path).To(Equal(""))
			Expect(vmProps.Bootstrap.Interpreter_Path).To(Equal(""))
			Expect(vmProps.Bootstrap.Ready_Process_Name).To(Equal(""))
			Expect(vmProps.Bootstrap.Username).To(Equal(""))
			Expect(vmProps.Bootstrap.Password).To(Equal(""))
			Expect(vmProps.Bootstrap.Min_Wait).To(Equal(time.Duration(0)))
			Expect(vmProps.Bootstrap.Max_Wait).To(Equal(600 * time.Second))
		})
	})
})
