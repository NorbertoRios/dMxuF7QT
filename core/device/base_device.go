package device

import (
	"genx-go/connection"
	"genx-go/core/sensors"
	"genx-go/logger"
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
		quit:                 make(chan struct{}),
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
	quit                  chan struct{}
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
	lastSynchronization   time.Time
}

func (device *BaseDevice) synchronizationCron() {
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("[TaskStorage] Recovered in recync cron:", r)
			}
			for {
				select {
				case <-ticker.C:
					{
						if time.Now().UTC().Sub(device.lastSynchronization).Minutes() > 120 {
							logger.Info("[BaseDevice | synchronizationCron] Start sinchronization task for ", device.identity)
							device.CreateNewTask(SynchronizationTask, "", nil)
						}
					}
				case <-device.quit:
					{
						ticker.Stop()
						return
					}
				}
			}
		}()
	}()
}

//OnDeviceRemoving on device removing
func (device *BaseDevice) OnDeviceRemoving() {
	close(device.quit)
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

//OnLoadConfig returns config
func (device *BaseDevice) OnLoadConfig(identity, configType string) *models.ConfigurationModel {
	return device.LoadConfig(device.identity, configType)
}

//SendFacadeCallback send callback to facade
func (device *BaseDevice) SendFacadeCallback(callbackID string) {
}

//OnSynchronizationTaskCompleted when synchonization complete
func (device *BaseDevice) OnSynchronizationTaskCompleted() {
	device.OnDeviceSynchronized(device)
	return
}

//CreateNewTask new task for device
func (device *BaseDevice) CreateNewTask(taskType, callbackID string, onCompleteCallback func(string)) {
	device.TaskStorage.createTask(taskType, callbackID, onCompleteCallback)
}

//Send command
func (device *BaseDevice) Send(message string) error {
	return device.Channel.Send(message)
}

//MessageArrived on message arrived
func (device *BaseDevice) MessageArrived(rawMessage *message.RawMessage) {
	device.lastActivityTimeStamp = time.Now().UTC()
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
