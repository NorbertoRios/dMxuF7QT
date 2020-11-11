package task

import (
	"time"
)

//NewDoneImmoTask ..
func NewDoneImmoTask(_task *ImmobilizerTask) *DoneImmoTask {
	return &DoneImmoTask{
		Task:     _task,
		doneTime: time.Now().UTC(),
	}
}

//DoneImmoTask done immovilizer task
type DoneImmoTask struct {
	Task     *ImmobilizerTask
	doneTime time.Time
}
