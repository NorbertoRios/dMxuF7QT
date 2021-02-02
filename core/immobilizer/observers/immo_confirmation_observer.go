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
		Watchdog: watchdog.NewWatchdog(_task.Device(), _task.Invoker().(interfaces.IImmoInvoker).WatchdogsCommands(_task, "DIAG HARDWARE"), 30),
	}
}

//ImmoConfitmationObserver ...
type ImmoConfitmationObserver struct {
	task     interfaces.ITask
	Watchdog *watchdog.Watchdog
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
			locationMessage := msg.(*message.Message)
			if commands := observer.checkSensorState(locationMessage.Sensors); commands != nil {
				return commands
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
	outNum = outNumDecorator.Index()
	for _, sens := range messgaeSensors {
		switch sens.(type) {
		case *sensors.Outputs:
			{
				relay := sens.(*sensors.Outputs).Relays
				if relay[outNum] == state {
					observer.Watchdog.Stop()
					return observer.task.Invoker().DoneTask(observer.task)
				}
			}
		}
	}
	return list.New()
}
