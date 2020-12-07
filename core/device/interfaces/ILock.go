package interfaces

import (
	"container/list"
	"genx-go/core/lock/request"
)

//ILock ...
type ILock interface {
	NewRequest(*request.UnlockRequest) *list.List
	CurrentTask() ITask
	Tasks() *list.List
}
