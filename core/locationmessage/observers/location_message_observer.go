package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/logger"
	"genx-go/message"
)

//NewLocationMessageObserver ...
func NewLocationMessageObserver(_task interfaces.ITask) *LocationMessageObserver {
	return &LocationMessageObserver{
		task: _task,
	}
}

//LocationMessageObserver ...
type LocationMessageObserver struct {
	task interfaces.ITask
}

//Task ...
func (observer *LocationMessageObserver) Task() interfaces.ITask {
	return observer.task
}

//Attached ...
func (observer *LocationMessageObserver) Attached() {
	logger.Logger().WriteToLog(logger.Info, "[LocationMessageObserver] Successfuly attached")
}

//Update ...
func (observer *LocationMessageObserver) Update(msg interface{}) *list.List {
	switch msg.(type) {
	case *message.Message:
		{
			locationMessage := msg.(*message.Message)
			observer.task.Device().NewState(locationMessage.Sensors)
			logger.Logger().WriteToLog(logger.Info, "[LocationMessageObserver | Update] New state for device ", locationMessage.Identity)
		}
	}
	return list.New()
}
