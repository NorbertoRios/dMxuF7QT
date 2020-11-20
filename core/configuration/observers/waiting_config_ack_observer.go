package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
	"genx-go/core/watchdog"
	"genx-go/logger"
	"genx-go/message"
)

//NewWaitingConfigAckObserver ...
func NewWaitingConfigAckObserver(_task interfaces.ITask) *WaitingConfigAckObserver {
	observer := &WaitingConfigAckObserver{
		task: _task,
	}
	anyMessageObserver := NewAnyMessageObserver(_task)
	wdList := list.New()
	wdList.PushBack(observers.NewDetachObserverCommand(observer))
	wdList.PushBack(observers.NewAttachObserverCommand(anyMessageObserver))
	wd := watchdog.NewWatchdog(wdList, observer.task.Device(), 5)
	observer.watchdog = wd
	return observer
}

//WaitingConfigAckObserver ..
type WaitingConfigAckObserver struct {
	task     interfaces.ITask
	watchdog *watchdog.Watchdog
}

//Update ...
func (observer *WaitingConfigAckObserver) Update(msg interface{}) *list.List {
	commands := list.New()
	switch msg.(type) {
	case *message.AckMessage:
		{
			ackMessage := msg.(*message.AckMessage)
			command := NewConfig(observer.task.Request().(string))
			if ackMessage.Value == command.Command() {
				go observer.watchdog.Stop()
				commands.PushBack(observers.NewDetachObserverCommand(observer))
				observer.task.Done()
			}
		}
	}
	return commands
}

//Task ...
func (observer *WaitingConfigAckObserver) Task() interfaces.ITask {
	return observer.task
}

//Attached ...
func (observer *WaitingConfigAckObserver) Attached() {
	observer.watchdog.Start()
	logger.Logger().WriteToLog(logger.Info, "[WaitingConfigAckObserver] Successfuly attached")
}
