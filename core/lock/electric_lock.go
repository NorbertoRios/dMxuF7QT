package lock

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/lock/request"
	"genx-go/core/lock/task"
	"genx-go/core/process"
	baseRequest "genx-go/core/request"
	"genx-go/logger"
	"reflect"
	"sync"
)

//NewElectricLock ...
func NewElectricLock(_outNum int) *ElectricLock {
	lock := &ElectricLock{
		OutputNumber: _outNum,
	}
	lock.Mutex = &sync.Mutex{}
	lock.ProcessTasks = list.New()
	//lock.ProcessDevice = _device
	return lock
}

//ElectricLock ...
type ElectricLock struct {
	process.BaseProcess
	OutputNumber int
}

//NewRequest ..
func (lock *ElectricLock) NewRequest(req baseRequest.IRequest, _device interfaces.IDevice) *list.List {
	newTask := task.NewElectricLockTask(req, _device, lock)
	if lock.ProcessCurrentTask == nil {
		lock.ProcessCurrentTask = newTask
		return lock.ProcessCurrentTask.Commands()
	}
	return lock.competitivenessOfTasks(newTask, lock.ProcessCurrentTask.Request().(*request.UnlockRequest))
}

func (lock *ElectricLock) competitivenessOfTasks(newTask interfaces.ITask, currentReq *request.UnlockRequest) *list.List {
	if currentReq.Equal(newTask.Request().(*request.UnlockRequest)) {
		return newTask.Invoker().CanselTask(newTask, "Duplicate")
	}
	cmdList := list.New()
	cmdList.PushBackList(lock.ProcessCurrentTask.Invoker().CanselTask(lock.ProcessCurrentTask, "Deprecated"))
	lock.ProcessCurrentTask = newTask
	cmdList.PushBackList(lock.ProcessCurrentTask.Commands())
	logger.Logger().WriteToLog(logger.Info, "[ElectricLock] Task created and run")
	return cmdList
}

//TaskCancel ...
func (lock *ElectricLock) TaskCancel(canseledTask interfaces.ITask, description string) {
	if reflect.DeepEqual(canseledTask, lock.ProcessCurrentTask) {
		lock.ProcessCurrentTask = nil
		logger.Logger().WriteToLog(logger.Info, "[ElectricLock] Current task is canceled. ", description)
	} else {
		logger.Logger().WriteToLog(logger.Info, "[ElectricLock] Task is canceled. ", description)
	}
	lock.PushToTasks(task.NewCanceledElectricLockTask(canseledTask, description), false)
}

//TaskDone ...
func (lock *ElectricLock) TaskDone(doneTask interfaces.ITask) {
	lock.ProcessCurrentTask = nil
	logger.Logger().WriteToLog(logger.Info, "[ElectricLock] Task is done")
	lock.PushToTasks(task.NewDoneElectricLockTask(doneTask), true)
}
