package interfaces

import (
	"container/list"
	"genx-go/core/request"
)

//IInvoker ..
type IInvoker interface {
	NewImmoRequest(req request.IRequest) *list.List
	CanselTask(ITask, string) *list.List
	DoneTask(ITask) *list.List
}
