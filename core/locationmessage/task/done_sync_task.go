package task

import (
	"genx-go/core/device/interfaces"
	"time"
)

//NewDoneSyncTask ...
func NewDoneSyncTask(_task interfaces.ITask) *DoneSyncTask {
	return &DoneSyncTask{
		task:     _task,
		doneTime: time.Now().UTC(),
	}
}

//DoneSyncTask ...
type DoneSyncTask struct {
	SyncTask
	task     interfaces.ITask
	doneTime time.Time
}
