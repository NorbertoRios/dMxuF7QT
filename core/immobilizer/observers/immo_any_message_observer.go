package observers

import (
	"container/list"
	"genx-go/core/commands"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
	"genx-go/logger"
)

//NewAnyImmoMessageObserver ...
func NewAnyImmoMessageObserver(_task interfaces.ITask, _message string) *AnyMessageObserver {
	return &AnyMessageObserver{
		task:    _task,
		message: _message,
	}
}

//AnyMessageObserver ...
type AnyMessageObserver struct {
	message string
	task    interfaces.ITask
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
	cmds := list.New()
	cmds.PushBack(commands.NewSendDeviceMessageCommand(observer.message))
	cmds.PushBack(observers.NewDetachObserverCommand(observer))
	return cmds
}
