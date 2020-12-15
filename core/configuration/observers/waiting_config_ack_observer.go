package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
	"genx-go/logger"
	"genx-go/message"
)

//NewWaitingConfigAckObserver ...
func NewWaitingConfigAckObserver(_task interfaces.ITask, _command string) *WaitingConfigAckObserver {
	observer := &WaitingConfigAckObserver{
		task:    _task,
		command: _command,
	}
	wdList := list.New()
	wdList.PushBack(observers.NewDetachObserverCommand(observer))
	wdList.PushBack(observers.NewAttachObserverCommand(NewAnyMessageObserver(_task, _command)))
	//wd := watchdog.NewWatchdog(wdList, observer.task.Device(), 5)
	//observer.watchdog = wd
	return observer
}

//WaitingConfigAckObserver ..
type WaitingConfigAckObserver struct {
	task interfaces.ITask
	//watchdog *watchdog.Watchdog
	command string
}

//Update ...
func (observer *WaitingConfigAckObserver) Update(msg interface{}) *list.List {
	commands := list.New()
	switch msg.(type) {
	case *message.AckMessage:
		{
			ackMessage := msg.(*message.AckMessage)
			if ackMessage.Value == observer.command {
				//observer.watchdog.Stop()
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
	//observer.watchdog.Start()
	logger.Logger().WriteToLog(logger.Info, "[WaitingConfigAckObserver] Successfuly attached")
}
