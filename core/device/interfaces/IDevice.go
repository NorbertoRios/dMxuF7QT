package interfaces

import (
	"container/list"
	"genx-go/connection/interfaces"
	"genx-go/core/sensors"
	"genx-go/message"
)

//IDevice device interface
type IDevice interface {
	// Observable() IObservable
	// Send(interface{}) error
	// ProcessCommands(*list.List)
	// State() map[sensors.ISensor]time.Time
	// Immobilizer(int, string) IImmobilizer
	// ElectricLock(int) IProcess
	// MessageArrived(interface{}) *list.List
	// LocationRequest() IProcess
	// Configuration() IProcess
	// Parser() parser.IParser

	//LastDeviceState
	// LastLocationMessage() *message.LocationMessage
	// Observable() IObservable
	// LastDeviceMessage() *message.Message

	// CurrentDeviceState() []sensors.ISensor
	// Immobilizer(int, string) IImmobilizer
	// ElectricLock(int) IProcess
	// ProcessCommands(*list.List)
	// State() []sensors.ISensor
	// MessageArrived(interface{}) *list.List
	// LocationRequest() IProcess
	// Configuration() IProcess

	Send(interface{}) error
	NewChannel(interfaces.IChannel)
	Configuration() IProcess
	LastDeviceLocationMessage() *message.LocationMessage
	LocationRequest() IProcess
	ElectricLock(int) IProcess
	State() map[string]sensors.ISensor
	MessageArrived(interface{}) *list.List
	Immobilizer(int, string) IImmobilizer
	Observable() IObservable
	ProcessCommands(*list.List)
}
