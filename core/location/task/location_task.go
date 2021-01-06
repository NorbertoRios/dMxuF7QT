package task

import (
	"container/list"
	"encoding/json"
	"genx-go/core/device/interfaces"
	"genx-go/core/filter"
	"genx-go/core/invoker"
	"genx-go/core/location/observers"
	"genx-go/core/request"
	"genx-go/logger"
	"time"
)

//NewLocationTask ...
func NewLocationTask(_request *request.BaseRequest, _device interfaces.IDevice, _process interfaces.IProcess) *LocationTask {
	return &LocationTask{
		FacadeRequest: _request,
		device:        _device,
		invoker:       invoker.NewLocationInvoker(_process),
	}
}

//LocationTask ...
type LocationTask struct {
	BornTime      time.Time            `json:"CreatedAt"`
	FacadeRequest *request.BaseRequest `json:"FacadeRequest"`
	invoker       interfaces.ILocationInvoker
	device        interfaces.IDevice
}

//Invoker ...
func (task *LocationTask) Invoker() interfaces.IInvoker {
	return task.invoker
}

//Commands ..
func (task *LocationTask) Commands() *list.List {
	cList := list.New()
	cList.PushBack(observers.NewSendLocationRequest(task))
	return cList
}

//Observers ..
func (task *LocationTask) Observers() []interfaces.IObserver {
	taskFilter := filter.NewObserversFilter(task.device.Observable())
	return taskFilter.Extract(task)
}

//Device ..
func (task *LocationTask) Device() interfaces.IDevice {
	return task.device
}

//Request ...
func (task *LocationTask) Request() interface{} {
	return task.FacadeRequest
}

//Marshal ...
func (task *LocationTask) Marshal() string {
	jTask, jerr := json.Marshal(task)
	if jerr != nil {
		logger.Logger().WriteToLog(logger.Error, "[ImmobilizerTask | Marshal] Error while marshaling task. Error:", jerr)
		return ""
	}
	return string(jTask)
}
