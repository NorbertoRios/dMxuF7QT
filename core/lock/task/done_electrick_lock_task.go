package task

import "time"

//NewDoneElectricLockTask ...
func NewDoneElectricLockTask(_task *ElectricLockTask) *DoneElectricLockTask {
	return &DoneElectricLockTask{
		Task:     _task,
		doneTime: time.Now().UTC(),
	}
}

//DoneElectricLockTask ...
type DoneElectricLockTask struct {
	ElectricLockTask
	Task     *ElectricLockTask
	doneTime time.Time
}
