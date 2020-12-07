package configuration

import (
	"container/list"
	"genx-go/core/configuration/request"
	"genx-go/core/configuration/task"
	"genx-go/core/device/interfaces"
	"genx-go/logger"
	"sync"
)

//NewConfiguration ...
func NewConfiguration(_device interfaces.IDevice) *Configuration {
	return &Configuration{
		mutex:  &sync.Mutex{},
		device: _device,
		tasks:  list.New(),
	}
}

//Configuration ..
type Configuration struct {
	mutex       *sync.Mutex
	device      interfaces.IDevice
	currentTask interfaces.ITask
	tasks       *list.List
}

//CurrentTask ...
func (config *Configuration) CurrentTask() interfaces.ITask {
	return config.currentTask
}

//Tasks ...
func (config *Configuration) Tasks() *list.List {
	return config.tasks
}

//NewRequest ..
func (config *Configuration) NewRequest(req *request.ConfigurationRequest) *list.List {
	newTask := task.New(req, config.device, config.taskCancel, config.taskDone)
	if config.currentTask != nil {
		config.currentTask.Cancel("Deprecated")
	}
	config.currentTask = newTask
	return config.currentTask.Commands()
}

//Done ...
func (config *Configuration) taskDone(_task *task.ConfigTask) {
	logger.Logger().WriteToLog(logger.Info, "Task is done")
	config.mutex.Lock()
	config.tasks.PushFront(task.NewDoneConfigTask(_task))
	config.mutex.Unlock()
	config.currentTask = nil
}

//Cancel ...
func (config *Configuration) taskCancel(_task *task.ConfigTask, description string) {
	logger.Logger().WriteToLog(logger.Info, "Task is canceled. ", description)
	config.mutex.Lock()
	config.tasks.PushBack(task.NewCanceledConfigTask(_task, description))
	config.mutex.Unlock()
	config.currentTask = nil
}
