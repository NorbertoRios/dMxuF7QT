package immobilizer

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/immobilizer/request"
	"genx-go/core/immobilizer/task"
	baseRequest "genx-go/core/request"
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
func (immo *Immobilizer) NewRequest(req baseRequest.IRequest) *list.List {
	newTask := task.NewImmobilizerTask(req.(*request.ChangeImmoStateRequest), immo, immo.Device())
	if immo.currentTask == nil {
		immo.currentTask = newTask
		return newTask.Commands()
	}
	return immo.competitivenessOfTasks(newTask, immo.currentTask.Request().(*request.ChangeImmoStateRequest))
}

func (immo *Immobilizer) competitivenessOfTasks(newTask interfaces.ITask, currentRequest *request.ChangeImmoStateRequest) *list.List {
	if currentRequest.Equal(newTask.Request().(*request.ChangeImmoStateRequest)) {
		return newTask.Invoker().CanselTask(newTask, "Duplicate")
	}
	cmdList := list.New()
	cmdList.PushBackList(immo.currentTask.Invoker().CanselTask(immo.currentTask, "Deprecated"))
	immo.currentTask = newTask
	cmdList.PushBackList(immo.currentTask.Commands())
	return cmdList
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

//TaskCancel ...
func (immo *Immobilizer) TaskCancel(canseledTask interfaces.ITask, description string) {
	logger.Logger().WriteToLog(logger.Info, "Task is canceled. ", description)
	immo.pushToTasks(task.NewCanceledImmoTask(canseledTask, description), false)
}

//TaskDone ...
func (immo *Immobilizer) TaskDone(doneTask interfaces.ITask) {
	logger.Logger().WriteToLog(logger.Info, "Task is done")
	immo.pushToTasks(task.NewDoneImmoTask(doneTask), true)
}

func (immo *Immobilizer) pushToTasks(_task interfaces.ITask, isDone bool) {
	immo.mutex.Lock()
	defer immo.mutex.Unlock()
	if isDone {
		immo.tasks.PushFront(_task)
	} else {
		immo.tasks.PushBack(_task)
	}
}
