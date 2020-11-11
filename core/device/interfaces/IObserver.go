package interfaces

import "container/list"

//IObserver base observer
type IObserver interface {
	Update(interface{}) *list.List
	Task() ITask
	Attached()
}
