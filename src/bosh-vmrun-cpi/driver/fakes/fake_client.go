// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"bosh-vmrun-cpi/driver"
	"sync"
	"time"
)

type FakeClient struct {
	AttachDiskStub        func(string, string) error
	attachDiskMutex       sync.RWMutex
	attachDiskArgsForCall []struct {
		arg1 string
		arg2 string
	}
	attachDiskReturns struct {
		result1 error
	}
	attachDiskReturnsOnCall map[int]struct {
		result1 error
	}
	BootstrapVMStub        func(string, string, string, string, string, string, string, time.Duration, time.Duration) error
	bootstrapVMMutex       sync.RWMutex
	bootstrapVMArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
		arg6 string
		arg7 string
		arg8 time.Duration
		arg9 time.Duration
	}
	bootstrapVMReturns struct {
		result1 error
	}
	bootstrapVMReturnsOnCall map[int]struct {
		result1 error
	}
	CloneVMStub        func(string, string) error
	cloneVMMutex       sync.RWMutex
	cloneVMArgsForCall []struct {
		arg1 string
		arg2 string
	}
	cloneVMReturns struct {
		result1 error
	}
	cloneVMReturnsOnCall map[int]struct {
		result1 error
	}
	CreateDiskStub        func(string, int) error
	createDiskMutex       sync.RWMutex
	createDiskArgsForCall []struct {
		arg1 string
		arg2 int
	}
	createDiskReturns struct {
		result1 error
	}
	createDiskReturnsOnCall map[int]struct {
		result1 error
	}
	CreateEphemeralDiskStub        func(string, int) error
	createEphemeralDiskMutex       sync.RWMutex
	createEphemeralDiskArgsForCall []struct {
		arg1 string
		arg2 int
	}
	createEphemeralDiskReturns struct {
		result1 error
	}
	createEphemeralDiskReturnsOnCall map[int]struct {
		result1 error
	}
	DestroyDiskStub        func(string) error
	destroyDiskMutex       sync.RWMutex
	destroyDiskArgsForCall []struct {
		arg1 string
	}
	destroyDiskReturns struct {
		result1 error
	}
	destroyDiskReturnsOnCall map[int]struct {
		result1 error
	}
	DestroyVMStub        func(string) error
	destroyVMMutex       sync.RWMutex
	destroyVMArgsForCall []struct {
		arg1 string
	}
	destroyVMReturns struct {
		result1 error
	}
	destroyVMReturnsOnCall map[int]struct {
		result1 error
	}
	DetachDiskStub        func(string, string) error
	detachDiskMutex       sync.RWMutex
	detachDiskArgsForCall []struct {
		arg1 string
		arg2 string
	}
	detachDiskReturns struct {
		result1 error
	}
	detachDiskReturnsOnCall map[int]struct {
		result1 error
	}
	GetVMInfoStub        func(string) (driver.VMInfo, error)
	getVMInfoMutex       sync.RWMutex
	getVMInfoArgsForCall []struct {
		arg1 string
	}
	getVMInfoReturns struct {
		result1 driver.VMInfo
		result2 error
	}
	getVMInfoReturnsOnCall map[int]struct {
		result1 driver.VMInfo
		result2 error
	}
	GetVMIsoPathStub        func(string) string
	getVMIsoPathMutex       sync.RWMutex
	getVMIsoPathArgsForCall []struct {
		arg1 string
	}
	getVMIsoPathReturns struct {
		result1 string
	}
	getVMIsoPathReturnsOnCall map[int]struct {
		result1 string
	}
	HasDiskStub        func(string) bool
	hasDiskMutex       sync.RWMutex
	hasDiskArgsForCall []struct {
		arg1 string
	}
	hasDiskReturns struct {
		result1 bool
	}
	hasDiskReturnsOnCall map[int]struct {
		result1 bool
	}
	HasVMStub        func(string) bool
	hasVMMutex       sync.RWMutex
	hasVMArgsForCall []struct {
		arg1 string
	}
	hasVMReturns struct {
		result1 bool
	}
	hasVMReturnsOnCall map[int]struct {
		result1 bool
	}
	ImportOvfStub        func(string, string) (bool, error)
	importOvfMutex       sync.RWMutex
	importOvfArgsForCall []struct {
		arg1 string
		arg2 string
	}
	importOvfReturns struct {
		result1 bool
		result2 error
	}
	importOvfReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	SetVMNetworkAdapterStub        func(string, string, string) error
	setVMNetworkAdapterMutex       sync.RWMutex
	setVMNetworkAdapterArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	setVMNetworkAdapterReturns struct {
		result1 error
	}
	setVMNetworkAdapterReturnsOnCall map[int]struct {
		result1 error
	}
	SetVMResourcesStub        func(string, int, int) error
	setVMResourcesMutex       sync.RWMutex
	setVMResourcesArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 int
	}
	setVMResourcesReturns struct {
		result1 error
	}
	setVMResourcesReturnsOnCall map[int]struct {
		result1 error
	}
	StartVMStub        func(string) error
	startVMMutex       sync.RWMutex
	startVMArgsForCall []struct {
		arg1 string
	}
	startVMReturns struct {
		result1 error
	}
	startVMReturnsOnCall map[int]struct {
		result1 error
	}
	StopVMStub        func(string) error
	stopVMMutex       sync.RWMutex
	stopVMArgsForCall []struct {
		arg1 string
	}
	stopVMReturns struct {
		result1 error
	}
	stopVMReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateVMIsoStub        func(string, string) error
	updateVMIsoMutex       sync.RWMutex
	updateVMIsoArgsForCall []struct {
		arg1 string
		arg2 string
	}
	updateVMIsoReturns struct {
		result1 error
	}
	updateVMIsoReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) AttachDisk(arg1 string, arg2 string) error {
	fake.attachDiskMutex.Lock()
	ret, specificReturn := fake.attachDiskReturnsOnCall[len(fake.attachDiskArgsForCall)]
	fake.attachDiskArgsForCall = append(fake.attachDiskArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("AttachDisk", []interface{}{arg1, arg2})
	fake.attachDiskMutex.Unlock()
	if fake.AttachDiskStub != nil {
		return fake.AttachDiskStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.attachDiskReturns
	return fakeReturns.result1
}

func (fake *FakeClient) AttachDiskCallCount() int {
	fake.attachDiskMutex.RLock()
	defer fake.attachDiskMutex.RUnlock()
	return len(fake.attachDiskArgsForCall)
}

func (fake *FakeClient) AttachDiskCalls(stub func(string, string) error) {
	fake.attachDiskMutex.Lock()
	defer fake.attachDiskMutex.Unlock()
	fake.AttachDiskStub = stub
}

func (fake *FakeClient) AttachDiskArgsForCall(i int) (string, string) {
	fake.attachDiskMutex.RLock()
	defer fake.attachDiskMutex.RUnlock()
	argsForCall := fake.attachDiskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) AttachDiskReturns(result1 error) {
	fake.attachDiskMutex.Lock()
	defer fake.attachDiskMutex.Unlock()
	fake.AttachDiskStub = nil
	fake.attachDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) AttachDiskReturnsOnCall(i int, result1 error) {
	fake.attachDiskMutex.Lock()
	defer fake.attachDiskMutex.Unlock()
	fake.AttachDiskStub = nil
	if fake.attachDiskReturnsOnCall == nil {
		fake.attachDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.attachDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) BootstrapVM(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 string, arg7 string, arg8 time.Duration, arg9 time.Duration) error {
	fake.bootstrapVMMutex.Lock()
	ret, specificReturn := fake.bootstrapVMReturnsOnCall[len(fake.bootstrapVMArgsForCall)]
	fake.bootstrapVMArgsForCall = append(fake.bootstrapVMArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
		arg6 string
		arg7 string
		arg8 time.Duration
		arg9 time.Duration
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9})
	fake.recordInvocation("BootstrapVM", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9})
	fake.bootstrapVMMutex.Unlock()
	if fake.BootstrapVMStub != nil {
		return fake.BootstrapVMStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.bootstrapVMReturns
	return fakeReturns.result1
}

func (fake *FakeClient) BootstrapVMCallCount() int {
	fake.bootstrapVMMutex.RLock()
	defer fake.bootstrapVMMutex.RUnlock()
	return len(fake.bootstrapVMArgsForCall)
}

func (fake *FakeClient) BootstrapVMCalls(stub func(string, string, string, string, string, string, string, time.Duration, time.Duration) error) {
	fake.bootstrapVMMutex.Lock()
	defer fake.bootstrapVMMutex.Unlock()
	fake.BootstrapVMStub = stub
}

func (fake *FakeClient) BootstrapVMArgsForCall(i int) (string, string, string, string, string, string, string, time.Duration, time.Duration) {
	fake.bootstrapVMMutex.RLock()
	defer fake.bootstrapVMMutex.RUnlock()
	argsForCall := fake.bootstrapVMArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7, argsForCall.arg8, argsForCall.arg9
}

func (fake *FakeClient) BootstrapVMReturns(result1 error) {
	fake.bootstrapVMMutex.Lock()
	defer fake.bootstrapVMMutex.Unlock()
	fake.BootstrapVMStub = nil
	fake.bootstrapVMReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) BootstrapVMReturnsOnCall(i int, result1 error) {
	fake.bootstrapVMMutex.Lock()
	defer fake.bootstrapVMMutex.Unlock()
	fake.BootstrapVMStub = nil
	if fake.bootstrapVMReturnsOnCall == nil {
		fake.bootstrapVMReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.bootstrapVMReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CloneVM(arg1 string, arg2 string) error {
	fake.cloneVMMutex.Lock()
	ret, specificReturn := fake.cloneVMReturnsOnCall[len(fake.cloneVMArgsForCall)]
	fake.cloneVMArgsForCall = append(fake.cloneVMArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CloneVM", []interface{}{arg1, arg2})
	fake.cloneVMMutex.Unlock()
	if fake.CloneVMStub != nil {
		return fake.CloneVMStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cloneVMReturns
	return fakeReturns.result1
}

func (fake *FakeClient) CloneVMCallCount() int {
	fake.cloneVMMutex.RLock()
	defer fake.cloneVMMutex.RUnlock()
	return len(fake.cloneVMArgsForCall)
}

func (fake *FakeClient) CloneVMCalls(stub func(string, string) error) {
	fake.cloneVMMutex.Lock()
	defer fake.cloneVMMutex.Unlock()
	fake.CloneVMStub = stub
}

func (fake *FakeClient) CloneVMArgsForCall(i int) (string, string) {
	fake.cloneVMMutex.RLock()
	defer fake.cloneVMMutex.RUnlock()
	argsForCall := fake.cloneVMArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) CloneVMReturns(result1 error) {
	fake.cloneVMMutex.Lock()
	defer fake.cloneVMMutex.Unlock()
	fake.CloneVMStub = nil
	fake.cloneVMReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CloneVMReturnsOnCall(i int, result1 error) {
	fake.cloneVMMutex.Lock()
	defer fake.cloneVMMutex.Unlock()
	fake.CloneVMStub = nil
	if fake.cloneVMReturnsOnCall == nil {
		fake.cloneVMReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cloneVMReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CreateDisk(arg1 string, arg2 int) error {
	fake.createDiskMutex.Lock()
	ret, specificReturn := fake.createDiskReturnsOnCall[len(fake.createDiskArgsForCall)]
	fake.createDiskArgsForCall = append(fake.createDiskArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("CreateDisk", []interface{}{arg1, arg2})
	fake.createDiskMutex.Unlock()
	if fake.CreateDiskStub != nil {
		return fake.CreateDiskStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createDiskReturns
	return fakeReturns.result1
}

func (fake *FakeClient) CreateDiskCallCount() int {
	fake.createDiskMutex.RLock()
	defer fake.createDiskMutex.RUnlock()
	return len(fake.createDiskArgsForCall)
}

func (fake *FakeClient) CreateDiskCalls(stub func(string, int) error) {
	fake.createDiskMutex.Lock()
	defer fake.createDiskMutex.Unlock()
	fake.CreateDiskStub = stub
}

func (fake *FakeClient) CreateDiskArgsForCall(i int) (string, int) {
	fake.createDiskMutex.RLock()
	defer fake.createDiskMutex.RUnlock()
	argsForCall := fake.createDiskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) CreateDiskReturns(result1 error) {
	fake.createDiskMutex.Lock()
	defer fake.createDiskMutex.Unlock()
	fake.CreateDiskStub = nil
	fake.createDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CreateDiskReturnsOnCall(i int, result1 error) {
	fake.createDiskMutex.Lock()
	defer fake.createDiskMutex.Unlock()
	fake.CreateDiskStub = nil
	if fake.createDiskReturnsOnCall == nil {
		fake.createDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CreateEphemeralDisk(arg1 string, arg2 int) error {
	fake.createEphemeralDiskMutex.Lock()
	ret, specificReturn := fake.createEphemeralDiskReturnsOnCall[len(fake.createEphemeralDiskArgsForCall)]
	fake.createEphemeralDiskArgsForCall = append(fake.createEphemeralDiskArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("CreateEphemeralDisk", []interface{}{arg1, arg2})
	fake.createEphemeralDiskMutex.Unlock()
	if fake.CreateEphemeralDiskStub != nil {
		return fake.CreateEphemeralDiskStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createEphemeralDiskReturns
	return fakeReturns.result1
}

func (fake *FakeClient) CreateEphemeralDiskCallCount() int {
	fake.createEphemeralDiskMutex.RLock()
	defer fake.createEphemeralDiskMutex.RUnlock()
	return len(fake.createEphemeralDiskArgsForCall)
}

func (fake *FakeClient) CreateEphemeralDiskCalls(stub func(string, int) error) {
	fake.createEphemeralDiskMutex.Lock()
	defer fake.createEphemeralDiskMutex.Unlock()
	fake.CreateEphemeralDiskStub = stub
}

func (fake *FakeClient) CreateEphemeralDiskArgsForCall(i int) (string, int) {
	fake.createEphemeralDiskMutex.RLock()
	defer fake.createEphemeralDiskMutex.RUnlock()
	argsForCall := fake.createEphemeralDiskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) CreateEphemeralDiskReturns(result1 error) {
	fake.createEphemeralDiskMutex.Lock()
	defer fake.createEphemeralDiskMutex.Unlock()
	fake.CreateEphemeralDiskStub = nil
	fake.createEphemeralDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CreateEphemeralDiskReturnsOnCall(i int, result1 error) {
	fake.createEphemeralDiskMutex.Lock()
	defer fake.createEphemeralDiskMutex.Unlock()
	fake.CreateEphemeralDiskStub = nil
	if fake.createEphemeralDiskReturnsOnCall == nil {
		fake.createEphemeralDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createEphemeralDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DestroyDisk(arg1 string) error {
	fake.destroyDiskMutex.Lock()
	ret, specificReturn := fake.destroyDiskReturnsOnCall[len(fake.destroyDiskArgsForCall)]
	fake.destroyDiskArgsForCall = append(fake.destroyDiskArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DestroyDisk", []interface{}{arg1})
	fake.destroyDiskMutex.Unlock()
	if fake.DestroyDiskStub != nil {
		return fake.DestroyDiskStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.destroyDiskReturns
	return fakeReturns.result1
}

func (fake *FakeClient) DestroyDiskCallCount() int {
	fake.destroyDiskMutex.RLock()
	defer fake.destroyDiskMutex.RUnlock()
	return len(fake.destroyDiskArgsForCall)
}

func (fake *FakeClient) DestroyDiskCalls(stub func(string) error) {
	fake.destroyDiskMutex.Lock()
	defer fake.destroyDiskMutex.Unlock()
	fake.DestroyDiskStub = stub
}

func (fake *FakeClient) DestroyDiskArgsForCall(i int) string {
	fake.destroyDiskMutex.RLock()
	defer fake.destroyDiskMutex.RUnlock()
	argsForCall := fake.destroyDiskArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) DestroyDiskReturns(result1 error) {
	fake.destroyDiskMutex.Lock()
	defer fake.destroyDiskMutex.Unlock()
	fake.DestroyDiskStub = nil
	fake.destroyDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DestroyDiskReturnsOnCall(i int, result1 error) {
	fake.destroyDiskMutex.Lock()
	defer fake.destroyDiskMutex.Unlock()
	fake.DestroyDiskStub = nil
	if fake.destroyDiskReturnsOnCall == nil {
		fake.destroyDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DestroyVM(arg1 string) error {
	fake.destroyVMMutex.Lock()
	ret, specificReturn := fake.destroyVMReturnsOnCall[len(fake.destroyVMArgsForCall)]
	fake.destroyVMArgsForCall = append(fake.destroyVMArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DestroyVM", []interface{}{arg1})
	fake.destroyVMMutex.Unlock()
	if fake.DestroyVMStub != nil {
		return fake.DestroyVMStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.destroyVMReturns
	return fakeReturns.result1
}

func (fake *FakeClient) DestroyVMCallCount() int {
	fake.destroyVMMutex.RLock()
	defer fake.destroyVMMutex.RUnlock()
	return len(fake.destroyVMArgsForCall)
}

func (fake *FakeClient) DestroyVMCalls(stub func(string) error) {
	fake.destroyVMMutex.Lock()
	defer fake.destroyVMMutex.Unlock()
	fake.DestroyVMStub = stub
}

func (fake *FakeClient) DestroyVMArgsForCall(i int) string {
	fake.destroyVMMutex.RLock()
	defer fake.destroyVMMutex.RUnlock()
	argsForCall := fake.destroyVMArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) DestroyVMReturns(result1 error) {
	fake.destroyVMMutex.Lock()
	defer fake.destroyVMMutex.Unlock()
	fake.DestroyVMStub = nil
	fake.destroyVMReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DestroyVMReturnsOnCall(i int, result1 error) {
	fake.destroyVMMutex.Lock()
	defer fake.destroyVMMutex.Unlock()
	fake.DestroyVMStub = nil
	if fake.destroyVMReturnsOnCall == nil {
		fake.destroyVMReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyVMReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DetachDisk(arg1 string, arg2 string) error {
	fake.detachDiskMutex.Lock()
	ret, specificReturn := fake.detachDiskReturnsOnCall[len(fake.detachDiskArgsForCall)]
	fake.detachDiskArgsForCall = append(fake.detachDiskArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DetachDisk", []interface{}{arg1, arg2})
	fake.detachDiskMutex.Unlock()
	if fake.DetachDiskStub != nil {
		return fake.DetachDiskStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.detachDiskReturns
	return fakeReturns.result1
}

func (fake *FakeClient) DetachDiskCallCount() int {
	fake.detachDiskMutex.RLock()
	defer fake.detachDiskMutex.RUnlock()
	return len(fake.detachDiskArgsForCall)
}

func (fake *FakeClient) DetachDiskCalls(stub func(string, string) error) {
	fake.detachDiskMutex.Lock()
	defer fake.detachDiskMutex.Unlock()
	fake.DetachDiskStub = stub
}

func (fake *FakeClient) DetachDiskArgsForCall(i int) (string, string) {
	fake.detachDiskMutex.RLock()
	defer fake.detachDiskMutex.RUnlock()
	argsForCall := fake.detachDiskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) DetachDiskReturns(result1 error) {
	fake.detachDiskMutex.Lock()
	defer fake.detachDiskMutex.Unlock()
	fake.DetachDiskStub = nil
	fake.detachDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DetachDiskReturnsOnCall(i int, result1 error) {
	fake.detachDiskMutex.Lock()
	defer fake.detachDiskMutex.Unlock()
	fake.DetachDiskStub = nil
	if fake.detachDiskReturnsOnCall == nil {
		fake.detachDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.detachDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) GetVMInfo(arg1 string) (driver.VMInfo, error) {
	fake.getVMInfoMutex.Lock()
	ret, specificReturn := fake.getVMInfoReturnsOnCall[len(fake.getVMInfoArgsForCall)]
	fake.getVMInfoArgsForCall = append(fake.getVMInfoArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetVMInfo", []interface{}{arg1})
	fake.getVMInfoMutex.Unlock()
	if fake.GetVMInfoStub != nil {
		return fake.GetVMInfoStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getVMInfoReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) GetVMInfoCallCount() int {
	fake.getVMInfoMutex.RLock()
	defer fake.getVMInfoMutex.RUnlock()
	return len(fake.getVMInfoArgsForCall)
}

func (fake *FakeClient) GetVMInfoCalls(stub func(string) (driver.VMInfo, error)) {
	fake.getVMInfoMutex.Lock()
	defer fake.getVMInfoMutex.Unlock()
	fake.GetVMInfoStub = stub
}

func (fake *FakeClient) GetVMInfoArgsForCall(i int) string {
	fake.getVMInfoMutex.RLock()
	defer fake.getVMInfoMutex.RUnlock()
	argsForCall := fake.getVMInfoArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) GetVMInfoReturns(result1 driver.VMInfo, result2 error) {
	fake.getVMInfoMutex.Lock()
	defer fake.getVMInfoMutex.Unlock()
	fake.GetVMInfoStub = nil
	fake.getVMInfoReturns = struct {
		result1 driver.VMInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetVMInfoReturnsOnCall(i int, result1 driver.VMInfo, result2 error) {
	fake.getVMInfoMutex.Lock()
	defer fake.getVMInfoMutex.Unlock()
	fake.GetVMInfoStub = nil
	if fake.getVMInfoReturnsOnCall == nil {
		fake.getVMInfoReturnsOnCall = make(map[int]struct {
			result1 driver.VMInfo
			result2 error
		})
	}
	fake.getVMInfoReturnsOnCall[i] = struct {
		result1 driver.VMInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetVMIsoPath(arg1 string) string {
	fake.getVMIsoPathMutex.Lock()
	ret, specificReturn := fake.getVMIsoPathReturnsOnCall[len(fake.getVMIsoPathArgsForCall)]
	fake.getVMIsoPathArgsForCall = append(fake.getVMIsoPathArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetVMIsoPath", []interface{}{arg1})
	fake.getVMIsoPathMutex.Unlock()
	if fake.GetVMIsoPathStub != nil {
		return fake.GetVMIsoPathStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getVMIsoPathReturns
	return fakeReturns.result1
}

func (fake *FakeClient) GetVMIsoPathCallCount() int {
	fake.getVMIsoPathMutex.RLock()
	defer fake.getVMIsoPathMutex.RUnlock()
	return len(fake.getVMIsoPathArgsForCall)
}

func (fake *FakeClient) GetVMIsoPathCalls(stub func(string) string) {
	fake.getVMIsoPathMutex.Lock()
	defer fake.getVMIsoPathMutex.Unlock()
	fake.GetVMIsoPathStub = stub
}

func (fake *FakeClient) GetVMIsoPathArgsForCall(i int) string {
	fake.getVMIsoPathMutex.RLock()
	defer fake.getVMIsoPathMutex.RUnlock()
	argsForCall := fake.getVMIsoPathArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) GetVMIsoPathReturns(result1 string) {
	fake.getVMIsoPathMutex.Lock()
	defer fake.getVMIsoPathMutex.Unlock()
	fake.GetVMIsoPathStub = nil
	fake.getVMIsoPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeClient) GetVMIsoPathReturnsOnCall(i int, result1 string) {
	fake.getVMIsoPathMutex.Lock()
	defer fake.getVMIsoPathMutex.Unlock()
	fake.GetVMIsoPathStub = nil
	if fake.getVMIsoPathReturnsOnCall == nil {
		fake.getVMIsoPathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getVMIsoPathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeClient) HasDisk(arg1 string) bool {
	fake.hasDiskMutex.Lock()
	ret, specificReturn := fake.hasDiskReturnsOnCall[len(fake.hasDiskArgsForCall)]
	fake.hasDiskArgsForCall = append(fake.hasDiskArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("HasDisk", []interface{}{arg1})
	fake.hasDiskMutex.Unlock()
	if fake.HasDiskStub != nil {
		return fake.HasDiskStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.hasDiskReturns
	return fakeReturns.result1
}

func (fake *FakeClient) HasDiskCallCount() int {
	fake.hasDiskMutex.RLock()
	defer fake.hasDiskMutex.RUnlock()
	return len(fake.hasDiskArgsForCall)
}

func (fake *FakeClient) HasDiskCalls(stub func(string) bool) {
	fake.hasDiskMutex.Lock()
	defer fake.hasDiskMutex.Unlock()
	fake.HasDiskStub = stub
}

func (fake *FakeClient) HasDiskArgsForCall(i int) string {
	fake.hasDiskMutex.RLock()
	defer fake.hasDiskMutex.RUnlock()
	argsForCall := fake.hasDiskArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) HasDiskReturns(result1 bool) {
	fake.hasDiskMutex.Lock()
	defer fake.hasDiskMutex.Unlock()
	fake.HasDiskStub = nil
	fake.hasDiskReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeClient) HasDiskReturnsOnCall(i int, result1 bool) {
	fake.hasDiskMutex.Lock()
	defer fake.hasDiskMutex.Unlock()
	fake.HasDiskStub = nil
	if fake.hasDiskReturnsOnCall == nil {
		fake.hasDiskReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.hasDiskReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeClient) HasVM(arg1 string) bool {
	fake.hasVMMutex.Lock()
	ret, specificReturn := fake.hasVMReturnsOnCall[len(fake.hasVMArgsForCall)]
	fake.hasVMArgsForCall = append(fake.hasVMArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("HasVM", []interface{}{arg1})
	fake.hasVMMutex.Unlock()
	if fake.HasVMStub != nil {
		return fake.HasVMStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.hasVMReturns
	return fakeReturns.result1
}

func (fake *FakeClient) HasVMCallCount() int {
	fake.hasVMMutex.RLock()
	defer fake.hasVMMutex.RUnlock()
	return len(fake.hasVMArgsForCall)
}

func (fake *FakeClient) HasVMCalls(stub func(string) bool) {
	fake.hasVMMutex.Lock()
	defer fake.hasVMMutex.Unlock()
	fake.HasVMStub = stub
}

func (fake *FakeClient) HasVMArgsForCall(i int) string {
	fake.hasVMMutex.RLock()
	defer fake.hasVMMutex.RUnlock()
	argsForCall := fake.hasVMArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) HasVMReturns(result1 bool) {
	fake.hasVMMutex.Lock()
	defer fake.hasVMMutex.Unlock()
	fake.HasVMStub = nil
	fake.hasVMReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeClient) HasVMReturnsOnCall(i int, result1 bool) {
	fake.hasVMMutex.Lock()
	defer fake.hasVMMutex.Unlock()
	fake.HasVMStub = nil
	if fake.hasVMReturnsOnCall == nil {
		fake.hasVMReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.hasVMReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeClient) ImportOvf(arg1 string, arg2 string) (bool, error) {
	fake.importOvfMutex.Lock()
	ret, specificReturn := fake.importOvfReturnsOnCall[len(fake.importOvfArgsForCall)]
	fake.importOvfArgsForCall = append(fake.importOvfArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ImportOvf", []interface{}{arg1, arg2})
	fake.importOvfMutex.Unlock()
	if fake.ImportOvfStub != nil {
		return fake.ImportOvfStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.importOvfReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) ImportOvfCallCount() int {
	fake.importOvfMutex.RLock()
	defer fake.importOvfMutex.RUnlock()
	return len(fake.importOvfArgsForCall)
}

func (fake *FakeClient) ImportOvfCalls(stub func(string, string) (bool, error)) {
	fake.importOvfMutex.Lock()
	defer fake.importOvfMutex.Unlock()
	fake.ImportOvfStub = stub
}

func (fake *FakeClient) ImportOvfArgsForCall(i int) (string, string) {
	fake.importOvfMutex.RLock()
	defer fake.importOvfMutex.RUnlock()
	argsForCall := fake.importOvfArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) ImportOvfReturns(result1 bool, result2 error) {
	fake.importOvfMutex.Lock()
	defer fake.importOvfMutex.Unlock()
	fake.ImportOvfStub = nil
	fake.importOvfReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ImportOvfReturnsOnCall(i int, result1 bool, result2 error) {
	fake.importOvfMutex.Lock()
	defer fake.importOvfMutex.Unlock()
	fake.ImportOvfStub = nil
	if fake.importOvfReturnsOnCall == nil {
		fake.importOvfReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.importOvfReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) SetVMNetworkAdapter(arg1 string, arg2 string, arg3 string) error {
	fake.setVMNetworkAdapterMutex.Lock()
	ret, specificReturn := fake.setVMNetworkAdapterReturnsOnCall[len(fake.setVMNetworkAdapterArgsForCall)]
	fake.setVMNetworkAdapterArgsForCall = append(fake.setVMNetworkAdapterArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("SetVMNetworkAdapter", []interface{}{arg1, arg2, arg3})
	fake.setVMNetworkAdapterMutex.Unlock()
	if fake.SetVMNetworkAdapterStub != nil {
		return fake.SetVMNetworkAdapterStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setVMNetworkAdapterReturns
	return fakeReturns.result1
}

func (fake *FakeClient) SetVMNetworkAdapterCallCount() int {
	fake.setVMNetworkAdapterMutex.RLock()
	defer fake.setVMNetworkAdapterMutex.RUnlock()
	return len(fake.setVMNetworkAdapterArgsForCall)
}

func (fake *FakeClient) SetVMNetworkAdapterCalls(stub func(string, string, string) error) {
	fake.setVMNetworkAdapterMutex.Lock()
	defer fake.setVMNetworkAdapterMutex.Unlock()
	fake.SetVMNetworkAdapterStub = stub
}

func (fake *FakeClient) SetVMNetworkAdapterArgsForCall(i int) (string, string, string) {
	fake.setVMNetworkAdapterMutex.RLock()
	defer fake.setVMNetworkAdapterMutex.RUnlock()
	argsForCall := fake.setVMNetworkAdapterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClient) SetVMNetworkAdapterReturns(result1 error) {
	fake.setVMNetworkAdapterMutex.Lock()
	defer fake.setVMNetworkAdapterMutex.Unlock()
	fake.SetVMNetworkAdapterStub = nil
	fake.setVMNetworkAdapterReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) SetVMNetworkAdapterReturnsOnCall(i int, result1 error) {
	fake.setVMNetworkAdapterMutex.Lock()
	defer fake.setVMNetworkAdapterMutex.Unlock()
	fake.SetVMNetworkAdapterStub = nil
	if fake.setVMNetworkAdapterReturnsOnCall == nil {
		fake.setVMNetworkAdapterReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setVMNetworkAdapterReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) SetVMResources(arg1 string, arg2 int, arg3 int) error {
	fake.setVMResourcesMutex.Lock()
	ret, specificReturn := fake.setVMResourcesReturnsOnCall[len(fake.setVMResourcesArgsForCall)]
	fake.setVMResourcesArgsForCall = append(fake.setVMResourcesArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("SetVMResources", []interface{}{arg1, arg2, arg3})
	fake.setVMResourcesMutex.Unlock()
	if fake.SetVMResourcesStub != nil {
		return fake.SetVMResourcesStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setVMResourcesReturns
	return fakeReturns.result1
}

func (fake *FakeClient) SetVMResourcesCallCount() int {
	fake.setVMResourcesMutex.RLock()
	defer fake.setVMResourcesMutex.RUnlock()
	return len(fake.setVMResourcesArgsForCall)
}

func (fake *FakeClient) SetVMResourcesCalls(stub func(string, int, int) error) {
	fake.setVMResourcesMutex.Lock()
	defer fake.setVMResourcesMutex.Unlock()
	fake.SetVMResourcesStub = stub
}

func (fake *FakeClient) SetVMResourcesArgsForCall(i int) (string, int, int) {
	fake.setVMResourcesMutex.RLock()
	defer fake.setVMResourcesMutex.RUnlock()
	argsForCall := fake.setVMResourcesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClient) SetVMResourcesReturns(result1 error) {
	fake.setVMResourcesMutex.Lock()
	defer fake.setVMResourcesMutex.Unlock()
	fake.SetVMResourcesStub = nil
	fake.setVMResourcesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) SetVMResourcesReturnsOnCall(i int, result1 error) {
	fake.setVMResourcesMutex.Lock()
	defer fake.setVMResourcesMutex.Unlock()
	fake.SetVMResourcesStub = nil
	if fake.setVMResourcesReturnsOnCall == nil {
		fake.setVMResourcesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setVMResourcesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) StartVM(arg1 string) error {
	fake.startVMMutex.Lock()
	ret, specificReturn := fake.startVMReturnsOnCall[len(fake.startVMArgsForCall)]
	fake.startVMArgsForCall = append(fake.startVMArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("StartVM", []interface{}{arg1})
	fake.startVMMutex.Unlock()
	if fake.StartVMStub != nil {
		return fake.StartVMStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.startVMReturns
	return fakeReturns.result1
}

func (fake *FakeClient) StartVMCallCount() int {
	fake.startVMMutex.RLock()
	defer fake.startVMMutex.RUnlock()
	return len(fake.startVMArgsForCall)
}

func (fake *FakeClient) StartVMCalls(stub func(string) error) {
	fake.startVMMutex.Lock()
	defer fake.startVMMutex.Unlock()
	fake.StartVMStub = stub
}

func (fake *FakeClient) StartVMArgsForCall(i int) string {
	fake.startVMMutex.RLock()
	defer fake.startVMMutex.RUnlock()
	argsForCall := fake.startVMArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) StartVMReturns(result1 error) {
	fake.startVMMutex.Lock()
	defer fake.startVMMutex.Unlock()
	fake.StartVMStub = nil
	fake.startVMReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) StartVMReturnsOnCall(i int, result1 error) {
	fake.startVMMutex.Lock()
	defer fake.startVMMutex.Unlock()
	fake.StartVMStub = nil
	if fake.startVMReturnsOnCall == nil {
		fake.startVMReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startVMReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) StopVM(arg1 string) error {
	fake.stopVMMutex.Lock()
	ret, specificReturn := fake.stopVMReturnsOnCall[len(fake.stopVMArgsForCall)]
	fake.stopVMArgsForCall = append(fake.stopVMArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("StopVM", []interface{}{arg1})
	fake.stopVMMutex.Unlock()
	if fake.StopVMStub != nil {
		return fake.StopVMStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stopVMReturns
	return fakeReturns.result1
}

func (fake *FakeClient) StopVMCallCount() int {
	fake.stopVMMutex.RLock()
	defer fake.stopVMMutex.RUnlock()
	return len(fake.stopVMArgsForCall)
}

func (fake *FakeClient) StopVMCalls(stub func(string) error) {
	fake.stopVMMutex.Lock()
	defer fake.stopVMMutex.Unlock()
	fake.StopVMStub = stub
}

func (fake *FakeClient) StopVMArgsForCall(i int) string {
	fake.stopVMMutex.RLock()
	defer fake.stopVMMutex.RUnlock()
	argsForCall := fake.stopVMArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) StopVMReturns(result1 error) {
	fake.stopVMMutex.Lock()
	defer fake.stopVMMutex.Unlock()
	fake.StopVMStub = nil
	fake.stopVMReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) StopVMReturnsOnCall(i int, result1 error) {
	fake.stopVMMutex.Lock()
	defer fake.stopVMMutex.Unlock()
	fake.StopVMStub = nil
	if fake.stopVMReturnsOnCall == nil {
		fake.stopVMReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopVMReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) UpdateVMIso(arg1 string, arg2 string) error {
	fake.updateVMIsoMutex.Lock()
	ret, specificReturn := fake.updateVMIsoReturnsOnCall[len(fake.updateVMIsoArgsForCall)]
	fake.updateVMIsoArgsForCall = append(fake.updateVMIsoArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("UpdateVMIso", []interface{}{arg1, arg2})
	fake.updateVMIsoMutex.Unlock()
	if fake.UpdateVMIsoStub != nil {
		return fake.UpdateVMIsoStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateVMIsoReturns
	return fakeReturns.result1
}

func (fake *FakeClient) UpdateVMIsoCallCount() int {
	fake.updateVMIsoMutex.RLock()
	defer fake.updateVMIsoMutex.RUnlock()
	return len(fake.updateVMIsoArgsForCall)
}

func (fake *FakeClient) UpdateVMIsoCalls(stub func(string, string) error) {
	fake.updateVMIsoMutex.Lock()
	defer fake.updateVMIsoMutex.Unlock()
	fake.UpdateVMIsoStub = stub
}

func (fake *FakeClient) UpdateVMIsoArgsForCall(i int) (string, string) {
	fake.updateVMIsoMutex.RLock()
	defer fake.updateVMIsoMutex.RUnlock()
	argsForCall := fake.updateVMIsoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) UpdateVMIsoReturns(result1 error) {
	fake.updateVMIsoMutex.Lock()
	defer fake.updateVMIsoMutex.Unlock()
	fake.UpdateVMIsoStub = nil
	fake.updateVMIsoReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) UpdateVMIsoReturnsOnCall(i int, result1 error) {
	fake.updateVMIsoMutex.Lock()
	defer fake.updateVMIsoMutex.Unlock()
	fake.UpdateVMIsoStub = nil
	if fake.updateVMIsoReturnsOnCall == nil {
		fake.updateVMIsoReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateVMIsoReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.attachDiskMutex.RLock()
	defer fake.attachDiskMutex.RUnlock()
	fake.bootstrapVMMutex.RLock()
	defer fake.bootstrapVMMutex.RUnlock()
	fake.cloneVMMutex.RLock()
	defer fake.cloneVMMutex.RUnlock()
	fake.createDiskMutex.RLock()
	defer fake.createDiskMutex.RUnlock()
	fake.createEphemeralDiskMutex.RLock()
	defer fake.createEphemeralDiskMutex.RUnlock()
	fake.destroyDiskMutex.RLock()
	defer fake.destroyDiskMutex.RUnlock()
	fake.destroyVMMutex.RLock()
	defer fake.destroyVMMutex.RUnlock()
	fake.detachDiskMutex.RLock()
	defer fake.detachDiskMutex.RUnlock()
	fake.getVMInfoMutex.RLock()
	defer fake.getVMInfoMutex.RUnlock()
	fake.getVMIsoPathMutex.RLock()
	defer fake.getVMIsoPathMutex.RUnlock()
	fake.hasDiskMutex.RLock()
	defer fake.hasDiskMutex.RUnlock()
	fake.hasVMMutex.RLock()
	defer fake.hasVMMutex.RUnlock()
	fake.importOvfMutex.RLock()
	defer fake.importOvfMutex.RUnlock()
	fake.setVMNetworkAdapterMutex.RLock()
	defer fake.setVMNetworkAdapterMutex.RUnlock()
	fake.setVMResourcesMutex.RLock()
	defer fake.setVMResourcesMutex.RUnlock()
	fake.startVMMutex.RLock()
	defer fake.startVMMutex.RUnlock()
	fake.stopVMMutex.RLock()
	defer fake.stopVMMutex.RUnlock()
	fake.updateVMIsoMutex.RLock()
	defer fake.updateVMIsoMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ driver.Client = new(FakeClient)
