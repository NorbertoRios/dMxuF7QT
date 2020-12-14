package observers

import (
	"container/list"

	"genx-go/core/device/interfaces"
	"genx-go/core/immobilizer/request"
	bRequest "genx-go/core/request"
	"genx-go/core/sensors"
	"genx-go/core/watchdog"
	"genx-go/logger"
	"genx-go/message"
)

//NewImmoConfitmationObserver ...
func NewImmoConfitmationObserver(_task interfaces.ITask) *ImmoConfitmationObserver {
	return &ImmoConfitmationObserver{
		task:     _task,
		Watchdog: watchdog.NewDiagImmoWatchdog(_task, 300),
	}
}

//ImmoConfitmationObserver ...
type ImmoConfitmationObserver struct {
	task     interfaces.ITask
	Watchdog *watchdog.DiagImmoWatchdog
}

//Attached ...
func (observer *ImmoConfitmationObserver) Attached() {
	logger.Logger().WriteToLog(logger.Info, "[ImmoConfitmationObserver] Successfuly attached")
	observer.Watchdog.Start()
}

//Task ...
func (observer *ImmoConfitmationObserver) Task() interfaces.ITask {
	return observer.task
}

//Update ...
func (observer *ImmoConfitmationObserver) Update(msg interface{}) *list.List {
	switch msg.(type) {
	case *message.LocationMessage:
		{
			locationMessage := msg.(*message.LocationMessage)
			for _, simpleMessage := range locationMessage.Messages {
				if commands := observer.checkSensorState(simpleMessage.Sensors); commands != nil {
					return commands
				}
				continue
			}
		}
	case *message.HardwareMessage:
		{
			return observer.checkSensorState(msg.(*message.HardwareMessage).Sensors)
		}
	}
	return list.New()
}

func (observer *ImmoConfitmationObserver) checkSensorState(messgaeSensors []sensors.ISensor) *list.List {
	var state byte
	var outNum int
	req := observer.task.Request().(*request.ChangeImmoStateRequest)
	stateDecorator := &request.ShouldStateByte{Data: req}
	outNumDecorator := &bRequest.OutputNumber{Data: req.Port}
	state = stateDecorator.State()
	outNum = outNumDecorator.Index() - 1
	for _, sens := range messgaeSensors {
		switch sens.(type) {
		case *sensors.Relay:
			{
				relay := sens.(*sensors.Relay)
				if relay.ID == outNum && relay.State == state {
					observer.Watchdog.Stop()
					return observer.task.Invoker().DoneTask(observer.task)
				}
			}
		}
	}
	return list.New()
}
