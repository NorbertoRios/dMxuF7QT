package invoker

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
)

//BaseInvoker ...
type BaseInvoker struct {
	process interfaces.IProcess
}

//CancelTask ...
func (invoker *BaseInvoker) CancelTask(_task interfaces.ITask, description string) *list.List {
	invoker.process.TaskCancel(_task, description)
	return invoker.dropAllObservers(_task)
}

//DoneTask ...
func (invoker *BaseInvoker) DoneTask(_task interfaces.ITask) *list.List {
	invoker.process.TaskDone(_task)
	return invoker.dropAllObservers(_task)
}

func (invoker *BaseInvoker) dropAllObservers(_task interfaces.ITask) *list.List {
	cmd := list.New()
	cmd.PushBack(observers.NewDetachAllTaskObserversCommad(_task))
	return cmd
}
