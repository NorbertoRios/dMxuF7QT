package main

import (
	"genx-go/configuration"
	"genx-go/connection"
	"genx-go/core/device"
	"genx-go/core/device_storage"
	"genx-go/core/immobilizer/request"
	"genx-go/logger"
	"genx-go/message"
	"genx-go/message/messagetype"
	"genx-go/parser"
	"genx-go/test/mock"
)

//NewGenxService ..
func NewGenxService() *GenxService {
	file := &mock.File{FilePath: "ReportConfiguration.xml"}
	xmlProvider := configuration.ConstructXMLProvider(file)
	config, _ := configuration.ConstructReportConfiguration(xmlProvider)
	return &GenxService{
		udpServer:     connection.ConstructUDPServer("", 10164),
		Configuration: config,
		factory:       message.CounstructRawMessageFactory(),
		storage:       device_storage.NewDeviceStorage(),
	}
}

//GenxService ...
type GenxService struct {
	Configuration *configuration.ReportConfiguration
	udpServer     *connection.UDPServer
	factory       *message.RawMessageFactory
	storage       *device_storage.DeviceStorage
}

//Run ...
func (service *GenxService) Run() {
	service.udpServer.OnNewPacket(service.onNewPacket)
	go service.udpServer.Listen()
}

func (service *GenxService) onNewPacket(channel *connection.UDPChannel, packet []byte) {
	rm := service.factory.BuildRawMessage(packet)
	service.parseMessage(rm, service.storage.Device(rm.Serial(), channel))
}

func (service *GenxService) startImmo(d *device.Device) {
	logger.Logger().WriteToLog(logger.Info, "Device founded")
	immo := d.ImmoStorage.Immobilizer(1, "high", d)
	data := &request.ChangeImmoStateRequest{
		SafetyOption: true,
		State:        "armed",
		Trigger:      "high",
	}
	data.Port = "OUT0"
	data.Identity = "genx_000003870006"
	immo.NewRequest(data)
}

func (service *GenxService) parseMessage(rawMessage *message.RawMessage, _device *device.Device) {
	switch rawMessage.MessageType {
	case messagetype.BinaryLocation:
		{
			fields := service.Configuration.GetFieldsByIds(_device.Param24)
			parser := &parser.GenxBinaryReportParser{
				ReportFields: fields,
			}
			messages, ack := parser.Parse(rawMessage)
			for _, msg := range messages {
				_device.MessageArrived(msg)
			}
			_device.Send(ack)
			break
		}
	case messagetype.Ack:
		{
			parser := parser.ConstructAckMesageParser()
			msg := parser.Parse(rawMessage)
			_device.MessageArrived(msg)
			break
		}
	case messagetype.DiagHardware:
		{
			parser := parser.BuildGenxHardwareMessageParser()
			msg := parser.Parse(rawMessage)
			_device.MessageArrived(msg)
			break
		}
	}
	if !immoStarted {
		immoStarted = true
		//time.Sleep(3 * time.Minute)
		service.startImmo(_device)
	}
}

var immoStarted bool = false
