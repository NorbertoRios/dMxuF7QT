package interfaces

import "container/list"

//ITask itask interface
type ITask interface {
	Commands() *list.List
	Observers() []IObserver
	Device() IDevice
	Cancel(string)
	Done()
	Request() interface{}
}
