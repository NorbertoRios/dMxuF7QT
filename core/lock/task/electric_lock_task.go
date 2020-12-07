package task

import (
	"container/list"
	"encoding/json"
	"genx-go/core/device/interfaces"
	"genx-go/core/filter"
	"genx-go/core/lock/observers"
	"genx-go/core/lock/request"
	coreObservers "genx-go/core/observers"
	"genx-go/logger"
	"time"
)

//NewElectricLockTask ...
func NewElectricLockTask(_request *request.UnlockRequest, _device interfaces.IDevice, _onCancel func(*ElectricLockTask, string), _onDone func(*ElectricLockTask)) *ElectricLockTask {
	return &ElectricLockTask{
		BornTime:      time.Now().UTC(),
		FacadeRequest: _request,
		device:        _device,
		onCancel:      _onCancel,
		onDone:        _onDone,
	}
}

//ElectricLockTask ...
type ElectricLockTask struct {
	BornTime      time.Time              `json:"CreatedAt"`
	FacadeRequest *request.UnlockRequest `json:"FacadeRequest"`
	device        interfaces.IDevice
	onCancel      func(*ElectricLockTask, string)
	onDone        func(*ElectricLockTask)
}

//Commands ...
func (task *ElectricLockTask) Commands() *list.List {
	if task.FacadeRequest.Time().Before(time.Now().UTC()) {
		logger.Logger().WriteToLog(logger.Info, "[ElectricLockTask | Start] Task time is over. Task expiration time: ", task.FacadeRequest.Time().String(), ". Current time: ", time.Now().UTC().String())
		task.Cancel("Task timed out")
		return list.New()
	}
	cList := list.New()
	cList.PushFront(observers.NewElectricLockSendCommand(task))
	logger.Logger().WriteToLog(logger.Info, "[ElectricLockTask | Start] Task starded. Task expiration time: ", task.FacadeRequest.Time().String(), ". Current time: ", time.Now().UTC().String())
	return cList
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

func (task *ElectricLockTask) detachAllTaskObservers() {
	cList := list.New()
	for _, observer := range task.Observers() {
		cList.PushBack(coreObservers.NewDetachObserverCommand(observer))
	}
	task.device.ProcessCommands(cList)
	logger.Logger().WriteToLog(logger.Info, "ALL OBSERVERS DETACHED", task.Observers())
}

//Done task
func (task *ElectricLockTask) Done() {
	task.onDone(task)
	logger.Logger().WriteToLog(logger.Info, "Task Done.")
	task.detachAllTaskObservers()
}

//Cancel task
func (task *ElectricLockTask) Cancel(description string) {
	task.onCancel(task, description)
	logger.Logger().WriteToLog(logger.Info, "Task CANCELED. ", description)
	task.detachAllTaskObservers()
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
