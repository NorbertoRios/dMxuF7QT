package device

import (
	"genx-go/connection"
	"genx-go/core/sensors"
	"genx-go/message"
	"genx-go/message/messagetype"
	"genx-go/parser"
	"genx-go/repository/models"
	"time"
)

//BuildBaseDevice returns base device
func BuildBaseDevice(identity string, channel *connection.UDPChannel, state []sensors.ISensor,
	loadCurrentConfig func(string, string) *models.ConfigurationModel, deviceSynchronized func(IDevice)) IDevice {
	device := &BaseDevice{
		Identity:   identity,
		Channel:    channel,
		Sensors:    state,
		LoadConfig: loadCurrentConfig,
	}
	device.TaskStorage = CounstructTaskStorage(device)
	return device
}

//BaseDevice unknown device
type BaseDevice struct {
	TaskStorage          *TaskStorage
	Sensors              []sensors.ISensor
	OnDeviceSynchronized func(IDevice)
	LoadConfig           func(string, string) *models.ConfigurationModel
	Identity             string
	Channel              *connection.UDPChannel
	LatsConfigLoadingTS  time.Time
	currentConfig        *models.ConfigurationModel
}

//OnLoadCurrentConfig when need to load current device config
func (device *BaseDevice) OnLoadCurrentConfig() *models.ConfigurationModel {
	if device.currentConfig == nil || time.Now().UTC().Sub(device.LatsConfigLoadingTS).Minutes() > 180 {
		device.currentConfig = device.LoadConfig(device.Identity, CurrentConfig)
	}
	return device.currentConfig
}

//OnLoadNonSendedConfig when need load non sended config
func (device *BaseDevice) OnLoadNonSendedConfig() *models.ConfigurationModel {
	return device.LoadConfig(device.Identity, Unsended)
}

//SendFacadeCallback send callback to facade
func (device *BaseDevice) SendFacadeCallback(callbackID string) {

}

//OnSynchronizationTaskCompleted when synchonization complete
func (device *BaseDevice) OnSynchronizationTaskCompleted() {
	device.OnDeviceSynchronized(device)
	return
}

//Send send command
func (device *BaseDevice) Send(message string) error {
	return device.Channel.Send(message)
}

//MessageArrived to device
func (device *BaseDevice) MessageArrived(rawMessage *message.RawMessage) {
	var msgParser parser.IParser
	switch rawMessage.MessageType {
	case messagetype.Parameter:
		{
			msgParser = parser.ConstructParametersMessageParser()
			break
		}
	case messagetype.Ack:
		{
			msgParser = parser.ConstructAckMesageParser()
			break
		}
	default:
		{
			return
		}
	}
	if message := msgParser.Parse(rawMessage); message != nil {
		device.processSystemMessage(message)
	}
}

func (device *BaseDevice) processSystemMessage(message interface{}) {
	device.TaskStorage.NewDeviceResponce(message)
}
