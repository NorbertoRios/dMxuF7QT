package interfaces

import (
	"container/list"
	"genx-go/core/request"
)

//IProcess ...
type IProcess interface {
	TaskCancel(ITask, string)
	TaskDone(ITask)
	CurrentTask() ITask
	Tasks() *list.List
	NewRequest(request.IRequest) *list.List
}
