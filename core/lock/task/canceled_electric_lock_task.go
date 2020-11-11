package device

import "time"

//NewCanceledElectricLockTask ..
func NewCanceledElectricLockTask(_task *ElectricLockTask, description string) *CanceledElectricLockTask {
	return &CanceledElectricLockTask{
		task:                _task,
		canceledTime:        time.Now().UTC(),
		canseledDescription: description,
	}
}

//CanceledElectricLockTask ...
type CanceledElectricLockTask struct {
	task                *ElectricLockTask
	canceledTime        time.Time
	canseledDescription string
}
