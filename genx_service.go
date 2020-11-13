package main

import (
	"genx-go/configuration"
	"genx-go/connection"
	"genx-go/core/device"
	"genx-go/core/device_storage"
	"genx-go/core/immobilizer/request"
	lockRequest "genx-go/core/lock/request"
	"genx-go/logger"
	"genx-go/message"
	"genx-go/message/messagetype"
	"genx-go/parser"
	"genx-go/test/mock"
	"net"
	"time"
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
	if !immoStarted {
		immoStarted = true
		adr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:15123")
		service.startLock(service.storage.Device("000003870006", connection.ConstructUDPChannel(adr, service.udpServer)))
	}
}

func (service *GenxService) onNewPacket(channel *connection.UDPChannel, packet []byte) {
	rm := service.factory.BuildRawMessage(packet)
	service.parseMessage(rm, service.storage.Device(rm.Serial(), channel))
}

func (service *GenxService) startLock(d *device.Device) {
	logger.Logger().WriteToLog(logger.Info, "Device founded")
	lock := d.ElectricLock(1)
	exT := time.Now().UTC().Add(6 * time.Minute)
	data := &lockRequest.UnlockRequest{
		ExpirationTime: exT.Format("2006-01-02T15:04:05Z"),
		TimeToPulse:    3,
	}
	data.Port = "OUT0"
	data.Identity = "genx_000003870006"
	lock.NewRequest(data)
}

func (service *GenxService) startImmo(d *device.Device) {
	logger.Logger().WriteToLog(logger.Info, "Device founded")
	immo := d.Immobilizer(1, "high")
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
}

var immoStarted bool = false
