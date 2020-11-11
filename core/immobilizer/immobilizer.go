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

//CurrentTask ...
func (immo *Immobilizer) CurrentTask() interfaces.ITask {
	return immo.currentTask
}

//NewRequest process new request to device
func (immo *Immobilizer) NewRequest(req *request.ChangeImmoStateRequest) {
	newTask := task.NewImmobilizerTask(req, immo.device, immo.taskCanceled, immo.taskDone)
	if immo.currentTask == nil {
		immo.currentTask = newTask
		immo.currentTask.Start()
		return
	}
	immo.competitivenessOfTasks(newTask)
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

func (immo *Immobilizer) competitivenessOfTasks(task *task.ImmobilizerTask) {
	req := immo.currentTask.Request().(*request.ChangeImmoStateRequest)
	if req.Equal(task.Request().(*request.ChangeImmoStateRequest)) {
		task.Cancel("Duplicate")
	} else {
		immo.currentTask.Cancel("Deprecated")
		immo.currentTask = task
		immo.currentTask.Start()
		logger.Logger().WriteToLog(logger.Info, "Task created and run")
	}
}

func (immo *Immobilizer) taskCanceled(canseledTask *task.ImmobilizerTask, description string) {
	logger.Logger().WriteToLog(logger.Info, "Task is canceled")
	immo.mutex.Lock()
	immo.tasks.PushBack(task.NewCanceledImmoTask(canseledTask, description))
	immo.mutex.Unlock()
	immo.currentTask = nil
}

func (immo *Immobilizer) taskDone(doneTask *task.ImmobilizerTask) {
	logger.Logger().WriteToLog(logger.Info, "Task is done")
	immo.mutex.Lock()
	immo.tasks.PushFront(task.NewDoneImmoTask(doneTask))
	immo.mutex.Unlock()
	immo.currentTask = nil
}
