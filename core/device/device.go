package device

import (
	"container/list"
	"encoding/json"
	connInterfaces "genx-go/connection/interfaces"
	"genx-go/core/configuration"
	"genx-go/core/device/interfaces"
	"genx-go/core/location"
	"genx-go/core/peripherystorage"
	"genx-go/core/sensors"
	"genx-go/logger"
	"genx-go/message"
	"genx-go/parser"
	"sync"
	"time"
)

//NewActivityLessDevice ...
func NewActivityLessDevice(_channel connInterfaces.IChannel) interfaces.IDevice {
	return NewDevice([]string{}, make(map[string]sensors.ISensor), _channel)
}

//NewDevice ...
func NewDevice(_param24 []string, _sensors map[string]sensors.ISensor, _channel connInterfaces.IChannel) interfaces.IDevice {
	// device := &Device{
	// 	Parameter24:         _param24,
	// 	LastLocationMessage: &message.LocationMessage{},
	// 	DeviceObservable:    NewObservable(),
	// 	LastStateUpdateTime: time.Now().UTC(),
	// 	CurrentState:        NewState(_sensors),
	// 	Mutex:               &sync.Mutex{},
	// 	UDPChannel:          _channel,
	// 	ImmoStorage:         peripherystorage.NewImmobilizerStorage(),
	// 	LockStorage:         peripherystorage.NewElectricLockStorage(),
	// }
	// //device.DeviceObservable.Attach()
	// return device
	return nil
}

//Device struct
type Device struct {
	LastLocationMessage *message.Message
	DeviceObservable    *Observable
	LastStateUpdateTime time.Time
	LastActivity        time.Time
	Config              interfaces.IProcess
	CurrentState        *State
	DeviceParser        parser.IParser
	ImmoStorage         *peripherystorage.ImmobilizerStorage
	LockStorage         *peripherystorage.ElectricLockStorage
	UDPChannel          connInterfaces.IChannel
	Mutex               *sync.Mutex
	LocationProcess     interfaces.IProcess
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

//MessageArrived new message
func (device *Device) MessageArrived(msg interface{}) *list.List {
	device.LastActivity = time.Now().UTC()
	bMessage, bErr := json.Marshal(msg)
	if bErr != nil {
		logger.Logger().WriteToLog(logger.Error, "[Device | MessageArrived] Error while marshaling incoming message. Error: ", bErr)
	} else {
		logger.Logger().WriteToLog(logger.Info, "[Device | MessageArrived] New message arrived. Message: ", string(bMessage))
	}
	return device.DeviceObservable.Notify(msg)
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

//Immobilizer ...
func (device *Device) Immobilizer(index int, trigger string) interfaces.IImmobilizer {
	return device.ImmoStorage.Immobilizer(index, trigger, device)
}

//Observable returns device Observable
func (device *Device) Observable() interfaces.IObservable {
	return device.DeviceObservable
}

//NewState ...
func (device *Device) NewState(messageSensors []sensors.ISensor) {
	device.LastStateUpdateTime = time.Now().UTC()
	device.CurrentState = NewSensorState(device.CurrentState, messageSensors)
}

//NewChannel ...
func (device *Device) NewChannel(_channel connInterfaces.IChannel) {
	device.UDPChannel = _channel
}

//Configuration ..
func (device *Device) Configuration() interfaces.IProcess {
	if device.Config == nil {
		device.Config = configuration.NewConfiguration()
	}
	return device.Config
}

//LocationRequest ..
func (device *Device) LocationRequest() interfaces.IProcess {
	if device.LocationProcess == nil {
		device.LocationProcess = location.New()
	}
	return device.LocationProcess
}

//LocationMessageProcess ...
func (device *Device) LocationMessageProcess() interfaces.ILocationMessageProcess {
	return nil
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
