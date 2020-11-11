package interfaces

import "container/list"

//ICommand command for device
type ICommand interface {
	Execute(IDevice) *list.List
}
