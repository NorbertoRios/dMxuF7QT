package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
)

//NewDetachAllTaskObserversCommad ...
func NewDetachAllTaskObserversCommad(_task interfaces.ITask) *DetachAllTaskObserversCommad {
	return &DetachAllTaskObserversCommad{
		task: _task,
	}
}

//DetachAllTaskObserversCommad ..
type DetachAllTaskObserversCommad struct {
	task interfaces.ITask
}

//Execute ..
func (c *DetachAllTaskObserversCommad) Execute(device interfaces.IDevice) *list.List {
	observers := c.task.Observers()
	for _, o := range observers {
		device.Observable().Detach(o)
	}
	return list.New()
}
