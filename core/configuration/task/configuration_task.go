package task

import (
	"container/list"
	"genx-go/core/configuration/observers"
	"genx-go/core/configuration/request"
	"genx-go/core/device/interfaces"
	"genx-go/core/filter"
	"genx-go/core/invoker"
	"time"
)

//New ...
func New(_request *request.ConfigurationRequest, device interfaces.IDevice, _config interfaces.IProcess) *ConfigTask {
	return &ConfigTask{
		BornTime:       time.Now().UTC(),
		FacadeRequest:  _request,
		ConfigCommands: _request.Commands(),
		device:         device,
		invoker:        invoker.NewConfigInvoker(_config),
	}
}

//ConfigTask ...
type ConfigTask struct {
	BornTime       time.Time                     `json:"CreatedAt"`
	FacadeRequest  *request.ConfigurationRequest `json:"FacadeRequest"`
	ConfigCommands *list.List                    `json:"Commands"`
	currentCommand *list.Element
	device         interfaces.IDevice
	invoker        interfaces.IConfigInvoker
}

//Device ...
func (task *ConfigTask) Device() interfaces.IDevice {
	return task.device
}

//Commands ...
func (task *ConfigTask) Commands() *list.List {
	if task.currentCommand == nil {
		task.currentCommand = task.ConfigCommands.Front()
	}
	cList := list.New()
	cList.PushBack(observers.NewSendConfigCommand(task, task.currentCommand.Value.(*request.Command).Command()))
	return cList
}

//Invoker ..
func (task *ConfigTask) Invoker() interfaces.IInvoker {
	return task.invoker
}

//Observers ...
func (task *ConfigTask) Observers() []interfaces.IObserver {
	f := filter.NewObserversFilter(task.device.Observable())
	return f.Extract(task)
}

//Request ...
func (task *ConfigTask) Request() interface{} {
	return task.FacadeRequest
}

//NextStep ...
func (task *ConfigTask) NextStep() *list.List {
	task.currentCommand.Value.(*request.Command).Complete()
	if cmd := task.currentCommand.Next(); cmd != nil {
		cList := list.New()
		task.currentCommand = cmd
		cList.PushBack(observers.NewSendConfigCommand(task, task.currentCommand.Value.(*request.Command).Command()))
		return cList
	}
	return task.Invoker().DoneTask(task)
}
