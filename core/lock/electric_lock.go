package lock

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/lock/request"
	"genx-go/core/lock/task"
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
func (lock *ElectricLock) NewRequest(req *request.UnlockRequest) {
	newTask := task.NewElectricLockTask(req, lock.device, lock.taskCanceled, lock.taskDone)
	if lock.currentTask == nil {
		lock.currentTask = newTask
		lock.currentTask.Start()
		return
	}
	lock.competitivenessOfTasks(newTask)
}

func (lock *ElectricLock) competitivenessOfTasks(newTask interfaces.ITask) {
	if lock.currentTask.Request().(*request.UnlockRequest).Equal(newTask.Request().(*request.UnlockRequest)) {
		newTask.Cancel("Duplicate")
	} else {
		lock.currentTask.Cancel("Deprecated")
		lock.currentTask = newTask
		lock.currentTask.Start()
		logger.Logger().WriteToLog(logger.Info, "Task created and run")
	}
}

//CurrentTask ..
func (lock *ElectricLock) CurrentTask() interfaces.ITask {
	return lock.currentTask
}

//Tasks ...
func (lock *ElectricLock) Tasks() *list.List {
	return lock.tasks
}

func (lock *ElectricLock) taskCanceled(cancelTask *task.ElectricLockTask, description string) {
	logger.Logger().WriteToLog(logger.Info, "Task is canceled. ", description)
	lock.mutex.Lock()
	lock.tasks.PushBack(task.NewCanceledElectricLockTask(cancelTask, description))
	lock.mutex.Unlock()
	lock.currentTask = nil
}

func (lock *ElectricLock) taskDone(doneTask *task.ElectricLockTask) {
	logger.Logger().WriteToLog(logger.Info, "Task is done")
	lock.mutex.Lock()
	lock.tasks.PushFront(task.NewDoneElectricLockTask(doneTask))
	lock.mutex.Unlock()
	lock.currentTask = nil
}
