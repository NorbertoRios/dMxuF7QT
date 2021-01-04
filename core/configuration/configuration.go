package configuration

import (
	"container/list"
	"genx-go/core/configuration/request"
	"genx-go/core/configuration/task"
	"genx-go/core/device/interfaces"
	baseRequest "genx-go/core/request"
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
func (config *Configuration) NewRequest(req baseRequest.IRequest) *list.List {
	cList := list.New()
	newTask := task.New(req.(*request.ConfigurationRequest), config.device, config)
	if config.currentTask != nil {
		cList.PushBackList(config.currentTask.Invoker().CanselTask(config.currentTask, "Deprecated"))
	}
	config.currentTask = newTask
	cList.PushBackList(config.currentTask.Commands())
	return cList
}

//TaskCancel ...
func (config *Configuration) TaskCancel(canseledTask interfaces.ITask, description string) {
	logger.Logger().WriteToLog(logger.Info, "Task is canceled. ", description)
	config.pushToTasks(task.NewCanceledConfigTask(canseledTask, description), false)
}

//TaskDone ...
func (config *Configuration) TaskDone(doneTask interfaces.ITask) {
	logger.Logger().WriteToLog(logger.Info, "Task is done")
	config.pushToTasks(task.NewDoneConfigTask(doneTask), true)
}

func (config *Configuration) pushToTasks(_task interfaces.ITask, isDone bool) {
	config.mutex.Lock()
	defer config.mutex.Unlock()
	if isDone {
		config.tasks.PushFront(_task)
	} else {
		config.tasks.PushBack(_task)
	}
}
