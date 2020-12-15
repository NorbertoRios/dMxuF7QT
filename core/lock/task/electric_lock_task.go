package task

import (
	"container/list"
	"encoding/json"
	"genx-go/core/device/interfaces"
	"genx-go/core/filter"
	"genx-go/core/invoker"
	"genx-go/core/lock/observers"
	"genx-go/core/lock/request"
	baseRequest "genx-go/core/request"
	"genx-go/logger"
	"time"
)

//NewElectricLockTask ...
func NewElectricLockTask(_request baseRequest.IRequest, _device interfaces.IDevice, _lock interfaces.IProcess) *ElectricLockTask {
	return &ElectricLockTask{
		BornTime:      time.Now().UTC(),
		FacadeRequest: _request,
		device:        _device,
		invoker:       invoker.NewElectricLockInvoker(_lock),
	}
}

//ElectricLockTask ...
type ElectricLockTask struct {
	BornTime      time.Time            `json:"CreatedAt"`
	FacadeRequest baseRequest.IRequest `json:"FacadeRequest"`
	device        interfaces.IDevice
	invoker       interfaces.IInvoker
}

//Commands ...
func (task *ElectricLockTask) Commands() *list.List {
	unlockRequest := task.FacadeRequest.(*request.UnlockRequest)
	if unlockRequest.Time().Before(time.Now().UTC()) {
		logger.Logger().WriteToLog(logger.Info, "[ElectricLockTask | Start] Task time is over. Task expiration time: ", unlockRequest.Time().String(), ". Current time: ", time.Now().UTC().String())
		return task.Invoker().CanselTask(task, "Task timed out")
	}
	cList := list.New()
	cList.PushFront(observers.NewElectricLockSendCommand(task))
	logger.Logger().WriteToLog(logger.Info, "[ElectricLockTask | Start] Task starded. Task expiration time: ", unlockRequest.Time().String(), ". Current time: ", time.Now().UTC().String())
	return cList
}

//Invoker ..
func (task *ElectricLockTask) Invoker() interfaces.IInvoker {
	return task.invoker
}

//Request ...
func (task *ElectricLockTask) Request() interface{} {
	return task.FacadeRequest
}

//Device ...
func (task *ElectricLockTask) Device() interfaces.IDevice {
	return task.device
}

//Observers ..
func (task *ElectricLockTask) Observers() []interfaces.IObserver {
	filter := filter.NewObserversFilter(task.device.Observable())
	return filter.Extract(task)
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
