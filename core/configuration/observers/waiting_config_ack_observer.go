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
		watchdog: watchdog.NewWatchdog(_task.Device(), _task.Invoker().(interfaces.IConfigInvoker).SendConfigAfterAnyMessage(_task), 10),
	}
}

//WaitingConfigAckObserver ..
type WaitingConfigAckObserver struct {
	task     interfaces.IConfigTask
	watchdog *watchdog.Watchdog
}

//Update ...
func (observer *WaitingConfigAckObserver) Update(msg interface{}) *list.List {
	commands := list.New()
	switch msg.(type) {
	case *message.AckMessage:
		{
			ackMessage := msg.(*message.AckMessage)
			if ackMessage.Value == observer.task.CurrentStringCommand() {
				observer.watchdog.Stop()
				commands.PushBack(observers.NewDetachObserverCommand(observer))
				observer.task.(interfaces.IConfigTask).CommandComplete()
				if observer.task.IsNextExist() {
					observer.task.GoToNextCommand()
					commands.PushBackList(observer.task.Commands())
				} else {
					commands.PushBackList(observer.task.Invoker().DoneTask(observer.task))
				}
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
