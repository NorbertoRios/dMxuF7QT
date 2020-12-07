package task

import (
	"container/list"
	"genx-go/core/configuration/observers"
	"genx-go/core/configuration/request"
	"genx-go/core/device/interfaces"
	"genx-go/core/filter"
	coreObservers "genx-go/core/observers"
	"genx-go/core/usecase"
	"time"
)

//New ...
func New(_request *request.ConfigurationRequest, device interfaces.IDevice, _onCancel func(*ConfigTask, string), _onDone func(*ConfigTask)) *ConfigTask {
	return &ConfigTask{
		BornTime:       time.Now().UTC(),
		FacadeRequest:  _request,
		ConfigCommands: _request.Commands(),
		device:         device,
		onCancel:       _onCancel,
		onDone:         _onDone,
	}
}

//ConfigTask ...
type ConfigTask struct {
	BornTime       time.Time                     `json:"CreatedAt"`
	FacadeRequest  *request.ConfigurationRequest `json:"FacadeRequest"`
	ConfigCommands *list.List                    `json:"Commands"`
	currentCommand *list.Element
	device         interfaces.IDevice
	onCancel       func(*ConfigTask, string)
	onDone         func(*ConfigTask)
}

//Device ...
func (task *ConfigTask) Device() interfaces.IDevice {
	return task.device
}

//Commands ...
func (task *ConfigTask) Commands() *list.List {
	if task.currentCommand == nil {
		task.currentCommand = task.ConfigCommands.Front()
	}
	cList := list.New()
	cList.PushBack(observers.NewSendConfigCommand(task, task.currentCommand.Value.(*request.Command).Command()))
	return cList
}

//Observers ...
func (task *ConfigTask) Observers() []interfaces.IObserver {
	f := filter.NewObserversFilter(task.device.Observable())
	return f.Extract(task)
}

//Request ...
func (task *ConfigTask) Request() interface{} {
	return task.FacadeRequest
}

//NextStep ...
func (task *ConfigTask) NextStep() *list.List {
	cList := list.New()
	task.currentCommand.Value.(*request.Command).Complete()
	if cmd := task.currentCommand.Next(); cmd != nil {
		task.currentCommand = cmd
		cList.PushBack(observers.NewSendConfigCommand(task, task.currentCommand.Value.(*request.Command).Command()))
	} else {
		task.Done()
	}
	return cList
}

//Done ...
func (task *ConfigTask) Done() {
	task.cleanObservers()
	task.onDone(task)
}

//Cancel ...
func (task *ConfigTask) Cancel(description string) {
	task.cleanObservers()
	task.onCancel(task, description)
}

func (task *ConfigTask) cleanObservers() {
	cList := list.New()
	oFilter := filter.NewObserversFilter(task.device.Observable())
	for _, o := range oFilter.Extract(task) {
		cList.PushBack(coreObservers.NewDetachObserverCommand(o))
	}
	useCase := usecase.NewBaseUseCase(task.device, cList)
	useCase.Launch()
}
