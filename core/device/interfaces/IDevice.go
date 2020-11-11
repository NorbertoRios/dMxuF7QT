package interfaces

import (
	"container/list"
	"genx-go/core/sensors"
	"time"
)

//IDevice device interface
type IDevice interface {
	GetObservable() IObservable
	Send(interface{}) error
	ProccessCommands(*list.List)
	PushToRabbit(string, string)
	State() map[sensors.ISensor]time.Time
	Immobilizer(int, string) IImmobilizer
	MessageArrived(interface{})
}
