package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/watchdog"
	"genx-go/logger"
	"genx-go/message"
)

//NewWaitingLocationMessageObserver ...
func NewWaitingLocationMessageObserver(_task interfaces.ITask) *WaitingLocationMessageObserver {
	return &WaitingLocationMessageObserver{
		task:     _task,
		watchdog: watchdog.NewWatchdog(_task.Device(), _task.Invoker().(interfaces.ILocationInvoker).LocationWatchdogCommands(_task), 10),
	}
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
		cList.PushBackList(observer.task.Invoker().DoneTask(observer.task))
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
