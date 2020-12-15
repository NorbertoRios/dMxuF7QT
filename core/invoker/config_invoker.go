package invoker

import (
	"container/list"
	"genx-go/core/device/interfaces"
)

//NewConfigInvoker ...
func NewConfigInvoker(_process interfaces.IProcess) *ConfigInvoker {
	invoker := &ConfigInvoker{}
	invoker.process = _process
	return invoker
}

//ConfigInvoker ...
type ConfigInvoker struct {
	BaseInvoker
}

//Next ...
func (invoker *ConfigInvoker) Next(_task interfaces.ITask) *list.List {
	return _task.(interfaces.IConfigTask).NextStep()
}

//SendConfigAfterAnyMessage ...
func (invoker *ConfigInvoker) SendConfigAfterAnyMessage(_task interfaces.ITask) *list.List {

}
