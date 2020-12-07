package interfaces

import (
	"container/list"
	"genx-go/core/immobilizer/request"
)

//IImmobilizer ...
type IImmobilizer interface {
	Trigger() string
	NewRequest(*request.ChangeImmoStateRequest) *list.List
	State() string
	CurrentTask() ITask
	Tasks() *list.List
}
