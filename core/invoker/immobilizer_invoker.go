package invoker

import (
	"container/list"
	"genx-go/core/device/interfaces"
	immoObservers "genx-go/core/immobilizer/observers"
	"genx-go/core/observers"
)

//NewImmobilizerInvoker ...
func NewImmobilizerInvoker(_process interfaces.IProcess) *ImmobilizerInvoker {
	invoker := &ImmobilizerInvoker{}
	invoker.process = _process
	return invoker
}

//ImmobilizerInvoker ...
type ImmobilizerInvoker struct {
	BaseInvoker
}

//WatchdogsCommands ...
func (invoker *ImmobilizerInvoker) WatchdogsCommands(task interfaces.ITask, _message string) *list.List {
	cmd := list.New()
	anyMessageObserver := immoObservers.NewAnyImmoMessageObserver(task, _message)
	cmd.PushBack(observers.NewAttachObserverCommand(anyMessageObserver))
	return cmd
}

//AckWatchdogsCommands ...
// func (invoker *ImmobilizerInvoker) AckWatchdogsCommands(task interfaces.ITask) *list.List {
// 	cmd := list.New()
// 	anyMessageObserver := immoObservers.NewAnyImmoMessageObserver(task)
// 	cmd.PushBack(observers.NewAttachObserverCommand(anyMessageObserver))
// 	return cmd
// }

//DiagWatchdogsCommands ...
// func (invoker *ImmobilizerInvoker) DiagWatchdogsCommands(task interfaces.ITask) *list.List {
// 	cmd := list.New()
// 	anyMessageObserver := immoObservers.NewAnyImmoDiagObserver(task)
// 	cmd.PushBack(observers.NewAttachObserverCommand(anyMessageObserver))
// 	return cmd
// }
