package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/immobilizer/request"
	"genx-go/core/observers"
	"genx-go/logger"
)

//NewImmoSendRelayCommand ...
func NewImmoSendRelayCommand(_task interfaces.ITask) *ImmoSendRelayCommand {
	return &ImmoSendRelayCommand{
		task: _task,
	}
}

//ImmoSendRelayCommand ..
type ImmoSendRelayCommand struct {
	task interfaces.ITask
}

//Execute ..
func (c *ImmoSendRelayCommand) Execute(device interfaces.IDevice) *list.List {
	commands := list.New()
	setRelayDrive := NewSetRelayDrive(c.task.Request().(*request.ChangeImmoStateRequest))
	if err := device.Send(setRelayDrive.Command()); err != nil {
		logger.Logger().WriteToLog(logger.Error, "[ImmoSendRelayCommand | Execute] Error while sending command ", setRelayDrive.Command())
	}
	commands.PushBack(observers.NewAttachObserverCommand(NewWaitingImmoAckObserver(c.task)))
	return commands
}
