package device

import (
	"container/list"
	"genx-go/logger"
	"genx-go/message"
)

//NewWaitingelEctricLockAck ..
func NewWaitingelEctricLockAck(_task *ElectricLockTask) *WaitingelEctricLockAck {
	return &WaitingelEctricLockAck{
		task: _task,
	}
}

//WaitingelEctricLockAck ..
type WaitingelEctricLockAck struct {
	task *ElectricLockTask
}

//Attached ..
func (observer *WaitingelEctricLockAck) Attached() {
	logger.Logger().WriteToLog(logger.Info, "[WaitingelEctricLockAck] Successfuly attached")
}

//Task returns observer's task
func (observer *WaitingelEctricLockAck) Task() ITask {
	return observer.task
}

//Update ...
func (observer *WaitingelEctricLockAck) Update(msg interface{}) *list.List {
	commands := list.New()
	switch msg.(type) {
	case *message.AckMessage:
		{
			ackMessage := msg.(*message.AckMessage)
			setRelayDrive := NewElectricLockSetRelayDrive(observer.task.Request)
			if ackMessage.Value == setRelayDrive.Command() {
				commands := list.New()
				commands.PushBack(NewDetachObserverCommand(observer))
				observer.task.Done()
			}
		}
	}
	return commands
}
