package task

import (
	"genx-go/core/device/interfaces"
	"time"
)

//BaseTask ..
type BaseTask struct {
	bornTime time.Time
	device   interfaces.IDevice
	invoker  interfaces.ILocationProcessInvoker
}

//Device ...
func (task *BaseTask) Device() interfaces.IDevice {
	return task.device
}

//Request ...
func (task *BaseTask) Request() interface{} {
	return struct{}{}
}

//Invoker ...
func (task *BaseTask) Invoker() interfaces.IInvoker {
	return task.invoker
}
