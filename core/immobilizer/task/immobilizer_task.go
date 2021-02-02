package task

import (
	"container/list"
	"encoding/json"
	"genx-go/core/device/interfaces"
	"genx-go/core/filter"
	"genx-go/core/immobilizer/observers"
	"genx-go/core/immobilizer/request"
	"genx-go/core/invoker"
	bRequest "genx-go/core/request"
	"genx-go/logger"
	"time"
)

//NewImmobilizerTask ...
func NewImmobilizerTask(_request *request.ChangeImmoStateRequest, immo interfaces.IImmobilizer, _device interfaces.IDevice) *ImmobilizerTask {
	return &ImmobilizerTask{
		BornTime:      time.Now().UTC(),
		FacadeRequest: _request,
		device:        _device,
		invoker:       invoker.NewImmobilizerInvoker(immo),
	}
}

//ImmobilizerTask in progress state
type ImmobilizerTask struct {
	invoker       interfaces.IImmoInvoker
	BornTime      time.Time                       `json:"CreatedAt"`
	FacadeRequest *request.ChangeImmoStateRequest `json:"FacadeRequest"`
	device        interfaces.IDevice
}

//Device returns task's observer
func (task *ImmobilizerTask) Device() interfaces.IDevice {
	return task.device
}

//Request ...
func (task *ImmobilizerTask) Request() interface{} {
	return task.FacadeRequest
}

//Invoker ...
func (task *ImmobilizerTask) Invoker() interfaces.IInvoker {
	return task.invoker
}

//Observers returns task's observer
func (task *ImmobilizerTask) Observers() []interfaces.IObserver {
	filter := filter.NewObserversFilter(task.device.Observable())
	return filter.Extract(task)
}

//Commands ..
func (task *ImmobilizerTask) Commands() *list.List {
	cList := list.New()
	out := &bRequest.OutputNumber{Data: task.FacadeRequest.Port}
	immo := task.device.Immobilizer(out.Index(), task.FacadeRequest.Trigger)
	if immo.State(task.device) == task.FacadeRequest.State {
		immo.CurrentTask().Invoker().CancelTask(immo.CurrentTask(), "Actual")
	} else {
		cList.PushFront(observers.NewImmoSendRelayCommand(task))
	}
	return cList
}

//Marshal ...
func (task *ImmobilizerTask) Marshal() string {
	jTask, jerr := json.Marshal(task)
	if jerr != nil {
		logger.Logger().WriteToLog(logger.Error, "[ImmobilizerTask | Marshal] Error while marshaling task. Error:", jerr)
		return ""
	}
	return string(jTask)
}
