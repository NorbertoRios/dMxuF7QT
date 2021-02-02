package commands

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/logger"
)

//NewSendDeviceMessageCommand ...
func NewSendDeviceMessageCommand(_command string) *SendDeviceMessageCommand {
	return &SendDeviceMessageCommand{
		command: _command,
	}
}

//SendDeviceMessageCommand ...
type SendDeviceMessageCommand struct {
	command string
}

//Execute ...
func (c *SendDeviceMessageCommand) Execute(device interfaces.IDevice) *list.List {
	if err := device.Send(c.command); err != nil {
		logger.Logger().WriteToLog(logger.Error, "[SendDeviceMessageCommand | Execute] Error while sending command: ", c.command)
	} else {
		logger.Logger().WriteToLog(logger.Info, "[SendDiagCommand | Execute] Command ", c.command, " sent")
	}
	return list.New()
}
