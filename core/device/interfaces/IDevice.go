package interfaces

import (
	"container/list"
	"genx-go/connection/interfaces"
	"genx-go/core/sensors"
	"genx-go/parser"
)

//IDevice device interface
type IDevice interface {
	Send(interface{}) error
	NewChannel(interfaces.IChannel)
	Configuration() IProcess
	LocationRequest() IProcess
	ElectricLock(int) IProcess
	State() map[string]sensors.ISensor
	MessageArrived(interface{}) *list.List
	Immobilizer(int, string) IImmobilizer
	Observable() IObservable
	ProcessCommands(*list.List)
	LocationMessageProcess() ILocationMessageProcess
	NewState([]sensors.ISensor)
	Parser() parser.IParser
}
