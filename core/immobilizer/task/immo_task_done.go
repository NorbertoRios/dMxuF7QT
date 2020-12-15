package task

import (
	"genx-go/core/device/interfaces"
	"time"
)

//NewDoneImmoTask ..
func NewDoneImmoTask(_task interfaces.ITask) *DoneImmoTask {
	return &DoneImmoTask{
		Task:     _task,
		doneTime: time.Now().UTC(),
	}
}

//DoneImmoTask done immovilizer task
type DoneImmoTask struct {
	ImmobilizerTask
	Task     interfaces.ITask
	doneTime time.Time
}
