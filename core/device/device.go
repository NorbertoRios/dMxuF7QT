package device

import (
	"container/list"
	serviceConfig "genx-go/configuration"
	connInterfaces "genx-go/connection/interfaces"
	"genx-go/core/configuration"
	"genx-go/core/device/interfaces"
	"genx-go/core/location"
	"genx-go/core/peripherystorage"
	"genx-go/core/sensors"
	"genx-go/logger"
	"genx-go/message"
	"genx-go/parser"
	"genx-go/types"
	"sync"
	"time"
)

//NewDevice ...
func NewDevice(_param24 []string, _sensors map[string]sensors.ISensor, _channel connInterfaces.IChannel) interfaces.IDevice {
	return &Device{
		Parameter24:         _param24,
		LastLocationMessage: &message.LocationMessage{},
		DeviceObservable:    NewObservable(),
		LastStateUpdateTime: time.Now().UTC(),
		CurrentState:        NewState(_sensors),
		Mutex:               &sync.Mutex{},
		UDPChannel:          _channel,
		ImmoStorage:         peripherystorage.NewImmobilizerStorage(),
		LockStorage:         peripherystorage.NewElectricLockStorage(),
	}
}

//Device struct
type Device struct {
	Parameter24         []string
	LastLocationMessage *message.LocationMessage
	DeviceObservable    *Observable
	LastStateUpdateTime time.Time
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

//Parser ...
func (device *Device) Parser() parser.IParser {
	if device.DeviceParser != nil {
		file := types.NewFile("/config/initializer/ReportConfiguration.xml")
		provider := serviceConfig.ConstructXMLProvider(file)
		device.DeviceParser = parser.NewGenxBinaryReportParser(device.Parameter24, provider)
	}
	return device.DeviceParser
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

//LastDeviceLocationMessage ..
func (device *Device) LastDeviceLocationMessage() *message.LocationMessage {
	return device.LastLocationMessage
}

//ElectricLock ..
func (device *Device) ElectricLock(index int) interfaces.IProcess {
	return device.LockStorage.ElectricLock(index, device)
}

//Ack ...
func (device *Device) Ack() {
	device.Send(device.LastLocationMessage.Ack)
}

//State returns device current state
func (device *Device) State() map[string]sensors.ISensor {
	device.Mutex.Lock()
	defer device.Mutex.Unlock()
	return device.CurrentState.State()
}

func (device *Device) handleLocationMessage(msg *message.LocationMessage) *list.List {
	commands := list.New()
	device.LastLocationMessage = msg
	for _, m := range msg.Messages {
		device.CurrentState = NewSensorState(device.CurrentState, m.Sensors)
		commands.PushBackList(device.DeviceObservable.Notify(m))
	}
	return commands
}

//MessageArrived new message
func (device *Device) MessageArrived(msg interface{}) *list.List {
	device.LastStateUpdateTime = time.Now().UTC()
	switch msg.(type) {
	case *message.LocationMessage:
		{
			return device.handleLocationMessage(msg.(*message.LocationMessage))
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
