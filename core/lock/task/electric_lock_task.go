package device

import (
	"encoding/json"
	"genx-go/core/device/interfaces"
	"genx-go/core/lock/request"
	"genx-go/logger"
	"time"
)

//NewElectricLockTask ...
func NewElectricLockTask(_request *request.UnlockRequest, _device interfaces.IDevice, _onCancel func(*ElectricLockTask, string), _onDone func(*ElectricLockTask)) *ElectricLockTask {
	return &ElectricLockTask{
		BornTime: time.Now().UTC(),
		Request:  _request,
		device:   _device,
		onCancel: _onCancel,
		onDone:   _onDone,
	}
}

//ElectricLockTask ...
type ElectricLockTask struct {
	BornTime time.Time              `json:"CreatedAt"`
	Request  *request.UnlockRequest `json:"FacadeRequest"`
	device   interfaces.IDevice
	onCancel func(*ElectricLockTask, string)
	onDone   func(*ElectricLockTask)
}

//Start ...
func (task *ElectricLockTask) Start() {
	if time.Now().UTC().Before(task.Request.ExpirationTime) {
		task.Cancel("Task timed out")
		return
	}

	//cList := list.New()
	//out := &request.OutputNumber{Data: task.Request}
	//immo := task.device.ImmobilizerStorage().Immobilizer(out.Index(), task.Request.Trigger, task.device)
	//if immo.State() == task.Request.State {
	//	response := response.NewResponse(task.Request.FacadeCallbackID, true, "Actual")
	//	cList.PushFront(NewPushToRabbitMessageCommand(response.Marshal(), FacadeResponse, Message))
	//} else {
	//	cList.PushFront(NewImmoSendRelayCommand(task))
	//}
	//task.device.ProccessCommands(cList)
}

//Done task
func (task *ElectricLockTask) Done() {
	task.onDone(task)
}

//Observers ..
func (task *ElectricLockTask) Observers() []IObserver {
	filter := NewObserversFilter(task.device.GetObservable())
	return filter.Extract(task)
}

//Cancel task
func (task *ElectricLockTask) Cancel(description string) {
	task.onCancel(task, description)
}

//Marshal ...
func (task *ElectricLockTask) Marshal() string {
	jTask, jerr := json.Marshal(task)
	if jerr != nil {
		logger.Logger().WriteToLog(logger.Error, "[ElectricLockTask | Marshal] Error while marshaling task. Error:", jerr)
		return ""
	}
	return string(jTask)
}
