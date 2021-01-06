package task

import (
	"genx-go/core/device/interfaces"
	"time"
)

//NewDoneConfigTask ..
func NewDoneConfigTask(_task interfaces.ITask) interfaces.ITask {
	return &DoneConfigTask{
		Task:     _task,
		doneTime: time.Now().UTC(),
	}
}

//DoneConfigTask done immovilizer task
type DoneConfigTask struct {
	ConfigTask
	Task     interfaces.ITask
	doneTime time.Time
}
