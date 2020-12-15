package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/lock/request"
	"genx-go/core/observers"
	"genx-go/core/watchdog"
	"genx-go/logger"
	"genx-go/message"
)

//NewWaitingEctricLockAck ..
func NewWaitingEctricLockAck(_task interfaces.ITask) *WaitingEctricLockAck {
	return &WaitingEctricLockAck{
		task: _task,
		wd:   watchdog.NewElectricLockWatchdog(_task),
	}
}

//WaitingEctricLockAck ..
type WaitingEctricLockAck struct {
	task interfaces.ITask
	wd   *watchdog.ElectricLockWatchdog
}

//Attached ..
func (observer *WaitingEctricLockAck) Attached() {
	observer.wd.Start()
	logger.Logger().WriteToLog(logger.Info, "[WaitingEctricLockAck] Successfuly attached")
}

//Task returns observer's task
func (observer *WaitingEctricLockAck) Task() interfaces.ITask {
	return observer.task
}

//Update ...
func (observer *WaitingEctricLockAck) Update(msg interface{}) *list.List {
	setRelayDrive := NewElectricLockSetRelayDrive(observer.task.Request().(*request.UnlockRequest))
	commands := list.New()
	switch msg.(type) {
	case *message.AckMessage:
		{
			ackMessage := msg.(*message.AckMessage)
			if ackMessage.Value == setRelayDrive.Command() {
				observer.wd.Stop()
				return observer.task.Invoker().DoneTask(observer.task)
			}
			commands.PushBack(observers.NewSendStringCommand(setRelayDrive.Command()))
		}
	default:
		{
			commands.PushBack(observers.NewSendStringCommand(setRelayDrive.Command()))
		}
	}
	return commands
}
