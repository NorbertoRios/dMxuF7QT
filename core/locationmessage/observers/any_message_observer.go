package observers

import (
	"container/list"
	"genx-go/core/commands"
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
	cmd := list.New()
	cmd.PushBack(commands.NewSendDeviceMessageCommand("DIAG PARAMS=24"))
	cmd.PushBack(observers.NewDetachObserverCommand(observer))
	return cmd
}

//Task ...
func (observer *AnyMessageObserver) Task() interfaces.ITask {
	return observer.task.(interfaces.ITask)
}

//Attached ...
func (observer *AnyMessageObserver) Attached() {
	logger.Logger().WriteToLog(logger.Info, "[Location process | AnyMessageObserver] Successfuly attached")
}
