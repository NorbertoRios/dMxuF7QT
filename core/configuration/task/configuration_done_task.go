package task

import "time"

//NewDoneConfigTask ..
func NewDoneConfigTask(_task *ConfigTask) *DoneConfigTask {
	return &DoneConfigTask{
		Task:     _task,
		doneTime: time.Now().UTC(),
	}
}

//DoneConfigTask done immovilizer task
type DoneConfigTask struct {
	Task     *ConfigTask
	doneTime time.Time
}
