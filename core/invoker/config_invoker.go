package invoker

import (
	"container/list"
	confObservers "genx-go/core/configuration/observers"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
	"genx-go/logger"
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

//SendCurrectCommand ...
func (invoker *ConfigInvoker) SendCurrectCommand(task interfaces.IConfigTask) *list.List {
	cmd := list.New()
	cmd.PushBack(confObservers.NewSendConfigCommand(task))
	return cmd
}

//SendNextCommand ...
func (invoker *ConfigInvoker) SendNextCommand(task interfaces.IConfigTask) *list.List {
	if !task.IsNextExist() {
		logger.Logger().WriteToLog(logger.Info, "[ConfigInvoker | SendNextCommand] There are no more unsent commands left in the task. The task is done")
		return invoker.DoneTask(task)
	}
	task.GoToNextCommand()
	return invoker.SendCurrectCommand(task)
}
