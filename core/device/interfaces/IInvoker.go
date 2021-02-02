package interfaces

import (
	"container/list"
)

//IInvoker ..
type IInvoker interface {
	CancelTask(ITask, string) *list.List
	DoneTask(ITask) *list.List
}
