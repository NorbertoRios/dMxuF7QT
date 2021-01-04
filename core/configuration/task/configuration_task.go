package task

import (
	"container/list"
	"genx-go/core/configuration/request"
	"genx-go/core/device/interfaces"
	"genx-go/core/filter"
	"genx-go/core/invoker"
	"time"
)

//New ...
func New(_request *request.ConfigurationRequest, device interfaces.IDevice, _config interfaces.IProcess) *ConfigTask {
	return &ConfigTask{
		BornTime:      time.Now().UTC(),
		FacadeRequest: _request,
		iterator:      newConfigIterator(_request.Commands()),
		device:        device,
		invoker:       invoker.NewConfigInvoker(_config),
	}
}

//ConfigTask ...
type ConfigTask struct {
	BornTime      time.Time                     `json:"CreatedAt"`
	FacadeRequest *request.ConfigurationRequest `json:"FacadeRequest"`
	iterator      *ConfigIterator
	device        interfaces.IDevice
	invoker       interfaces.IConfigInvoker
}

//Commands ...
func (task *ConfigTask) Commands() *list.List {
	return task.invoker.SendCurrectCommand(task)
}

//Device ...
func (task *ConfigTask) Device() interfaces.IDevice {
	return task.device
}

//ConfigCommands ...
func (task *ConfigTask) ConfigCommands() *list.List {
	return task.iterator.configCommands()
}

//Invoker ..
func (task *ConfigTask) Invoker() interfaces.IInvoker {
	return task.invoker
}

//CurrentStringCommand ...
func (task *ConfigTask) CurrentStringCommand() string {
	return task.iterator.current().Command()
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

//CommandComplete ...
func (task *ConfigTask) CommandComplete() {
	task.iterator.current().Complete()
}

//IsNextExist ...
func (task *ConfigTask) IsNextExist() bool {
	return task.iterator.nextExisting()
}

//GoToNextCommand ..
func (task *ConfigTask) GoToNextCommand() {
	task.iterator.goToNext()
}
