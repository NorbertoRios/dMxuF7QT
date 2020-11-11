package device

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/logger"
)

//NewPushToRabbitMessageCommand ..
func NewPushToRabbitMessageCommand(_message string, _destinations ...string) *PushToRabbitMessageCommand {
	return &PushToRabbitMessageCommand{
		message:      _message,
		destinations: _destinations,
	}
}

//PushToRabbitMessageCommand ..
type PushToRabbitMessageCommand struct {
	message      string
	destinations []string
}

//Send ..
func (c *PushToRabbitMessageCommand) Send(device interfaces.IDevice) *list.List {
	for _, d := range c.destinations {
		switch d {
		case FacadeResponse, Message:
			{
				device.PushToRabbit(c.message, d)
			}
		default:
			{
				logger.Logger().WriteToLog(logger.Error, "[PushToRabbitMessageCommand | Execute] Unexpected destination ", d)
				continue
			}
		}
	}
	return nil

}

//Execute command
func (c *PushToRabbitMessageCommand) Execute(device interfaces.IDevice) *list.List {
	return c.Send(device)
}
