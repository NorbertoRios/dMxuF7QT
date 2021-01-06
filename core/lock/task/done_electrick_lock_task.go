package task

import (
	"genx-go/core/device/interfaces"
	"time"
)

//NewDoneElectricLockTask ...
func NewDoneElectricLockTask(_task interfaces.ITask) interfaces.ITask {
	return &DoneElectricLockTask{
		Task:     _task,
		doneTime: time.Now().UTC(),
	}
}

//DoneElectricLockTask ...
type DoneElectricLockTask struct {
	ElectricLockTask
	Task     interfaces.ITask
	doneTime time.Time
}
