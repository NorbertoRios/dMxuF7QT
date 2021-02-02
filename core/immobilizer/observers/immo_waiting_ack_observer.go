package observers

import (
	"container/list"
	"genx-go/core/commands"
	"genx-go/core/device/interfaces"
	"genx-go/core/immobilizer/request"
	"genx-go/core/observers"
	"genx-go/core/watchdog"
	"genx-go/logger"
	"genx-go/message"
)

//NewWaitingImmoAckObserver ...
func NewWaitingImmoAckObserver(_task interfaces.ITask) *WaitingImmoAckObserver {
	setRelayDrive := NewSetRelayDrive(_task.Request().(*request.ChangeImmoStateRequest))
	return &WaitingImmoAckObserver{
		task:     _task,
		Watchdog: watchdog.NewWatchdog(_task.Device(), _task.Invoker().(interfaces.IImmoInvoker).WatchdogsCommands(_task, setRelayDrive.Command()), 30),
	}
}

//Attached ..
func (observer *WaitingImmoAckObserver) Attached() {
	logger.Logger().WriteToLog(logger.Info, "[WaitingImmoAckObserver] Successfuly attached")
	observer.Watchdog.Start()
}

//WaitingImmoAckObserver ...
type WaitingImmoAckObserver struct {
	task     interfaces.ITask
	Watchdog *watchdog.Watchdog
}

//Task returns observer's task
func (observer *WaitingImmoAckObserver) Task() interfaces.ITask {
	return observer.task
}

//Update observer
func (observer *WaitingImmoAckObserver) Update(msg interface{}) *list.List {
	cmds := list.New()
	switch msg.(type) {
	case *message.AckMessage:
		{
			ackMessage := msg.(*message.AckMessage)
			setRelayDrive := NewSetRelayDrive(observer.task.Request().(*request.ChangeImmoStateRequest))
			if ackMessage.Value == setRelayDrive.Command() {
				observer.Watchdog.Stop()
				immoConfObserver := NewImmoConfitmationObserver(observer.task)
				cmds.PushBack(observers.NewDetachObserverCommand(observer))
				cmds.PushBack(observers.NewAttachObserverCommand(immoConfObserver))
				cmds.PushBack(commands.NewSendDeviceMessageCommand("DIAG HARDWARE"))
			}
		}
	}
	return cmds
}
