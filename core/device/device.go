package device

import (
	"container/list"
	"genx-go/connection"
	"genx-go/core/device/interfaces"
	"genx-go/core/peripherystorage"
	"genx-go/core/sensors"
	"genx-go/logger"
	"sync"
	"time"
)

//NewDevice ...
func NewDevice(_serial string, _param24 []string, _channel connection.IChannel) interfaces.IDevice {
	return &Device{
		Param24:             _param24,
		CurrentState:        make(map[sensors.ISensor]time.Time),
		UDPChannel:          _channel,
		SerialNumber:        _serial,
		LastStateUpdateTime: time.Now().UTC(),
		Mutex:               &sync.Mutex{},
		DeviceObservable:    NewObservable(),
		ImmoStorage:         peripherystorage.NewImmobilizerStorage(),
		LockStorage:         peripherystorage.NewElectricLockStorage(),
	}
}

//Device struct
type Device struct {
	Param24             []string
	DeviceObservable    *Observable
	LastStateUpdateTime time.Time
	ImmoStorage         *peripherystorage.ImmobilizerStorage
	LockStorage         *peripherystorage.ElectricLockStorage
	UDPChannel          connection.IChannel
	Mutex               *sync.Mutex
	SerialNumber        string
	CurrentState        map[sensors.ISensor]time.Time
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

//ElectricLock ..
func (device *Device) ElectricLock(index int) interfaces.ILock {
	return device.LockStorage.ElectricLock(index, device)
}

//State returns device current state
func (device *Device) State() map[sensors.ISensor]time.Time {
	device.Mutex.Lock()
	defer device.Mutex.Unlock()
	return device.CurrentState
}

//PushToRabbit ...
func (device *Device) PushToRabbit(message, destination string) {
}

//MessageArrived new message
func (device *Device) MessageArrived(msg interface{}) {
	commands := device.DeviceObservable.Notify(msg)
	device.ProccessCommands(commands)
}

//Immobilizer ...
func (device *Device) Immobilizer(index int, trigger string) interfaces.IImmobilizer {
	return device.ImmoStorage.Immobilizer(index, trigger, device)
}

//Observable returns device Observable
func (device *Device) Observable() interfaces.IObservable {
	return device.DeviceObservable
}

//ProccessCommands process commands
func (device *Device) ProccessCommands(commands *list.List) {
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
