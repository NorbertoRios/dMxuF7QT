package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	coreObservers "genx-go/core/observers"
	"genx-go/logger"
)

//NewSendConfigCommand ...
func NewSendConfigCommand(_task interfaces.ITask, _command string) *SendConfigCommand {
	return &SendConfigCommand{
		task:    _task,
		command: _command,
	}
}

//SendConfigCommand ..
type SendConfigCommand struct {
	task    interfaces.ITask
	command string
}

//Execute ...
func (c *SendConfigCommand) Execute(device interfaces.IDevice) *list.List {
	commands := list.New()
	if err := device.Send(c.command); err != nil {
		logger.Logger().WriteToLog(logger.Error, "[ImmoSendRelayCommand | Execute] Error while sending command ", c.command)
	}
	commands.PushBack(coreObservers.NewAttachObserverCommand(NewWaitingConfigAckObserver(c.task, c.command)))
	return commands
}
