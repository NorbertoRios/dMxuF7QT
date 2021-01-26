package configuration

import (
	"container/list"
	"genx-go/core/configuration/request"
	"genx-go/core/configuration/task"
	"genx-go/core/device/interfaces"
	"genx-go/core/process"
	baseRequest "genx-go/core/request"
	"genx-go/logger"
	"reflect"
	"sync"
)

//NewConfiguration ...
func NewConfiguration() *Configuration {
	config := &Configuration{}
	config.Mutex = &sync.Mutex{}
	config.ProcessTasks = list.New()
	return config
}

//Configuration ..
type Configuration struct {
	process.BaseProcess
}

//NewRequest ..
func (config *Configuration) NewRequest(req baseRequest.IRequest, device interfaces.IDevice) *list.List {
	cList := list.New()
	newTask := task.New(req.(*request.ConfigurationRequest), device, config)
	if config.ProcessCurrentTask != nil {
		if _, v := config.ProcessCurrentTask.(*task.ConfigTask); v {
			cList.PushBackList(config.ProcessCurrentTask.Invoker().CanselTask(config.ProcessCurrentTask, "Deprecated"))
		}
	}
	config.ProcessCurrentTask = newTask
	cList.PushBackList(config.ProcessCurrentTask.Commands())
	return cList
}

//TaskCancel ...
func (config *Configuration) TaskCancel(canseledTask interfaces.ITask, description string) {
	if reflect.DeepEqual(canseledTask, config.ProcessCurrentTask) {
		config.ProcessCurrentTask = nil
		logger.Logger().WriteToLog(logger.Info, "[Configuration] Current task is canceled. ", description)
	}
	logger.Logger().WriteToLog(logger.Info, "[Configuration] Task is canceled. ", description)
	config.PushToTasks(task.NewCanceledConfigTask(canseledTask, description), false)
}

//TaskDone ...
func (config *Configuration) TaskDone(doneTask interfaces.ITask) {
	config.ProcessCurrentTask = nil
	logger.Logger().WriteToLog(logger.Info, "Task is done")
	config.PushToTasks(task.NewDoneConfigTask(doneTask), true)
}
