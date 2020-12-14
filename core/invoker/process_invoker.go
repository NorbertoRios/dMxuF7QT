package invoker

import (
	"container/list"
	"genx-go/core/device/interfaces"
	immoObservers "genx-go/core/immobilizer/observers"
	"genx-go/core/observers"
	"genx-go/core/request"
)

//NewProcessInvoker ...
func NewProcessInvoker(_process interfaces.IProcess) *ProcessInvoker {
	return &ProcessInvoker{
		process: _process,
	}
}

//ProcessInvoker ...
type ProcessInvoker struct {
	process interfaces.IProcess
}

//NewImmoRequest ...
func (invoker *ProcessInvoker) NewImmoRequest(req request.IRequest) *list.List {
	return invoker.process.NewRequest(req)
}

//AckWatchdogsCommands ...
func (invoker *ProcessInvoker) AckWatchdogsCommands(task interfaces.ITask) *list.List {
	cmd := list.New()
	anyMessageObserver := immoObservers.NewAnyImmoMessageObserver(task)
	cmd.PushBack(observers.NewAttachObserverCommand(anyMessageObserver))
	return cmd
}

//DiagWatchdogsCommands ...
func (invoker *ProcessInvoker) DiagWatchdogsCommands(task interfaces.ITask) *list.List {
	cmd := list.New()
	anyMessageObserver := immoObservers.NewAnyImmoDiagObserver(task)
	cmd.PushBack(observers.NewAttachObserverCommand(anyMessageObserver))
	return cmd
}

//CanselTask ...
func (invoker *ProcessInvoker) CanselTask(_task interfaces.ITask, description string) *list.List {
	invoker.process.TaskCancel(_task, description)
	return invoker.dropAllObservers(_task)
}

//DoneTask ...
func (invoker *ProcessInvoker) DoneTask(_task interfaces.ITask) *list.List {
	invoker.process.TaskDone(_task)
	return invoker.dropAllObservers(_task)
}

func (invoker *ProcessInvoker) dropAllObservers(_task interfaces.ITask) *list.List {
	cmd := list.New()
	cmd.PushBack(observers.NewDetachAllTaskObserversCommad(_task))
	return cmd
}
