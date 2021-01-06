package filter

import (
	"genx-go/core/device/interfaces"
	"reflect"
)

//NewObserversFilter new filter
func NewObserversFilter(_observable interfaces.IObservable) *ObserversFilter {
	return &ObserversFilter{
		observable: _observable,
	}
}

//ObserversFilter ...
type ObserversFilter struct {
	observable interfaces.IObservable
}

//Extract observers by task
func (filter *ObserversFilter) Extract(task interfaces.ITask) []interfaces.IObserver {
	observers := make([]interfaces.IObserver, 0)
	for _, observer := range filter.observable.Observers() {
		if reflect.DeepEqual(observer.Task(), task) {
			observers = append(observers, observer)
		}
	}
	return observers
}
