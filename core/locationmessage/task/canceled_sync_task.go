package task

import (
	"genx-go/core/device/interfaces"
	"time"
)

//NewCanceledSyncTask ...
func NewCanceledSyncTask(_task interfaces.ITask, _description string) *CanceledSyncTask {
	return &CanceledSyncTask{
		task:                _task,
		canceledTime:        time.Now().UTC(),
		canseledDescription: _description,
	}
}

//CanceledSyncTask ...
type CanceledSyncTask struct {
	SyncTask
	task                interfaces.ITask
	canceledTime        time.Time
	canseledDescription string
}
