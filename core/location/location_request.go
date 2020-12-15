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
func (request *Request) NewRequest(req *request.BaseRequest) *list.List {
	if request.task != nil {
		request.task.Cancel("Deprecated")
	}
	request.task = task.NewLocationTask(req, request.device, request.cancel, request.done)
	return request.task.Commands()
}

//CurrentTask ...
func (request *Request) CurrentTask() interfaces.ITask {
	return request.task
}

//Tasks ...
func (request *Request) Tasks() *list.List {
	return list.New()
}

//TaskCancel ...
func (request *Request) TaskCancel(_task interfaces.ITask) {
	logger.Logger().WriteToLog(logger.Info, "[LocationRequest | done] Task done")
}

//TaskDone ...
func (request *Request) TaskDone(_task interfaces.ITask, description string) {
	logger.Logger().WriteToLog(logger.Info, "[LocationRequest | cancel] Task canceled. Description: ", description)
}
