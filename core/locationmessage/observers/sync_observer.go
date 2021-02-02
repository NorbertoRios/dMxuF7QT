package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/watchdog"
	"genx-go/logger"
	"genx-go/message"
)

//NewSyncObserver ...
func NewSyncObserver(_task interfaces.ITask) *SyncObserver {
	return &SyncObserver{
		task:     _task,
		watchdog: watchdog.NewWatchdog(_task.Device(), _task.Invoker().(interfaces.ILocationProcessInvoker).SendDiagCommandAfterAnyMessage(_task), 10),
	}
}

//SyncObserver ...
type SyncObserver struct {
	task     interfaces.ITask
	watchdog *watchdog.Watchdog
}

//Task ...
func (observer *SyncObserver) Task() interfaces.ITask {
	return observer.task
}

//Attached ...
func (observer *SyncObserver) Attached() {
	logger.Logger().WriteToLog(logger.Info, "[SyncObserver] Successfuly attached")
	observer.watchdog.Start()
}

func (observer *SyncObserver) synchronized(value string) *list.List {
	cList := list.New()
	cList.PushBackList(observer.task.Invoker().(interfaces.ILocationProcessInvoker).DeviceSynchronized(value, observer.task.Device()))
	observer.watchdog.Stop()
	return cList
}

//Update ...
func (observer *SyncObserver) Update(msg interface{}) *list.List {
	switch msg.(type) {
	case *message.ParametersMessage:
		{
			paramMessage := msg.(*message.ParametersMessage)
			if value, f := paramMessage.Parameters["24"]; f {
				return observer.synchronized(value)
			}
		}
	case *message.AckMessage:
		{
			ackMessage := msg.(*message.ParametersMessage)
			if value, f := ackMessage.Parameters["24"]; f {
				return observer.synchronized(value)
			}
		}
	}
	return list.New()
}
