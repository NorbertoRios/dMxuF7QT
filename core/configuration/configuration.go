package configuration

import (
	"container/list"
	"genx-go/core/configuration/request"
	"genx-go/core/configuration/task"
	"genx-go/core/device/interfaces"
	core "genx-go/core/interfaces"
	"genx-go/core/observers"
	"genx-go/logger"
	"sync"
)

//NewConfiguration ...
func NewConfiguration(_device interfaces.IDevice, _client core.IClient) *Configuration {
	return &Configuration{
		mutex:  &sync.Mutex{},
		device: _device,
		tasks:  list.New(),
		client: _client,
	}
}

//Configuration ..
type Configuration struct {
	mutex       *sync.Mutex
	device      interfaces.IDevice
	currentTask interfaces.ITask
	tasks       *list.List
	client      core.IClient
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
func (config *Configuration) NewRequest(req *request.ConfigurationRequest) {
	configTask := task.New(req, config.device, config.taskCancel, config.taskDone)
	if config.currentTask == nil {
		config.currentTask = configTask
	} else {
		config.pauseCurrentTask()
		config.competitivenessOfTasks(req)
	}
	config.currentTask.Start()
}

func (config *Configuration) competitivenessOfTasks(req *request.ConfigurationRequest) {
	commands := config.client.Execute(NewFacadeRequest(req.Identity, config.currentTask.(*task.ConfigTask).SentCommands, req.Commands()))
	if commands.(*list.List).Len() == 0 {
		logger.Logger().WriteToLog(logger.Info, "[Configuration | competitivenessOfTasks] Facade response does not contains commands. The current task remains the same.")
		return
	}
	newTask := task.NewConfigTask(commands.(*list.List), config.currentTask.(*task.ConfigTask).SentCommands, req, config.device, config.taskCancel, config.taskDone)
	config.currentTask.Cancel("Deprecated")
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
