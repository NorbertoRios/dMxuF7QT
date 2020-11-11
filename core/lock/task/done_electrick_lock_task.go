package device

import "time"

//DoneElectricLockTask ...
type DoneElectricLockTask struct {
	Task     *ElectricLockTask
	doneTime time.Time
}

//NewDoneElectricLockTask ...
func NewDoneElectricLockTask(_task *ElectricLockTask) *DoneElectricLockTask {
	return &DoneElectricLockTask{
		Task:     _task,
		doneTime: time.Now().UTC(),
	}
}
