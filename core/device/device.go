package device

import (
	"container/list"
	"fmt"
	serviceConfiguration "genx-go/configuration"
	"genx-go/connection"
	"genx-go/core/configuration"
	"genx-go/core/device/interfaces"
	"genx-go/core/location"
	"genx-go/core/peripherystorage"
	"genx-go/core/sensors"
	"genx-go/logger"
	"genx-go/parser"
	"genx-go/types"
	"os"
	"sync"
	"time"
)

//NewDevice ...
func NewDevice(_serial string, _param24 []string, _channel connection.IChannel) interfaces.IDevice {
	device := &Device{
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
	device.LocationTask = location.New(device)
	return device
}

//Device struct
type Device struct {
	Param24             []string
	DeviceObservable    *Observable
	LastStateUpdateTime time.Time
	Config              interfaces.IConfiguration
	DeviceParser        parser.IParser
	ImmoStorage         *peripherystorage.ImmobilizerStorage
	LockStorage         *peripherystorage.ElectricLockStorage
	UDPChannel          connection.IChannel
	Mutex               *sync.Mutex
	LocationTask        interfaces.ILocationRequest
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

//Parser ...
func (device *Device) Parser() parser.IParser {
	if device.DeviceParser == nil {
		file := &types.File{FilePath: fmt.Sprintf("%v/ReportConfiguration.xml", os.Args[0])}
		xmlProvider := serviceConfiguration.ConstructXMLProvider(file)
		device.DeviceParser = parser.NewGenxBinaryReportParser(device.Param24, xmlProvider)
	}
	return device.DeviceParser
}

//Configuration ..
func (device *Device) Configuration() interfaces.IConfiguration {
	if device.Config == nil {
		device.Config = configuration.NewConfiguration(device)
	}
	return device.Config
}

//LocationRequest ..
func (device *Device) LocationRequest() interfaces.ILocationRequest {
	return device.LocationTask
}

//ElectricLock ..
func (device *Device) ElectricLock(index int) interfaces.IProcess {
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
func (device *Device) MessageArrived(msg interface{}) *list.List {
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
