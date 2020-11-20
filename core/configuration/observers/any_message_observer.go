package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
	"genx-go/logger"
)

//NewAnyMessageObserver ...
func NewAnyMessageObserver(_task interfaces.ITask) *AnyMessageObserver {
	return &AnyMessageObserver{}
}

//AnyMessageObserver ...
type AnyMessageObserver struct {
	task interfaces.ITask
}

//Update ...
func (observer *AnyMessageObserver) Update(msg interface{}) *list.List {
	commands := list.New()
	commands.PushBack(NewSendConfigCommand(observer.task))
	commands.PushBack(observers.NewDetachObserverCommand(observer))
	return commands
}

//Task ...
func (observer *AnyMessageObserver) Task() interfaces.ITask {
	return observer.task
}

//Attached ...
func (observer *AnyMessageObserver) Attached() {
	logger.Logger().WriteToLog(logger.Info, "[AnyMessageObserver] Successfuly attached")
}
