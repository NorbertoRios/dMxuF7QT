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
func BuildBaseDevice(identity string, channel *connection.UDPChannel, state *LastKnownDeviceState,
	loadConfig func(string, string) *models.ConfigurationModel, deviceSynchronized func(IDevice), publishMessage func(interface{})) IDevice {
	device := &BaseDevice{
		identity:             identity,
		Channel:              channel,
		Sensors:              state.Sensors,
		LoadConfig:           loadConfig,
		PublishMessage:       publishMessage,
		OnDeviceSynchronized: deviceSynchronized,
	}
	device.TaskStorage = CounstructTaskStorage(device)
	return device
}

//BaseDevice unknown device
type BaseDevice struct {
	parameter500          string
	parameter24           string
	TaskStorage           *TaskStorage
	PublishMessage        func(interface{})
	Sensors               []sensors.ISensor
	OnDeviceSynchronized  func(IDevice)
	LoadConfig            func(string, string) *models.ConfigurationModel
	identity              string
	Channel               *connection.UDPChannel
	lastActivityTimeStamp time.Time
}

//NewRequiredParameter when configuration task ack device parameter
func (device *BaseDevice) NewRequiredParameter(key, value string) {
	switch key {
	case "24":
		{
			device.parameter24 = value
			return
		}
	case "500":
		{
			device.parameter500 = value
			return
		}
	}
}

//LastActivityTimeStamp returns device last activity
func (device *BaseDevice) LastActivityTimeStamp() time.Time {
	return device.lastActivityTimeStamp
}

//Parameter500 returns 500 parameter
func (device *BaseDevice) Parameter500() string {
	return device.parameter500
}

//Parameter24 returns 24
func (device *BaseDevice) Parameter24() string {
	return device.parameter24
}

//Identity returns device identity
func (device *BaseDevice) Identity() string {
	return device.identity
}

//OnLoadCurrentConfig returns current config
func (device *BaseDevice) OnLoadCurrentConfig() *models.ConfigurationModel {
	return device.LoadConfig(device.identity, CurrentConfig)
}

//OnLoadNonSendedConfig when need load non sended config
func (device *BaseDevice) OnLoadNonSendedConfig() *models.ConfigurationModel {
	return device.LoadConfig(device.identity, Unsended)
}

//SendFacadeCallback send callback to facade
func (device *BaseDevice) SendFacadeCallback(callbackID string) {
}

//OnSynchronizationTaskCompleted when synchonization complete
func (device *BaseDevice) OnSynchronizationTaskCompleted() {
	device.OnDeviceSynchronized(device)
	return
}

//Send command
func (device *BaseDevice) Send(message string) error {
	return device.Channel.Send(message)
}

//MessageArrived on message arrived
func (device *BaseDevice) MessageArrived(rawMessage *message.RawMessage) {
	switch rawMessage.MessageType {

	case messagetype.Parameter:
		{
			device.processSystemMessage(parser.ConstructParametersMessageParser(), rawMessage)
			return
		}
	case messagetype.Ack:
		{
			device.processSystemMessage(parser.ConstructAckMesageParser(), rawMessage)
			return
		}
	}
}

func (device *BaseDevice) processSystemMessage(parser parser.IParser, rawMessage *message.RawMessage) {
	if message := parser.Parse(rawMessage); message != nil {
		device.PublishMessage(message)
		device.TaskStorage.NewDeviceResponce(message)
	}
}
