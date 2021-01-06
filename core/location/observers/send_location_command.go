package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
	"genx-go/logger"
)

//NewSendLocationRequest ...
func NewSendLocationRequest(_task interfaces.ITask) *SendLocationRequest {
	return &SendLocationRequest{
		task: _task,
	}
}

//SendLocationRequest ...
type SendLocationRequest struct {
	task interfaces.ITask
}

//Execute ...
func (request *SendLocationRequest) Execute(device interfaces.IDevice) *list.List {
	cList := list.New()
	if err := device.Send("DIAG POLQ"); err != nil {
		logger.Logger().WriteToLog(logger.Error, "[SendLocationRequest | Execute] Error while sending command \"DIAG POLQ\"")
	}	
	cList.PushBack(observers.NewAttachObserverCommand(NewWaitingLocationMessageObserver(request.task)))
	//commands.PushBack(NewPushToRabbitMessageCommand(setRelayDrive.Command(), Message))
	return cList
}
