package task

import (
	"genx-go/core/device/interfaces"
	"time"
)

//NewCanceledElectricLockTask ..
func NewCanceledElectricLockTask(_task interfaces.ITask, description string) *CanceledElectricLockTask {
	return &CanceledElectricLockTask{
		task:                _task,
		canceledTime:        time.Now().UTC(),
		canseledDescription: description,
	}
}

//CanceledElectricLockTask ...
type CanceledElectricLockTask struct {
	ElectricLockTask
	task                interfaces.ITask
	canceledTime        time.Time
	canseledDescription string
}
