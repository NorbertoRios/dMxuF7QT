package lock

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/lock/request"
	"genx-go/core/lock/task"
	baseRequest "genx-go/core/request"
	"genx-go/logger"
	"sync"
)

//NewElectricLock ...
func NewElectricLock(_device interfaces.IDevice, _outNum int) *ElectricLock {
	return &ElectricLock{
		mutex:        &sync.Mutex{},
		device:       _device,
		OutputNumber: _outNum,
		tasks:        list.New(),
	}
}

//ElectricLock ...
type ElectricLock struct {
	mutex        *sync.Mutex
	device       interfaces.IDevice
	OutputNumber int
	currentTask  interfaces.ITask
	tasks        *list.List
}

//NewRequest ..
func (lock *ElectricLock) NewRequest(req baseRequest.IRequest) *list.List {
	newTask := task.NewElectricLockTask(req, lock.device, lock)
	if lock.currentTask == nil {
		lock.currentTask = newTask
		return lock.currentTask.Commands()
	}
	return lock.competitivenessOfTasks(newTask, lock.currentTask.Request().(*request.UnlockRequest))
}

func (lock *ElectricLock) competitivenessOfTasks(newTask interfaces.ITask, currentReq *request.UnlockRequest) *list.List {
	if currentReq.Equal(newTask.Request().(*request.UnlockRequest)) {
		return newTask.Invoker().CanselTask(newTask, "Duplicate")
	}
	cmdList := list.New()
	cmdList.PushBackList(lock.currentTask.Invoker().CanselTask(lock.currentTask, "Deprecated"))
	lock.currentTask = newTask
	cmdList.PushBackList(lock.currentTask.Commands())
	logger.Logger().WriteToLog(logger.Info, "Task created and run")
	return cmdList
}

//CurrentTask ..
func (lock *ElectricLock) CurrentTask() interfaces.ITask {
	return lock.currentTask
}

//Tasks ...
func (lock *ElectricLock) Tasks() *list.List {
	return lock.tasks
}

//TaskCancel ...
func (lock *ElectricLock) TaskCancel(canseledTask interfaces.ITask, description string) {
	logger.Logger().WriteToLog(logger.Info, "Task is canceled. ", description)
	lock.pushToTasks(task.NewCanceledElectricLockTask(canseledTask, description), false)
}

//TaskDone ...
func (lock *ElectricLock) TaskDone(doneTask interfaces.ITask) {
	logger.Logger().WriteToLog(logger.Info, "Task is done")
	lock.pushToTasks(task.NewDoneElectricLockTask(doneTask), true)
}

func (lock *ElectricLock) pushToTasks(_task interfaces.ITask, isDone bool) {
	lock.mutex.Lock()
	defer lock.mutex.Unlock()
	if isDone {
		lock.tasks.PushFront(_task)
	} else {
		lock.tasks.PushBack(_task)
	}
}
