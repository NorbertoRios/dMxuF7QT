package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
	"genx-go/logger"
)

//NewAnyMessageObserver ...
func NewAnyMessageObserver(_task interfaces.ITask) *AnyMessageObserver {
	return &AnyMessageObserver{
		task: _task,
	}
}

//AnyMessageObserver ...
type AnyMessageObserver struct {
	task interfaces.ITask
}

//Update ...
func (observer *AnyMessageObserver) Update(msg interface{}) *list.List {
	commands := list.New()
	commands.PushBack(NewSendDiagCommand(observer.task))
	commands.PushBack(observers.NewDetachObserverCommand(observer))
	return commands
}

//Task ...
func (observer *AnyMessageObserver) Task() interfaces.ITask {
	return observer.task.(interfaces.ITask)
}

//Attached ...
func (observer *AnyMessageObserver) Attached() {
	logger.Logger().WriteToLog(logger.Info, "[Location process | AnyMessageObserver] Successfuly attached")
}
