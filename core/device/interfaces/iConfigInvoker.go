package interfaces

import (
	"container/list"
)

//IConfigInvoker ...
type IConfigInvoker interface {
	IInvoker
	Next(ITask) *list.List
	SendConfigAfterAnyMessage(ITask) *list.List
}
