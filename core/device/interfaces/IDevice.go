package interfaces

import (
	"container/list"
	"genx-go/core/sensors"
	"genx-go/parser"
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
	ElectricLock(int) IProcess
	MessageArrived(interface{}) *list.List
	LocationRequest() IProcess
	Configuration() IProcess
	Parser() parser.IParser
}
