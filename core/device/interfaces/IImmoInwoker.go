package interfaces

import "container/list"

//IImmoInvoker ...
type IImmoInvoker interface {
	IInvoker
	AckWatchdogsCommands(ITask) *list.List
	DiagWatchdogsCommands(ITask) *list.List
}
