package location

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/location/task"
	"genx-go/core/observers"
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

//Task ...
func (request *Request) Task() interfaces.ITask {
	return request.task
}

func (request *Request) detachTaskObservers(_task interfaces.ITask) {
	detach := observers.NewDetachTaskObservers(_task)
	request.device.ProcessCommands(detach.Commands())
}

func (request *Request) done(_task *task.LocationTask) {
	logger.Logger().WriteToLog(logger.Info, "[LocationRequest | done] Task done")
	request.detachTaskObservers(_task)
}

func (request *Request) cancel(_task *task.LocationTask, description string) {
	logger.Logger().WriteToLog(logger.Info, "[LocationRequest | cancel] Task canceled. Description: ", description)
	request.detachTaskObservers(_task)
}
