package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/lock/request"
	"genx-go/core/observers"
	"genx-go/logger"
	"time"
)

//NewElectricLockSendCommand ...
func NewElectricLockSendCommand(_task interfaces.ITask) interfaces.ICommand {
	return &ElectricLockSendCommand{
		task: _task,
	}
}

//ElectricLockSendCommand ...
type ElectricLockSendCommand struct {
	task interfaces.ITask
}

//Execute ...
func (c *ElectricLockSendCommand) Execute(device interfaces.IDevice) *list.List {
	req := c.task.Request().(*request.UnlockRequest)
	commands := list.New()
	setRelayDrive := NewElectricLockSetRelayDrive(req)
	if req.Time().Before(time.Now().UTC()) {
		logger.Logger().WriteToLog(logger.Info, "[ElectricLockSendCommand | Execute] Time is over. ExpirationTime: ", req.Time().String(), ". Current time: ", time.Now().UTC().String())
		return c.task.Invoker().CancelTask(c.task, "Time is over")
	}
	if err := device.Send(setRelayDrive.Command()); err != nil {
		logger.Logger().WriteToLog(logger.Error, "[ElectricLockSendCommand | Execute] Error while sending command ", setRelayDrive.Command())
	}
	commands.PushBack(observers.NewAttachObserverCommand(NewWaitingEctricLockAck(c.task)))
	return commands
}
