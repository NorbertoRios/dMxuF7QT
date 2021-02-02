package locationmessage

import (
	"container/list"
	"fmt"
	"genx-go/core/device/interfaces"
	"genx-go/core/locationmessage/task"
	"genx-go/core/process"
	"genx-go/core/usecase"
	"sync"
)

//NewLocationProcess ...
func NewLocationProcess(parameter24 []string) *Process {
	var p *Process
	if len(parameter24) == 0 {
		p = &Process{}
	} else {
		p = &Process{
			param24: parameter24,
		}
	}
	p.Mutex = &sync.Mutex{}
	p.ProcessTasks = list.New()
	return p
}

//Process ...
type Process struct {
	param24 []string
	process.BaseProcess
}

//Param24Arriver ...
func (p *Process) Param24Arriver(param24 []string, _device interfaces.IDevice) *list.List {
	cmdList := list.New()
	cmdList.PushBackList(p.ProcessCurrentTask.Invoker().CancelTask(p.ProcessCurrentTask, fmt.Sprintf("[LocationMessageProcess] New 24 parameter arrived : %v", param24)))
	newTask := task.NewLocationMessageTask(p, _device)
	_device.New24Param(param24)
	cmdList.PushBackList(newTask.Commands())
	return cmdList
}

func (p *Process) initTask(_device interfaces.IDevice) {
	if p.ProcessCurrentTask != nil {
		return
	}
	if len(p.param24) == 0 {
		p.ProcessCurrentTask = task.NewSyncTask(p, _device)
	} else {
		p.ProcessCurrentTask = task.NewLocationMessageTask(p, _device)
	}
	usecase.NewBaseUseCase(_device, p.ProcessCurrentTask.Commands()).Launch()
}

//NewRequest ...
func (p *Process) NewRequest(incomeMessage interface{}, device interfaces.IDevice) *list.List {
	p.initTask(device)
	usecase.NewMessageArrivedUseCase(device, incomeMessage).Launch()
	return list.New()
}

//TaskCancel ...
func (p *Process) TaskCancel(canseledTask interfaces.ITask, description string) {
	var _task interfaces.ITask
	switch canseledTask.(type) {
	case *task.SyncTask:
		{
			_task = task.NewCanceledSyncTask(canseledTask, description)
		}
	default:
		{
			_task = task.NewCanceledLocationMessageTask(canseledTask, description)
		}
	}
	p.PushToTasks(_task, false)
}

//TaskDone ...
func (p *Process) TaskDone(doneTask interfaces.ITask) {
	var _task interfaces.ITask
	switch doneTask.(type) {
	case *task.SyncTask:
		{
			_task = task.NewDoneSyncTask(doneTask)
		}
	default:
		{
			_task = task.NewDoneLocationMessageTask(doneTask)
		}
	}
	p.PushToTasks(_task, true)
}
