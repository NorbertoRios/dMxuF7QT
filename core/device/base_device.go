package device

import (
	"genx-go/connection"
	"genx-go/core/sensors"
	"genx-go/message"
	"genx-go/message/messagetype"
	"genx-go/parser"
	"genx-go/repository/models"
	"log"
	"time"
)

//BuildBaseDevice returns base device
func BuildBaseDevice(identity string, state []sensors.ISensor, currentConfig *models.ConfigurationModel) IDevice {
	device := &BaseDevice{
		Sensors:       state,
		CurrentConfig: currentConfig,
	}
	defer device.resyncCron()
	return device
}

//BaseDevice unknown device
type BaseDevice struct {
	Sensors              []sensors.ISensor
	ConfigurationTask    *ConfigurationTask
	SynchronizarionTask  *SynchronizarionTask
	OnDeviceSynchronized func(*BaseDevice)
	Identity             string
	CurrentConfig        *models.ConfigurationModel
	Channel              *connection.UDPChannel
	LastSynchronizarion  time.Time
}

//OnSynchronizationTaskCompleted when synchonization complete
func (device *BaseDevice) OnSynchronizationTaskCompleted(needResendConfig bool) {
	if !needResendConfig {
		device.LastSynchronizarion = time.Now().UTC()
		device.OnDeviceSynchronized(device)
		return
	}
	if device.ConfigurationTask == nil {
		device.ConfigurationTask = BuildConfigurationTask(device, device.CurrentConfig)
		device.ConfigurationTask.Execute()
	}
}

//Config returns current device config (string)
func (device *BaseDevice) Config() string {
	return device.CurrentConfig.Command
}

func (device *BaseDevice) resyncCron() {
	ticker := time.NewTicker(300 * time.Second)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("[BuildBaseDevice] Recovered in recync cron:", r)
			}
		}()
		for {
			select {
			case <-ticker.C:
				device.Synchronize()
			}
		}
	}()
}

//Send send command
func (device *BaseDevice) Send(message string) error {
	return device.Channel.Send(message)
}

//OnConfigTaskCompleted when config task completed
func (device *BaseDevice) OnConfigTaskCompleted() {
	device.Synchronize()
}

//Synchronize run device Synchronization
func (device *BaseDevice) Synchronize() {
	device.SynchronizarionTask = ConstructSynchronizarionTask(device)
	device.SynchronizarionTask.Execute()
}

//MessageArrived to device
func (device *BaseDevice) MessageArrived(rawMessage *message.RawMessage) {
	switch rawMessage.MessageType {
	case messagetype.Parameter:
		{
			parser := parser.BuildParametersMessageParser()
			if message := parser.Parse(rawMessage); message != nil {
				device.processParametersMessage(message)
			}
		}
	case messagetype.Ack:
		{
			parser := parser.ConstructAckMesageParser()
			if message := parser.Parse(rawMessage); message != nil {
				device.processAckMessage(message)
			}
		}
	}
}

func (device *BaseDevice) processParametersMessage(message *message.ParametersMessage) {
	if device.SynchronizarionTask != nil {
		device.SynchronizarionTask.OnParametersRecieved(message)
	}
}

func (device *BaseDevice) processAckMessage(message *message.AckMessage) {
	if device.ConfigurationTask != nil {
		device.ConfigurationTask.AckArrived(message)
	}
}
