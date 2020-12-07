package interfaces

import (
	"container/list"
	"genx-go/core/configuration/request"
)

//IConfiguration ...
type IConfiguration interface {
	NewRequest(*request.ConfigurationRequest) *list.List
	CurrentTask() ITask
	Tasks() *list.List
}
