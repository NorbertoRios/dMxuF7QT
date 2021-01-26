package interfaces

import (
	"container/list"
	"genx-go/core/locationmessage/request"
)

//ILocationMessageProcess ...
type ILocationMessageProcess interface {
	Start() *list.List
	Param24Arriver([]string) *list.List
	MessageIncome(*request.MessageRequest, interface{})
}
