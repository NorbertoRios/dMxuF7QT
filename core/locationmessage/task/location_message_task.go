package task

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/filter"
	"genx-go/core/invoker"
	"genx-go/core/locationmessage/observers"
	coreCommand "genx-go/core/observers"
	"time"
)

//NewLocationMessageTask ...
func NewLocationMessageTask(_process interfaces.ILocationMessageProcess, _device interfaces.IDevice) *LocationMessageTask {
	t := &LocationMessageTask{}
	t.bornTime = time.Now().UTC()
	t.device = _device
	t.invoker = invoker.NewLocationProcessInvoker(_process)
	return t
}

//LocationMessageTask ...
type LocationMessageTask struct {
	BaseTask
}

//Commands ..
func (task *LocationMessageTask) Commands() *list.List { //При создании таска должна повесить обзервера который будет крутить состояние девайса
	cList := list.New()
	cList.PushBack(coreCommand.NewAttachObserverCommand(observers.NewLocationMessageObserver(task)))
	return cList
}

//Observers returns task's observer
func (task *LocationMessageTask) Observers() []interfaces.IObserver {
	filter := filter.NewObserversFilter(task.device.Observable())
	return filter.Extract(task)
}
