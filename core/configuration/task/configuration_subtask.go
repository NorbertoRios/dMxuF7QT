package task

import (
	"container/list"
	"genx-go/core/configuration/observers"
	"genx-go/core/device/interfaces"
	"genx-go/core/filter"
)

//NewConfigurationSubtask ..
func NewConfigurationSubtask(_task *ConfigTask, _command string) *ConfigurationSubtask {
	return &ConfigurationSubtask{
		mainTask: _task,
		command:  _command,
	}
}

//ConfigurationSubtask ..
type ConfigurationSubtask struct {
	mainTask *ConfigTask
	command  string
}

//Request ..
func (task *ConfigurationSubtask) Request() interface{} {
	return task.command
}

//Device ..
func (task *ConfigurationSubtask) Device() interfaces.IDevice {
	return task.mainTask.Device()
}

//Observers ...
func (task *ConfigurationSubtask) Observers() []interfaces.IObserver {
	f := filter.NewObserversFilter(task.Device().Observable())
	return f.Extract(task)
}

//Start ...
func (task *ConfigurationSubtask) Start() {
	cList := list.New()
	cList.PushFront(observers.NewSendConfigCommand(task))
	task.mainTask.Device().ProccessCommands(cList)
}

//Done ...
func (task *ConfigurationSubtask) Done() {

}

//Cancel ..
func (task *ConfigurationSubtask) Cancel(description string) {

}
