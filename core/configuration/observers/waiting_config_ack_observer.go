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
func NewWaitingConfigAckObserver(_task interfaces.IConfigTask) *WaitingConfigAckObserver {
	return &WaitingConfigAckObserver{
		task:     _task,
		watchdog: watchdog.NewConfigWatchdog(_task, 10),
	}
}

//WaitingConfigAckObserver ..
type WaitingConfigAckObserver struct {
	task     interfaces.IConfigTask
	watchdog *watchdog.ConfigWatchdog
}

//Update ...
func (observer *WaitingConfigAckObserver) Update(msg interface{}) *list.List {
	commands := list.New()
	switch msg.(type) {
	case *message.AckMessage:
		{
			ackMessage := msg.(*message.AckMessage)
			if ackMessage.Value == observer.task.CurrentCommand() {
				observer.watchdog.Stop()
				commands.PushBack(observers.NewDetachObserverCommand(observer))
				commands.PushBackList(observer.task.(interfaces.IConfigTask).NextStep())
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
