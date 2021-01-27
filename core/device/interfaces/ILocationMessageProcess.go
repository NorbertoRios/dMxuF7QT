package interfaces

import (
	"container/list"
)

//ILocationMessageProcess ...
type ILocationMessageProcess interface {
	Param24Arriver([]string, IDevice) *list.List
	MessageIncome(interface{}, IDevice)
}
