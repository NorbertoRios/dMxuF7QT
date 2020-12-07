package interfaces

import (
	"container/list"
	"genx-go/core/request"
)

//ILocationRequest ...
type ILocationRequest interface {
	NewRequest(*request.BaseRequest) *list.List
	Task() ITask
}
