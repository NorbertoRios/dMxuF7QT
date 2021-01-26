package interfaces

import "container/list"

//ILocationMessageProcess ...
type ILocationMessageProcess interface {
	Start() *list.List
	Param24Arriver([]string) *list.List
	//MessageIncome(*message.RawMessage, interface{})
}
