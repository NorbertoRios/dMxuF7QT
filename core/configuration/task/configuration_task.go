package task

import (
	"container/list"
	"genx-go/core/configuration/request"
	"genx-go/core/device/interfaces"
	"time"
)

//NewConfigTask ...
func NewConfigTask(_request *request.ConfigurationRequest, device interfaces.IDevice, _onCancel func(*ConfigTask, string), _onDone func(*ConfigTask)) *ConfigTask {
	task := &ConfigTask{
		BornTime:      time.Now().UTC(),
		FacadeRequest: _request,
		subTasks:      list.New(),
		device:        device,
		onCancel:      _onCancel,
		onDone:        _onDone,
	}
	task.queuedSubTasks = buildSubtasks(_request.Commands(), task)
	return task
}

func buildSubtasks(configs []string, task *ConfigTask) *list.List {
	sList := list.New()
	for _, config := range configs {
		sList.PushBack(NewConfigurationSubtask(task, config))
	}
	return sList
}

//ConfigTask ...
type ConfigTask struct {
	BornTime       time.Time                     `json:"CreatedAt"`
	FacadeRequest  *request.ConfigurationRequest `json:"FacadeRequest"`
	currentSubTask *ConfigurationSubtask
	queuedSubTasks *list.List
	subTasks       *list.List
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
	subtask := task.queuedSubTasks.Front()
	task.currentSubTask = subtask.Value.(*ConfigurationSubtask)
	task.queuedSubTasks.Remove(subtask)
	task.currentSubTask.Start()
}

func (task *ConfigTask) Merge(_task *ConfigTask) *ConfigTask {

}

//Observers ...
func (task *ConfigTask) Observers() []interfaces.IObserver {
	return task.currentSubTask.Observers()
}

//Request ...
func (task *ConfigTask) Request() interface{} {
	return task.FacadeRequest
}

//Done ...
func (task *ConfigTask) Done() {

}

//Cancel ...
func (task *ConfigTask) Cancel(description string) {

}
