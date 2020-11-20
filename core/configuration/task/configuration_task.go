package task

import (
	"container/list"
	"genx-go/core/configuration/observers"
	"genx-go/core/configuration/request"
	"genx-go/core/device/interfaces"
	"genx-go/core/filter"
	"time"
)

//NewConfigTask ...
func NewConfigTask(_commands, _sentCommands *list.List, _request *request.ConfigurationRequest, device interfaces.IDevice, _onCancel func(*ConfigTask, string), _onDone func(*ConfigTask)) *ConfigTask {
	return &ConfigTask{
		BornTime:      time.Now().UTC(),
		FacadeRequest: _request,
		Commands:      _commands,
		SentCommands:  _sentCommands,
		device:        device,
		onCancel:      _onCancel,
		onDone:        _onDone,
	}
}

//New ...
func New(_request *request.ConfigurationRequest, device interfaces.IDevice, _onCancel func(*ConfigTask, string), _onDone func(*ConfigTask)) *ConfigTask {
	return &ConfigTask{
		BornTime:      time.Now().UTC(),
		FacadeRequest: _request,
		Commands:      _request.Commands(),
		SentCommands:  list.New(),
		device:        device,
		onCancel:      _onCancel,
		onDone:        _onDone,
	}
}

//ConfigTask ...
type ConfigTask struct {
	BornTime       time.Time                     `json:"CreatedAt"`
	FacadeRequest  *request.ConfigurationRequest `json:"FacadeRequest"`
	Commands       *list.List                    `json:"Commands"`
	SentCommands   *list.List                    `json:"SentCommands"`
	currentCommand *list.Element
	device         interfaces.IDevice
	onCancel       func(*ConfigTask, string)
	onDone         func(*ConfigTask)
}

//Device ...
func (task *ConfigTask) Device() interfaces.IDevice {
	return task.device
}

//Start ...
func (task *ConfigTask) Start() {
	if task.currentCommand == nil {
		task.currentCommand = task.Commands.Front()
	}
	task.sendCurrentCommand()
}

func (task *ConfigTask) sendCurrentCommand() {
	cList := list.New()
	cList.PushBack(observers.NewSendConfigCommand(task, task.currentCommand.Value.(string)))
	task.device.ProccessCommands(cList)
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

//Done ...
func (task *ConfigTask) Done() {
	task.SentCommands.PushBack(task.currentCommand)
	task.Commands.Remove(task.currentCommand)
	if task.Commands.Len() > 0 {
		task.currentCommand = task.Commands.Front()
		task.sendCurrentCommand()
	} else {
		task.onDone(task)
	}
}

//Cancel ...
func (task *ConfigTask) Cancel(description string) {
	task.onCancel(task, description)
}
