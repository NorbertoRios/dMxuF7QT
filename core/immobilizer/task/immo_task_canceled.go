package task

import (
	"genx-go/core/device/interfaces"
	"time"
)

//NewCanceledImmoTask ..
func NewCanceledImmoTask(_task interfaces.ITask, description string) *CanceledImmoTask {
	return &CanceledImmoTask{
		task:                _task,
		canceledTime:        time.Now().UTC(),
		canseledDescription: description,
	}
}

//CanceledImmoTask ...
type CanceledImmoTask struct {
	ImmobilizerTask
	task                interfaces.ITask
	canceledTime        time.Time
	canseledDescription string
}
