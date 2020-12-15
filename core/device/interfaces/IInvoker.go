package interfaces

import (
	"container/list"
)

//IInvoker ..
type IInvoker interface {
	CanselTask(ITask, string) *list.List
	DoneTask(ITask) *list.List
}
