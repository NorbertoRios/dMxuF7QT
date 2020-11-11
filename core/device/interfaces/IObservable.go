package interfaces

import "container/list"

//IObservable interface ...
type IObservable interface {
	Attach(IObserver)
	Detach(IObserver)
	Notify(interface{}) *list.List
	Observers() []IObserver
}
