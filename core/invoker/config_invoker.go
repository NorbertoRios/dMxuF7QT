package invoker

import (
	"container/list"
	confObservers "genx-go/core/configuration/observers"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
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

//SendConfigAfterAnyMessage ...
func (invoker *ConfigInvoker) SendConfigAfterAnyMessage(task interfaces.IConfigTask) *list.List {
	cmd := list.New()
	cmd.PushBack(observers.NewAttachObserverCommand(confObservers.NewAnyMessageObserver(task)))
	return cmd
}

//SendCommand ...
func (invoker *ConfigInvoker) SendCommand(task interfaces.IConfigTask) *list.List {
	cmd := list.New()
	cmd.PushBack(confObservers.NewSendConfigCommand(task))
	return cmd
}
