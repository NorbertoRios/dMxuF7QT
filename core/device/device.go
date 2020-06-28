package device

import (
	"genx-go/logger"
	"genx-go/configuration"
	"genx-go/message"
	"genx-go/message/messagetype"
	"genx-go/parser"
	"log"
)

func extractReportFields

//BuildDevice build new device
func BuildDevice(baseDevice *BaseDevice, onDeviceStateUpdated func(IDevice)) *Device {
	bReportParser := parser.BuildGenxBinaryReportParser(baseDevice.Parameter24())
	if bReportParser == nil {
		logger.Error("[BuildDevice] Cant create device. Binary report parser is nil")
		return nil
	}
	device := &Device{
		baseDevice:         baseDevice,
		deviceStateUpdated: onDeviceStateUpdated,
		binaryReportParser: bReportParser,
	}
}

//Device struct
type Device struct {
	reportFields       []*configuration.Field
	baseDevice         *BaseDevice
	binaryReportParser *parser.GenxBinaryReportParser
	deviceStateUpdated func(IDevice)
}



//MessageArrived on message arrived
func (device *Device) MessageArrived(rawMessage *message.RawMessage) {
	switch rawMessage.MessageType {

	case messagetype.Parameter:
	case messagetype.Ack:
		{
			device.baseDevice.MessageArrived(rawMessage)
			return
		}
	}
}

//OnSynchronizationTaskCompleted when synchonization complete
func (device *Device) OnSynchronizationTaskCompleted(param24, param500 string) {
	device.parameter24 = param24
	device.parameter500 = param500
	device.OnDeviceSynchronized(device)
	return
}