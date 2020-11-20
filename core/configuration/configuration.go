package configuration

import (
	"container/list"
	"genx-go/core/configuration/request"
	"genx-go/core/configuration/task"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
	"sync"
)

//Configuration ..
type Configuration struct {
	mutex       *sync.Mutex
	device      interfaces.IDevice
	currentTask interfaces.ITask
	tasks       *list.List
}

//NewRequest ..
func (config *Configuration) NewRequest(req *request.ConfigurationRequest) {
	configTask := task.NewConfigTask(req, config.device, config.Cancel, config.Done)
	if config.currentTask == nil {
		config.currentTask = configTask
	} else {
		config.pauseCurrentTask()
		config.mergeTasks(configTask)
	}
	config.currentTask.Start()
}

func (config *Configuration) mergeTasks(_task *task.ConfigTask) {
	newTask := _task.Merge(config.currentTask.(*task.ConfigTask))
	config.currentTask.Cancel("Depricated")
	config.currentTask = newTask
}

func (config *Configuration) pauseCurrentTask() {
	cList := list.New()
	for _, observer := range config.currentTask.Observers() {
		cList.PushBack(observers.NewDetachObserverCommand(observer))
	}
	config.device.ProccessCommands(cList)
}

//Done ...
func (config *Configuration) Done(_task *task.ConfigTask) {

}

//Cancel ...
func (config *Configuration) Cancel(_task *task.ConfigTask, description string) {

}
