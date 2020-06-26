package device

import (
	"genx-go/message"
	"genx-go/message/messagetype"
	"genx-go/parser"
)

//BuildDevice build new device
func BuildDevice(device *BaseDevice) *Device {
	return nil
}

//Device struct
type Device struct {
	BaseDevice
	binaryReportParser *parser.GenxBinaryReportParser
	deviceStateUpdated func(IDevice)
	publishMessage     func(interface{})
}

//MessageArrived on message arrived
func (device *Device) MessageArrived(rawMessage *message.RawMessage) {
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

func (device *Device) processSystemMessage(parser parser.IParser, rawMessage *message.RawMessage) {
	if message := parser.Parse(rawMessage); message != nil {
		device.publishMessage(message)
		device.TaskStorage.NewDeviceResponce(message)
	}
}
