package device

import (
	"container/list"
	"genx-go/connection"
	"genx-go/core/device/interfaces"
	"genx-go/core/immostorage"
	"genx-go/core/sensors"
	"genx-go/logger"
	"sync"
	"time"
)

//Device struct
type Device struct {
	Param24             []string
	Observable          *Observable
	LastStateUpdateTime time.Time
	ImmoStorage         *immostorage.ImmobilizerStorage
	UDPChannel          connection.IChannel
	Mutex               *sync.Mutex
	SerialNumber        string
	fsdfsdfsd           func()
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
	commands := device.Observable.Notify(msg)
	device.ProccessCommands(commands)
}

//Immobilizer ...
func (device *Device) Immobilizer(index int, trigger string) interfaces.IImmobilizer {
	return device.ImmoStorage.Immobilizer(index, trigger, device)
}

//GetObservable returns device Observable
func (device *Device) GetObservable() interfaces.IObservable {
	return device.Observable
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
