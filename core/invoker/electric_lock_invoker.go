package invoker

import "genx-go/core/device/interfaces"

//NewElectricLockInvoker ...
func NewElectricLockInvoker(_process interfaces.IProcess) *ElectricLockInvoker {
	invoker := &ElectricLockInvoker{}
	invoker.process = _process
	return invoker
}

//ElectricLockInvoker ...
type ElectricLockInvoker struct {
	BaseInvoker
}
