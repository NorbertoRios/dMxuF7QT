package location

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/location/task"
	"genx-go/core/request"
	"genx-go/logger"
)

//New ...
func New(_device interfaces.IDevice) *Request {
	return &Request{
		device: _device,
	}
}

//Request ..
type Request struct {
	device interfaces.IDevice
	task   interfaces.ITask
}

//NewRequest ...
func (lRequest *Request) NewRequest(req request.IRequest) *list.List {
	cList := list.New()
	if lRequest.task != nil {
		cList.PushBackList(lRequest.task.Invoker().CanselTask(lRequest.task, "Deprecated"))
	}
	lRequest.task = task.NewLocationTask(req.(*request.BaseRequest), lRequest.device)
	cList.PushBackList(lRequest.task.Commands())
	return cList
}

//CurrentTask ...
func (lRequest *Request) CurrentTask() interfaces.ITask {
	return lRequest.task
}

//Tasks ...
func (Request) Tasks() *list.List {
	return list.New()
}

//TaskDone ...
func (Request) TaskDone(_task interfaces.ITask) {
	logger.Logger().WriteToLog(logger.Info, "[LocationRequest | Done] Task done")
}

//TaskCancel ...
func (Request) TaskCancel(_task interfaces.ITask, description string) {
	logger.Logger().WriteToLog(logger.Info, "[LocationRequest | Cancel] Task canceled. Description: ", description)
}
