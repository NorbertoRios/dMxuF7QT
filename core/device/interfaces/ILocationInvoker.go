package interfaces

import "container/list"

//ILocationInvoker ...
type ILocationInvoker interface {
	IInvoker
	LocationWatchdogCommands(ITask) *list.List
}
