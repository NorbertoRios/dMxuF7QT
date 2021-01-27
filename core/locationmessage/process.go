package locationmessage

import (
	"container/list"
	"fmt"
	"genx-go/core/device/interfaces"
	"genx-go/core/locationmessage/task"
	"genx-go/core/process"
	"genx-go/core/usecase"
)

//NewLocationProcess ...
func NewLocationProcess(parameter24 []string) *Process {
	var process *Process
	if len(parameter24) == 0 {
		process = &Process{}
	} else {
		process = &Process{
			param24: parameter24,
		}
	}
	return process
}

//Process ...
type Process struct {
	param24 []string
	process.BaseProcess
}

//Param24Arriver ...
func (p *Process) Param24Arriver(param24 []string, _device interfaces.IDevice) *list.List {
	cmdList := list.New()
	cmdList.PushBackList(p.ProcessCurrentTask.Invoker().CanselTask(p.ProcessCurrentTask, fmt.Sprintf("[LocationMessageProcess] New 24 parameter arrived : %v", param24)))
	newTask := task.NewLocationMessageTask(p, _device)
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

//MessageIncome ...
func (p *Process) MessageIncome(incomeMessage interface{}, device interfaces.IDevice) {
	p.initTask(device)
	usecase.NewMessageArrivedUseCase(device, incomeMessage).Launch()
}
