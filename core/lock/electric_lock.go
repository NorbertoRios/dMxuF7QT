package device

import (
	"container/list"
	"genx-go/core/device/lock/request"
	"genx-go/logger"
	"sync"
)

//ElectricLock ...
type ElectricLock struct {
	mutex        *sync.Mutex
	device       IDevice
	OutputNumber int
	CurrentTask  *ElectricLockTask
	Tasks        *list.List
}

//NewRequest ..
func (lock *ElectricLock) NewRequest(req *request.UnlockRequest) {
	newTask := NewElectricLockTask(req, lock.device, lock.taskCanceled, lock.taskDone)
	if lock.CurrentTask == nil {
		lock.CurrentTask = newTask
		lock.CurrentTask.Start()
		return
	}
	lock.competitivenessOfTasks(newTask)
}

func (lock *ElectricLock) competitivenessOfTasks(task *ElectricLockTask) {
	if lock.CurrentTask.Request.Equal(task.Request) {
		task.Cancel("Duplicate")
	} else {
		lock.CurrentTask.Cancel("Deprecated")
		lock.CurrentTask = task
		lock.CurrentTask.Start()
		logger.Logger().WriteToLog(logger.Info, "Task created and run")
	}
}

func (lock *ElectricLock) taskCanceled(task *ElectricLockTask, description string) {
	logger.Logger().WriteToLog(logger.Info, "Task is canceled")
	lock.mutex.Lock()
	lock.Tasks.PushBack(NewCanceledElectricLockTask(task, description))
	lock.mutex.Unlock()
	lock.CurrentTask = nil
}

func (lock *ElectricLock) taskDone(task *ElectricLockTask) {
	logger.Logger().WriteToLog(logger.Info, "Task is done")
	lock.mutex.Lock()
	lock.Tasks.PushFront(NewDoneElectricLockTask(task))
	lock.mutex.Unlock()
	lock.CurrentTask = nil
}
