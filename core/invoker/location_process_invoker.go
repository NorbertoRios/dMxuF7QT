package invoker

import (
	"container/list"
	"genx-go/core/device/interfaces"
	locationObservers "genx-go/core/locationmessage/observers"
	"genx-go/core/observers"
	"strings"
)

//NewLocationProcessInvoker ...
func NewLocationProcessInvoker(_process interfaces.ILocationMessageProcess) *LocationProcessInvoker {
	return &LocationProcessInvoker{
		process: _process,
	}
}

//LocationProcessInvoker ...
type LocationProcessInvoker struct {
	BaseInvoker
	process interfaces.ILocationMessageProcess
}

//SendDiagCommandAfterAnyMessage ...
func (invoker *LocationProcessInvoker) SendDiagCommandAfterAnyMessage(task interfaces.ITask) *list.List {
	cmd := list.New()
	cmd.PushBack(observers.NewAttachObserverCommand(locationObservers.NewAnyMessageObserver(task)))
	return cmd
}

//DeviceSynchronized ...
func (invoker *LocationProcessInvoker) DeviceSynchronized(param24 string, _device interfaces.IDevice) *list.List {
	cmd := list.New()
	value := strings.ReplaceAll(strings.Split(param24, "=")[1], ";", "")
	param24Arr := strings.Split(value, ".")
	cmd.PushBackList(invoker.process.Param24Arriver(param24Arr, _device))
	return cmd
}
