package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/immobilizer/request"
	"genx-go/core/observers"
	"genx-go/core/watchdog"
	"genx-go/logger"
	"genx-go/message"
)

//NewWaitingImmoAckObserver ...
func NewWaitingImmoAckObserver(_task interfaces.ITask) *WaitingImmoAckObserver {
	observer := &WaitingImmoAckObserver{
		task: _task,
	}
	return observer
}

//Attached ..
func (observer *WaitingImmoAckObserver) Attached() {
	wdList := list.New()
	anyMessageObserver := NewAnyImmoMessageObserver(observer.task)
	wdList.PushBack(observers.NewDetachObserverCommand(observer))
	wdList.PushBack(observers.NewAttachObserverCommand(anyMessageObserver))
	wd := watchdog.NewWatchdog(wdList, observer.task.Device(), 5)
	observer.Watchdog = wd
	logger.Logger().WriteToLog(logger.Info, "[WaitingImmoAckObserver] Successfuly attached")
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
	commands := list.New()
	switch msg.(type) {
	case *message.AckMessage:
		{
			ackMessage := msg.(*message.AckMessage)
			setRelayDrive := NewSetRelayDrive(observer.task.Request().(*request.ChangeImmoStateRequest))
			if ackMessage.Value == setRelayDrive.Command() {
				observer.Watchdog.Stop()
				immoConfObserver := NewImmoConfitmationObserver(observer.task)
				commands.PushBack(observers.NewDetachObserverCommand(observer))
				commands.PushBack(observers.NewAttachObserverCommand(immoConfObserver))
				commands.PushBack(NewImmoSendDiagCommand("DIAG HARDWARE"))
			}
		}
	}
	return commands
}
