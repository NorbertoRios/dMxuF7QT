package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
	"genx-go/logger"
)

//NewAnyMessageObserver ...
func NewAnyMessageObserver(_task interfaces.ITask, _command string) *AnyMessageObserver {
	return &AnyMessageObserver{
		task:    _task,
		command: _command,
	}
}

//AnyMessageObserver ...
type AnyMessageObserver struct {
	task    interfaces.ITask
	command string
}

//Update ...
func (observer *AnyMessageObserver) Update(msg interface{}) *list.List {
	commands := list.New()
	commands.PushBack(NewSendConfigCommand(observer.task, observer.command))
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
