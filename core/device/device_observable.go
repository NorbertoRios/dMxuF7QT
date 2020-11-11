package device

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"sync"
)

//NewObservable return new observable
func NewObservable() *Observable {
	return &Observable{
		mutex:     &sync.Mutex{},
		observers: make(map[interfaces.IObserver]bool, 0),
	}
}

//Observable ...
type Observable struct {
	mutex     *sync.Mutex
	observers map[interfaces.IObserver]bool
}

//Observers return all observers
func (o *Observable) Observers() []interfaces.IObserver {
	observers := make([]interfaces.IObserver, 0)
	o.mutex.Lock()
	defer o.mutex.Unlock()
	for observer := range o.observers {
		observers = append(observers, observer)
	}
	return observers
}

//Attach observer
func (o *Observable) Attach(observer interfaces.IObserver) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.observers[observer] = true
	observer.Attached()
}

//Detach observer
func (o *Observable) Detach(observer interfaces.IObserver) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	delete(o.observers, observer)
}

//Notify observers
func (o *Observable) Notify(msg interface{}) *list.List {
	commands := list.New()
	for _, observer := range o.Observers() {
		command := observer.Update(msg)
		if command == nil {
			continue
		} else {
			commands.PushBackList(command)
		}

	}
	return commands
}
