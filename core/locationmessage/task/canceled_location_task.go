package task

import (
	"genx-go/core/device/interfaces"
	"time"
)

//NewCanceledLocationMessageTask ...
func NewCanceledLocationMessageTask(_task interfaces.ITask, _description string) *CanceledLocationMessageTask {
	return &CanceledLocationMessageTask{
		task:                _task,
		canceledTime:        time.Now().UTC(),
		canseledDescription: _description,
	}
}

//CanceledLocationMessageTask ...
type CanceledLocationMessageTask struct {
	LocationMessageTask
	task                interfaces.ITask
	canceledTime        time.Time
	canseledDescription string
}
