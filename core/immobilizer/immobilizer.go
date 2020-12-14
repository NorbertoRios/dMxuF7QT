package immobilizer

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/immobilizer/request"
	"genx-go/core/immobilizer/task"
	"genx-go/core/sensors"
	"genx-go/logger"
	"sync"
)

//NewImmobilizer ...
func NewImmobilizer(_device interfaces.IDevice, index int, trigger string) *Immobilizer {
	return &Immobilizer{
		device:       _device,
		tasks:        list.New(),
		OutputNumber: index,
		trigger:      trigger,
		mutex:        &sync.Mutex{},
	}
}

//Immobilizer process
type Immobilizer struct {
	mutex        *sync.Mutex
	device       interfaces.IDevice
	OutputNumber int
	trigger      string
	currentTask  interfaces.ITask
	tasks        *list.List
}

//Trigger ...
func (immo *Immobilizer) Trigger() string {
	return immo.trigger
}

//Device ...
func (immo *Immobilizer) Device() interfaces.IDevice {
	return immo.device
}

//CurrentTask ...
func (immo *Immobilizer) CurrentTask() interfaces.ITask {
	return immo.currentTask
}

//NewRequest ...
func (immo *Immobilizer) NewRequest(req *request.ChangeImmoStateRequest) *list.List {
	newTask := task.NewImmobilizerTask(req, immo, immo.Device())
	if immo.currentTask == nil {
		immo.currentTask = newTask
		return task.Commands()
	}
	req := immo.currentTask.Request().(*request.ChangeImmoStateRequest)
	if req.Equal(task.Request().(*request.ChangeImmoStateRequest)) {
		task.Invoker().Cancel(task, "Duplicate")
		return list.New()
	}
	immo.currentTask.Cancel("Deprecated")
	immo.currentTask = newTask
	return immo.currentTask.Commands()
}

//State ...
func (immo *Immobilizer) State() string {
	deviceState := immo.device.State()
	for sensor := range deviceState {
		switch sensor.(type) {
		case *sensors.Relay:
			{
				relay := sensor.(*sensors.Relay)
				if relay.ID == immo.OutputNumber {
					sState := request.NewImmoStateRelayBased(relay.State, immo.trigger)
					return sState.State()
				}
			}
		}
	}
	return ""
}

//Tasks ...
func (immo *Immobilizer) Tasks() *list.List {
	return immo.tasks
}

//TaskCanceled ...
func (immo *Immobilizer) TaskCanceled(canseledTask interfaces.ITask, description string) {
	logger.Logger().WriteToLog(logger.Info, "Task is canceled. ", description)
	immo.pushToTasks(doneTask, false)
}

//TaskDone ...
func (immo *Immobilizer) TaskDone(doneTask interfaces.ITask) {
	logger.Logger().WriteToLog(logger.Info, "Task is done")
	immo.pushToTasks(doneTask, true)
}

func (immo *Immobilizer) pushToTasks(_task interfaces.ITask, isFront bool) {
	immo.mutex.Lock()
	defer immo.mutex.Unlock()
	if isFront {
		immo.tasks.PushFront(task.NewDoneImmoTask(_task))
	} else {
		immo.tasks.PushBack(task.NewDoneImmoTask(_task))
	}
}
