package interfaces

import "container/list"

//ILocationProcessInvoker ...
type ILocationProcessInvoker interface {
	IInvoker
	DeviceSynchronized(string, IDevice) *list.List
	SendDiagCommandAfterAnyMessage(ITask) *list.List
}
