package invoker

import (
	"container/list"
	"genx-go/core/device/interfaces"
	locationObserver "genx-go/core/location/observers"
	"genx-go/core/observers"
)

//NewLocationInvoker ...
func NewLocationInvoker(_process interfaces.IProcess) *LocationInvoker {
	invoker := &LocationInvoker{}
	invoker.process = _process
	return invoker
}

//LocationInvoker ...
type LocationInvoker struct {
	BaseInvoker
}

//LocationWatchdogCommands ...
func (invoker *LocationInvoker) LocationWatchdogCommands(task interfaces.ITask) *list.List {
	cmd := list.New()
	cmd.PushBack(observers.NewAttachObserverCommand(locationObserver.NewLocationAnyMessageObserver(task)))
	return cmd
}
