package interfaces

import (
	"container/list"
)

//ILocationMessageProcess ...
type ILocationMessageProcess interface {
	IProcess
	Param24Arriver([]string, IDevice) *list.List
}
