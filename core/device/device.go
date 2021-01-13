package device

import (
	"container/list"
	connInterfaces "genx-go/connection/interfaces"
	"genx-go/core/configuration"
	"genx-go/core/device/interfaces"
	"genx-go/core/peripherystorage"
	"genx-go/core/sensors"
	"genx-go/logger"
	"genx-go/message"
	"genx-go/parser"
	"sync"
	"time"
)

//NewDevice ...
// func NewDevice(_serial string, _param24 []string, _channel connInterfaces.IChannel) interfaces.IDevice {
// 	device := &Device{
// 		Param24:             _param24,
// 		CurrentState:        make(map[sensors.ISensor]time.Time),
// 		UDPChannel:          _channel,
// 		SerialNumber:        _serial,
// 		LastStateUpdateTime: time.Now().UTC(),
// 		Mutex:               &sync.Mutex{},
// 		DeviceObservable:    NewObservable(),
// 		ImmoStorage:         peripherystorage.NewImmobilizerStorage(),
// 		LockStorage:         peripherystorage.NewElectricLockStorage(),
// 	}
// 	device.LocationTask = location.New(device)
// 	return device
// }

//Device struct
type Device struct {
	lastLocationMessage *message.LocationMessage
	DeviceObservable    *Observable
	LastStateUpdateTime time.Time
	Config              interfaces.IProcess
	CurrentState        *State
	DeviceParser        parser.IParser
	ImmoStorage         *peripherystorage.ImmobilizerStorage
	LockStorage         *peripherystorage.ElectricLockStorage
	UDPChannel          connInterfaces.IChannel
	Mutex               *sync.Mutex
	LocationTask        interfaces.IProcess
	SerialNumber        string
}

//Send send command to device
func (device *Device) Send(message interface{}) error {
	err := device.UDPChannel.Send(message)
	if err != nil {
		logger.Logger().WriteToLog(logger.Error, "[Device | Send] Error:", err)
		return err
	}
	return nil
}

//NewChannel ...
func (device *Device) NewChannel(_channel connInterfaces.IChannel) {
	device.UDPChannel = _channel
}

//Configuration ..
func (device *Device) Configuration() interfaces.IProcess {
	if device.Config == nil {
		device.Config = configuration.NewConfiguration(device)
	}
	return device.Config
}

//LastDeviceMessage ..
func (device *Device) LastDeviceMessage() *message.LocationMessage {
	return device.lastLocationMessage
}

//LocationRequest ..
func (device *Device) LocationRequest() interfaces.IProcess {
	return device.LocationTask
}

//ElectricLock ..
func (device *Device) ElectricLock(index int) interfaces.IProcess {
	return device.LockStorage.ElectricLock(index, device)
}

//State returns device current state
func (device *Device) State() map[string]sensors.ISensor {
	device.Mutex.Lock()
	defer device.Mutex.Unlock()
	return device.CurrentState.State()
}

func (device *Device) handleLocationMessage(msg *message.LocationMessage) {
	device.lastLocationMessage = msg
	for _, m := range msg.Messages {
		device.CurrentState = NewSensorState(device.CurrentState, m.Sensors)
		device.DeviceObservable.Notify(m)
	}
}

//MessageArrived new message
func (device *Device) MessageArrived(msg interface{}) *list.List {
	device.LastStateUpdateTime = time.Now().UTC()
	switch msg.(type) {
	case *message.LocationMessage:
		{
			return device.handleLocationMessage(msg)
		}
	}
	return device.DeviceObservable.Notify(msg)
}

//Immobilizer ...
func (device *Device) Immobilizer(index int, trigger string) interfaces.IImmobilizer {
	return device.ImmoStorage.Immobilizer(index, trigger, device)
}

//Observable returns device Observable
func (device *Device) Observable() interfaces.IObservable {
	return device.DeviceObservable
}

//ProcessCommands process commands
func (device *Device) ProcessCommands(commands *list.List) {
	device.Mutex.Lock()
	defer device.Mutex.Unlock()
	for commands.Len() > 0 {
		cmd := commands.Front()
		command, valid := cmd.Value.(interfaces.ICommand)
		if valid {
			nList := command.Execute(device)
			if nList != nil && nList.Len() > 0 {
				commands.PushFrontList(nList)
			}
			commands.Remove(cmd)
		}
	}
}
