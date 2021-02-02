package interfaces

import (
	"container/list"
)

//IProcess ...
type IProcess interface {
	TaskCancel(ITask, string)
	TaskDone(ITask)
	CurrentTask() ITask
	Tasks() *list.List
	NewRequest(interface{}, IDevice) *list.List
	Pause()
	Resume()
}
