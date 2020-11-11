package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
	"genx-go/logger"
)

//NewAnyImmoMessageObserver ...
func NewAnyImmoMessageObserver(_task interfaces.ITask) *AnyMessageObserver {
	return &AnyMessageObserver{
		task: _task,
	}
}

//AnyMessageObserver ...
type AnyMessageObserver struct {
	task interfaces.ITask
}

//Attached ...
func (observer *AnyMessageObserver) Attached() {
	logger.Logger().WriteToLog(logger.Info, "[AnyMessageObserver] Successfuly attached")
}

//Task ...
func (observer *AnyMessageObserver) Task() interfaces.ITask {
	return observer.task
}

//Update ...
func (observer *AnyMessageObserver) Update(msg interface{}) *list.List {
	commands := list.New()
	commands.PushBack(NewImmoSendRelayCommand(observer.task))
	commands.PushBack(observers.NewDetachObserverCommand(observer))
	return commands
}
