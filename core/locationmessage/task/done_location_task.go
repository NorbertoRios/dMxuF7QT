package task

import (
	"genx-go/core/device/interfaces"
	"time"
)

//DoneLocationMessageTask ...
func NewDoneLocationMessageTask(_task interfaces.ITask) *DoneLocationMessageTask {
	return &DoneLocationMessageTask{
		task:     _task,
		doneTime: time.Now().UTC(),
	}
}

//DoneLocationMessageTask ...
type DoneLocationMessageTask struct {
	LocationMessageTask
	task     interfaces.ITask
	doneTime time.Time
}
