package task

import (
	"genx-go/core/device/interfaces"
	"time"
)

//NewCanceledConfigTask ..
func NewCanceledConfigTask(_task interfaces.ITask, _description string) interfaces.ITask {
	return &CanceledConfigTask{
		Task:                _task,
		canceledTime:        time.Now().UTC(),
		canseledDescription: _description,
	}
}

//CanceledConfigTask ...
type CanceledConfigTask struct {
	ConfigTask
	Task                interfaces.ITask
	canceledTime        time.Time
	canseledDescription string
}
