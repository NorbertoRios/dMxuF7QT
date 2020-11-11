package task

import (
	"container/list"
	"encoding/json"
	"genx-go/core/device/interfaces"
	"genx-go/core/filter"
	"genx-go/core/immobilizer/observers"
	"genx-go/core/immobilizer/request"
	bRequest "genx-go/core/request"
	"genx-go/logger"
	"time"
)

//NewImmobilizerTask ...
func NewImmobilizerTask(_request *request.ChangeImmoStateRequest, _device interfaces.IDevice, _onCancel func(*ImmobilizerTask, string), _onDone func(*ImmobilizerTask)) *ImmobilizerTask {
	return &ImmobilizerTask{
		BornTime:      time.Now().UTC(),
		FacadeRequest: _request,
		device:        _device,
		onCancel:      _onCancel,
		onDone:        _onDone,
	}
}

//ImmobilizerTask in progress state
type ImmobilizerTask struct {
	BornTime      time.Time                       `json:"CreatedAt"`
	FacadeRequest *request.ChangeImmoStateRequest `json:"FacadeRequest"`
	device        interfaces.IDevice
	onCancel      func(*ImmobilizerTask, string)
	onDone        func(*ImmobilizerTask)
}

//Device returns task's observer
func (task *ImmobilizerTask) Device() interfaces.IDevice {
	return task.device
}

//Request ...
func (task *ImmobilizerTask) Request() interface{} {
	return task.FacadeRequest
}

//Observers returns task's observer
func (task *ImmobilizerTask) Observers() []interfaces.IObserver {
	filter := filter.NewObserversFilter(task.device.GetObservable())
	return filter.Extract(task)
}

//Start ..
func (task *ImmobilizerTask) Start() {
	cList := list.New()
	out := &bRequest.OutputNumber{Data: task.FacadeRequest.Port}
	immo := task.device.Immobilizer(out.Index(), task.FacadeRequest.Trigger)
	if immo.State() == task.FacadeRequest.State {
		immo.CurrentTask().Cancel("Actual")
	} else {
		cList.PushFront(observers.NewImmoSendRelayCommand(task))
	}
	task.device.ProccessCommands(cList)
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

//Done task
func (task *ImmobilizerTask) Done() {
	// cList := list.New()
	// command := NewPushToRabbitMessageCommand("dasdas",FacadeResponse)
	// cList.PushBack(command)
	// task.device.ProccessCommands(cList)
	task.onDone(task)
}

//Cancel task
func (task *ImmobilizerTask) Cancel(description string) {
	task.onCancel(task, description)
}
