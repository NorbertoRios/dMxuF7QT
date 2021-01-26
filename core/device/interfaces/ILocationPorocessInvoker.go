package interfaces

import "container/list"

//ILocationProcessInvoker ...
type ILocationProcessInvoker interface {
	IInvoker
	DeviceSynchronized(string) *list.List
	SendDiagCommandAfterAnyMessage(ITask) *list.List
}
