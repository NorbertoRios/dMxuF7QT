package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
	"genx-go/logger"
	"genx-go/message"
)

//NewLocationAnyMessageObserver ...
func NewLocationAnyMessageObserver(_task interfaces.ITask) *LocationAnyMessageObserver {
	return &LocationAnyMessageObserver{
		task: _task,
	}
}

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
	logger.Logger().WriteToLog(logger.Info, "[LocationAnyMessageObserver | Attached] Successfully attached")
}

//Update ...
func (observer *LocationAnyMessageObserver) Update(msg interface{}) *list.List {
	cList := list.New()
	cList.PushFront(observers.NewDetachObserverCommand(observer))
	if _, f := msg.(*message.LocationMessage); f {
		observer.task.Invoker().DoneTask(observer.task)
	} else {
		cList.PushBack(NewSendLocationRequest(observer.task))
	}
	return cList
}
