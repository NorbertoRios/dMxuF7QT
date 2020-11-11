package interfaces

import (
	"container/list"
	"genx-go/core/immobilizer/request"
)

//IImmobilizer ...
type IImmobilizer interface {
	Trigger() string
	NewRequest(*request.ChangeImmoStateRequest)
	State() string
	CurrentTask() ITask
	Tasks() *list.List
}
