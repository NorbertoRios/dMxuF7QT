package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
)

//NewDetachTaskObservers ...
func NewDetachTaskObservers(_task interfaces.ITask) *DetachTaskObservers {
	return &DetachTaskObservers{
		task: _task,
	}
}

//DetachTaskObservers ...
type DetachTaskObservers struct {
	task interfaces.ITask
}

//Commands ...
func (dto *DetachTaskObservers) Commands() *list.List {
	cList := list.New()
	for _, observer := range dto.task.Observers() {
		cList.PushBack(NewDetachObserverCommand(observer))
	}
	return cList
}
