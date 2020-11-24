package interfaces

import "container/list"

//IConfigTask ...
type IConfigTask interface {
	ITask
	NextStep() *list.List
}
