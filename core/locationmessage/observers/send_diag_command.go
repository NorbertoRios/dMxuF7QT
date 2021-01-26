package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	coreObservers "genx-go/core/observers"
	"genx-go/logger"
)

//NewSendDiagCommand ...
func NewSendDiagCommand(_task interfaces.ITask) *SendDiagCommand {
	return &SendDiagCommand{
		task: _task,
	}
}

//SendDiagCommand ...
type SendDiagCommand struct {
	task interfaces.ITask
}

//Execute ...
func (c *SendDiagCommand) Execute(device interfaces.IDevice) *list.List {
	commands := list.New()
	if err := device.Send("DIAG PARAMS=24"); err != nil {
		logger.Logger().WriteToLog(logger.Error, "[SendDiagCommand | Execute] Error while sending command \"DIAG PARAMS=24\"")
	} else {
		logger.Logger().WriteToLog(logger.Info, "[SendDiagCommand | Execute] Command \"DIAG PARAMS=24\" sent")
	}
	commands.PushBack(coreObservers.NewAttachObserverCommand(NewSyncObserver(c.task)))
	return commands
}
