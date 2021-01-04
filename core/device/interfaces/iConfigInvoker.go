package interfaces

import (
	"container/list"
)

//IConfigInvoker ...
type IConfigInvoker interface {
	IInvoker
	SendConfigAfterAnyMessage(IConfigTask) *list.List
	SendCurrectCommand(IConfigTask) *list.List
	SendNextCommand(IConfigTask) *list.List
}
