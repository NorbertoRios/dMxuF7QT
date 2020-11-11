package commands

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/lock/task"
	"genx-go/core/observers"
	"genx-go/logger"
)

//ElectricLockSendCommand ...
type ElectricLockSendCommand struct {
	task *task.ElectricLockTask
}

//Execute ...
func (c *ElectricLockSendCommand) Execute(device interfaces.IDevice) *list.List {
	commands := list.New()
	setRelayDrive := NewElectricLockSetRelayDrive(c.task.Request)
	if err := device.Send(setRelayDrive.Command()); err != nil {
		logger.Logger().WriteToLog(logger.Error, "[ElectricLockSendCommand | Execute] Error while sending command ", setRelayDrive.Command())
	}
	commands.PushBack(observers.NewAttachObserverCommand(NewWaitingelEctricLockAck(c.task)))
	commands.PushBack(NewPushToRabbitMessageCommand(setRelayDrive.Command(), Message))
	return commands
}
