package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
	"genx-go/logger"
)

//NewAnyImmoDiagObserver ...
func NewAnyImmoDiagObserver(_task interfaces.ITask) *AnyImmoDiagObserver {
	return &AnyImmoDiagObserver{
		task: _task,
	}
}

//AnyImmoDiagObserver ...
type AnyImmoDiagObserver struct {
	task interfaces.ITask
}

//Update ...
func (observer *AnyImmoDiagObserver) Update(msg interface{}) *list.List {
	cList := list.New()
	cList.PushBack(observers.NewSendStringCommand("DIAG HARDWARE"))
	cList.PushBack(observers.NewDetachObserverCommand(observer))
	return cList
}

//Attached ...
func (observer *AnyImmoDiagObserver) Attached() {
	logger.Logger().WriteToLog(logger.Info, "[AnyMessageObserver] Successfuly attached")
}

//Task ...
func (observer *AnyImmoDiagObserver) Task() interfaces.ITask {
	return observer.task
}
