package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
	"genx-go/message"
)

//LocationAnyMessageObserver ...
type LocationAnyMessageObserver struct {
	task interfaces.ITask
}

//Task ..
func (observer *LocationAnyMessageObserver) Task() interfaces.ITask {
	return observer.task
}

//Attached ..
func (observer *LocationAnyMessageObserver) Attached() {

}

//Update ...
func (observer *LocationAnyMessageObserver) Update(msg interface{}) *list.List {
	cList := list.New()
	cList.PushFront(observers.NewDetachObserverCommand(observer))
	if _, f := msg.(*message.LocationMessage); f {
		observer.task.Done()
	} else {
		cList.PushBack(NewSendLocationRequest(observer.task))
	}
	return cList
}
