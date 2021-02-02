package interfaces

import "container/list"

//IImmoInvoker ...
type IImmoInvoker interface {
	IInvoker
	WatchdogsCommands(ITask, string) *list.List
}
