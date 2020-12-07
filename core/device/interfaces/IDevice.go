package interfaces

import (
	"container/list"
	"genx-go/core/sensors"
	"time"
)

//IDevice device interface
type IDevice interface {
	Observable() IObservable
	Send(interface{}) error
	ProcessCommands(*list.List)
	PushToRabbit(string, string)
	State() map[sensors.ISensor]time.Time
	Immobilizer(int, string) IImmobilizer
	ElectricLock(int) ILock
	MessageArrived(interface{}) *list.List
	LocationRequest() ILocationRequest
	Configuration() IConfiguration
}
