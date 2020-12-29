package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	coreObservers "genx-go/core/observers"
	"genx-go/logger"
)

//NewSendConfigCommand ...
func NewSendConfigCommand(_task interfaces.IConfigTask) *SendConfigCommand {
	return &SendConfigCommand{
		task: _task,
	}
}

//SendConfigCommand ..
type SendConfigCommand struct {
	task interfaces.IConfigTask
}

//Execute ...
func (c *SendConfigCommand) Execute(device interfaces.IDevice) *list.List {
	commands := list.New()
	if err := device.Send(c.task.CurrentCommand()); err != nil {
		logger.Logger().WriteToLog(logger.Error, "[ImmoSendRelayCommand | Execute] Error while sending command ", c.task.CurrentCommand())
	}
	commands.PushBack(coreObservers.NewAttachObserverCommand(NewWaitingConfigAckObserver(c.task)))
	return commands
}
