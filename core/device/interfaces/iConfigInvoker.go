package interfaces

import (
	"container/list"
)

//IConfigInvoker ...
type IConfigInvoker interface {
	IInvoker
	SendConfigAfterAnyMessage(IConfigTask) *list.List
	SendCommand(IConfigTask)
}
