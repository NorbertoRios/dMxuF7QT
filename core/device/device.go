package device

import (
	"genx-go/configuration"
	"genx-go/logger"
	"genx-go/message"
	"genx-go/message/messagetype"
	"genx-go/parser"
	"genx-go/utils"
	"strings"
	"time"
)

//BuildDevice build new device
func BuildDevice(baseDevice *BaseDevice, onDeviceStateUpdated func(IDevice)) *Device {
	device := &Device{}
	baseDevice.TaskStorage.Device = device
	device.identity = baseDevice.Identity()
	device.parameter24 = baseDevice.parameter24
	device.parameter500 = baseDevice.parameter500
	device.TaskStorage = baseDevice.TaskStorage
	device.Sensors = baseDevice.Sensors
	device.updateBinaryReportParser()
	return device
}

//Device struct
type Device struct {
	*BaseDevice
	binaryReportParser *parser.GenxBinaryReportParser
	deviceStateUpdated func(IDevice)
	LatsFuelRequest    time.Time
	Lats1WireRequest   time.Time
}

func (device *Device) new24Parameter(new24 string) {
	if device.parameter24 == new24 {
		return
	}
	device.parameter24 = new24
	device.updateBinaryReportParser()
}

func (device *Device) updateBinaryReportParser() {
	param24 := strings.ReplaceAll(strings.Split(device.parameter24, "=")[1], ";", "")
	param24Columns := strings.Split(param24, ".")
	file := &utils.File{FilePath: "/configuration/initialize/ReportConfiguration.xml"}
	xmlProvider := configuration.ConstructXMLProvider(file)
	config, err := configuration.ConstructReportConfiguration(xmlProvider)
	if err == nil {
		logger.Error("[updateBinaryReportParser] Cant create binary message parser")
		return
	}
	fields := config.GetFieldsByIds(param24Columns)
	device.binaryReportParser = &parser.GenxBinaryReportParser{
		ReportFields: fields,
	}
}

func (device *Device) periodicalRequests(fields *[]configuration.Field) {

}



//NewRequiredParameter when configuration task ack device parameter
func (device *Device) NewRequiredParameter(key, value string) {
	switch key {
	case "24":
		{
			device.new24Parameter(value)
			return
		}
	case "500":
		{
			device.parameter500 = value
			return
		}
	}
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
