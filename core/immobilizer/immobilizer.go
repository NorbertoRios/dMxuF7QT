package immobilizer

import (
	"container/list"
	"fmt"
	"genx-go/core/device/interfaces"
	"genx-go/core/immobilizer/request"
	"genx-go/core/immobilizer/task"
	"genx-go/core/process"
	baseRequest "genx-go/core/request"
	"genx-go/core/sensors"
	"genx-go/logger"
	"reflect"
	"sync"
)

//NewImmobilizer ...
func NewImmobilizer(index int, trigger string) *Immobilizer {
	immo := &Immobilizer{
		OutputNumber: index,
		trigger:      trigger,
	}
	immo.Mutex = &sync.Mutex{}
	immo.ProcessTasks = list.New()
	return immo
}

//Immobilizer process
type Immobilizer struct {
	process.BaseProcess
	OutputNumber int
	trigger      string
}

//Trigger ...
func (immo *Immobilizer) Trigger() string {
	return immo.trigger
}

//NewRequest ...
func (immo *Immobilizer) NewRequest(req baseRequest.IRequest, _device interfaces.IDevice) *list.List {
	newTask := task.NewImmobilizerTask(req.(*request.ChangeImmoStateRequest), immo, _device)
	if immo.ProcessCurrentTask == nil {
		immo.ProcessCurrentTask = newTask
		return newTask.Commands()
	}
	return immo.competitivenessOfTasks(newTask, immo.ProcessCurrentTask.Request().(*request.ChangeImmoStateRequest))
}

func (immo *Immobilizer) competitivenessOfTasks(newTask interfaces.ITask, currentRequest *request.ChangeImmoStateRequest) *list.List {
	if currentRequest.Equal(newTask.Request().(*request.ChangeImmoStateRequest)) {
		logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[Immobilizer | competitivenessOfTasks] Duplicate request %v", currentRequest.Marshal()))
		return newTask.Invoker().CanselTask(newTask, "Duplicate")
	}
	cmdList := list.New()
	cmdList.PushBackList(immo.ProcessCurrentTask.Invoker().CanselTask(immo.ProcessCurrentTask, "Deprecated"))
	cmdList.PushBackList(immo.ProcessCurrentTask.Commands())
	immo.ProcessCurrentTask = newTask
	return cmdList
}

//State ...
func (immo *Immobilizer) State(_device interfaces.IDevice) string {
	deviceState := _device.State()
	for _, sensor := range deviceState {
		switch sensor.(type) {
		case *sensors.Outputs:
			{
				relays := sensor.(*sensors.Outputs).Relays
				sState := request.NewImmoStateRelayBased(relays[immo.OutputNumber], immo.trigger)
				return sState.State()
			}
		}
	}
	return ""
}

//TaskCancel ...
func (immo *Immobilizer) TaskCancel(canseledTask interfaces.ITask, description string) {
	if reflect.DeepEqual(immo.ProcessCurrentTask, canseledTask) {
		logger.Logger().WriteToLog(logger.Info, "[Immobilizer] Current task canceled. Reason: ", description)
		immo.ProcessCurrentTask = nil
	} else {
		logger.Logger().WriteToLog(logger.Info, "[Immobilizer] Task is canceled. ", description)
	}
	immo.PushToTasks(task.NewCanceledImmoTask(canseledTask, description), false)
}

//TaskDone ...
func (immo *Immobilizer) TaskDone(doneTask interfaces.ITask) {
	immo.ProcessCurrentTask = nil
	logger.Logger().WriteToLog(logger.Info, "[Immobilizer] Task is done")
	immo.PushToTasks(task.NewDoneImmoTask(doneTask), true)
}
