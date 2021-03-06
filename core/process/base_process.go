package process

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"sync"
)

//BaseProcess ...
type BaseProcess struct {
	Mutex              *sync.Mutex
	ProcessCurrentTask interfaces.ITask
	ProcessTasks       *list.List
}

//CurrentTask ...
func (process *BaseProcess) CurrentTask() interfaces.ITask {
	return process.ProcessCurrentTask
}

//Tasks ...
func (process *BaseProcess) Tasks() *list.List {
	return process.ProcessTasks
}

//PushToTasks ...
func (process *BaseProcess) PushToTasks(_task interfaces.ITask, isDone bool) {
	process.Mutex.Lock()
	defer process.Mutex.Unlock()
	if isDone {
		process.ProcessTasks.PushFront(_task)
	} else {
		process.ProcessTasks.PushBack(_task)
	}
}

//Pause ...
func (process *BaseProcess) Pause() {
}

//Resume ...
func (process *BaseProcess) Resume() {
}
