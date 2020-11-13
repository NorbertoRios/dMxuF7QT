package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/logger"
)

//NewSendStringCommand ...
func NewSendStringCommand(_command string) *SendStringCommand {
	return &SendStringCommand{
		command: _command,
	}
}

//SendStringCommand ...
type SendStringCommand struct {
	command string
}

//Execute ...
func (c *SendStringCommand) Execute(device interfaces.IDevice) *list.List {
	commands := list.New()
	if err := device.Send(c.command); err != nil {
		logger.Logger().WriteToLog(logger.Error, "[SendStringCommand | Execute] Error while sending command ", c.command)
	}
	//commands.PushBack(NewPushToRabbitMessageCommand(c.command, Message))
	return commands
}
