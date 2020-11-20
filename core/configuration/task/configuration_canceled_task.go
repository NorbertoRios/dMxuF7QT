package task

import "time"

//NewCanceledConfigTask ..
func NewCanceledConfigTask(_task *ConfigTask, _description string) *CanceledConfigTask {
	return &CanceledConfigTask{
		Task:                _task,
		canceledTime:        time.Now().UTC(),
		canseledDescription: _description,
	}
}

//CanceledConfigTask ...
type CanceledConfigTask struct {
	Task                *ConfigTask
	canceledTime        time.Time
	canseledDescription string
}
