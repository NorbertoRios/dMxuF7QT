package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
	"genx-go/logger"
)

//NewSendConfigCommand ...
func NewSendConfigCommand(_task interfaces.ITask) *SendConfigCommand {
	return &SendConfigCommand{
		task: _task,
	}
}

//SendConfigCommand ..
type SendConfigCommand struct {
	task interfaces.ITask
}

//Execute ...
func (c *SendConfigCommand) Execute(device interfaces.IDevice) *list.List {
	commands := list.New()
	config := NewConfig(c.task.Request().(string))
	if err := device.Send(config.Command()); err != nil {
		logger.Logger().WriteToLog(logger.Error, "[ImmoSendRelayCommand | Execute] Error while sending command ", config.Command())
	}
	commands.PushBack(observers.NewAttachObserverCommand(NewWaitingConfigAckObserver(c.task)))
	//commands.PushBack(NewPushToRabbitMessageCommand(setRelayDrive.Command(), Message))
	return commands
}
