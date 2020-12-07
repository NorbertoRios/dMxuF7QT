package task

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/filter"
	"genx-go/core/location/observers"
	"genx-go/core/request"
	"time"
)

//NewLocationTask ...
func NewLocationTask(_request *request.BaseRequest, _device interfaces.IDevice, onCancel func(*LocationTask, string), onDone func(*LocationTask)) *LocationTask {
	return &LocationTask{
		request:      _request,
		device:       _device,
		onTaskCancel: onCancel,
		onTaskDone:   onDone,
	}
}

//LocationTask ...
type LocationTask struct {
	BornTieme    time.Time
	device       interfaces.IDevice
	onTaskCancel func(*LocationTask, string)
	onTaskDone   func(*LocationTask)
	request      *request.BaseRequest
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

//Cancel ..
func (task *LocationTask) Cancel(description string) {
	task.onTaskCancel(task, description)
}

//Done ..
func (task *LocationTask) Done() {
	task.onTaskDone(task)
}

//Request ...
func (task *LocationTask) Request() interface{} {
	return task.request
}
