package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/logger"
)

//NewImmoSendDiagCommand ...
func NewImmoSendDiagCommand(_command string) *ImmoSendDiagCommand {
	return &ImmoSendDiagCommand{
		command: _command,
	}
}

//ImmoSendDiagCommand ...
type ImmoSendDiagCommand struct {
	command string
}

//Execute ...
func (c *ImmoSendDiagCommand) Execute(device interfaces.IDevice) *list.List {
	commands := list.New()
	if err := device.Send(c.command); err != nil {
		logger.Logger().WriteToLog(logger.Error, "[ImmoSendDiagCommand | Execute] Error while sending command ", c.command)
	}
	//commands.PushBack(NewPushToRabbitMessageCommand(c.command, Message))
	return commands
}
