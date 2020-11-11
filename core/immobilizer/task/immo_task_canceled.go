package task

import "time"

//NewCanceledImmoTask ..
func NewCanceledImmoTask(_task *ImmobilizerTask, description string) *CanceledImmoTask {
	return &CanceledImmoTask{
		task:                _task,
		canceledTime:        time.Now().UTC(),
		canseledDescription: description,
	}
}

//CanceledImmoTask ...
type CanceledImmoTask struct {
	task                *ImmobilizerTask
	canceledTime        time.Time
	canseledDescription string
}
