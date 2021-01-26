package task

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/filter"
	"genx-go/core/invoker"
	"genx-go/core/locationmessage/observers"
	"time"
)

//NewSyncTask ...
func NewSyncTask(_process interfaces.ILocationMessageProcess, _device interfaces.IDevice) *SyncTask {
	t := &SyncTask{}
	t.bornTime = time.Now().UTC()
	t.device = _device
	t.invoker = invoker.NewLocationProcessInvoker(_process)
	return t
}

//SyncTask ...
type SyncTask struct {
	BaseTask
}

//Observers returns task's observer
func (task *SyncTask) Observers() []interfaces.IObserver {
	filter := filter.NewObserversFilter(task.device.Observable())
	return filter.Extract(task)
}

//Commands ...
func (task *SyncTask) Commands() *list.List { //Таска должна отправить комманду для опроса 24 параметра + повесить обзервер на эту команду
	cList := list.New()
	cList.PushBack(observers.NewSyncObserver(task))
	return cList
}
