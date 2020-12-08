package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/observers"
	"genx-go/core/watchdog"
	"genx-go/logger"
	"genx-go/message"
)

//NewWaitingLocationMessageObserver ...
func NewWaitingLocationMessageObserver(_task interfaces.ITask) *WaitingLocationMessageObserver {
	observer := &WaitingLocationMessageObserver{
		task: _task,
	}
	cList := list.New()
	cList.PushBack(observers.NewDetachObserverCommand(observer))
	cList.PushBack((NewSendLocationRequest(_task)))
	wd := watchdog.NewWatchdog(cList, _task.Device(), 300)
	observer.watchdog = wd
	return observer
}

//WaitingLocationMessageObserver ...
type WaitingLocationMessageObserver struct {
	task     interfaces.ITask
	watchdog *watchdog.Watchdog
}

//Update ...
func (observer *WaitingLocationMessageObserver) Update(msg interface{}) *list.List {
	cList := list.New()
	if _, f := msg.(*message.LocationMessage); f {
		observer.watchdog.Stop()
		observer.task.Done()
	}
	return cList
}

//Task ...
func (observer *WaitingLocationMessageObserver) Task() interfaces.ITask {
	return observer.task
}

//Attached ...
func (observer *WaitingLocationMessageObserver) Attached() {
	logger.Logger().WriteToLog(logger.Info, "[WaitingLocationMessageObserver | Attached] WaitingLocationMessageObserver successfuly attached")
	observer.watchdog.Start()
}
