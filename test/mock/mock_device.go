package mock

import (
	serviceConfiguration "genx-go/configuration"
	"genx-go/core/device"
	"genx-go/core/device/interfaces"
	"genx-go/core/location"
	"genx-go/core/peripherystorage"
	"genx-go/core/sensors"
	"genx-go/logger"
	"genx-go/parser"
	"strings"
	"sync"
)

var createdDevice interfaces.IDevice

//NewDevice ...
func NewDevice() interfaces.IDevice {
	param24 := "24=1.7.13.36.3.4.23.65.10.17.11.79.46.44.43.82.152.41.48.56.70.77.93.130;"
	param24 = strings.ReplaceAll(strings.Split(param24, "=")[1], ";", "")
	param24Columns := strings.Split(param24, ".")
	dev := &Device{}
	dev.Parameter24 = param24Columns
	dev.CurrentState = device.NewState(make(map[string]sensors.ISensor))
	dev.UDPChannel = &UDPChannel{}
	//dev.SerialNumber = "000003870006"
	dev.Mutex = &sync.Mutex{}
	dev.DeviceObservable = device.NewObservable()
	dev.ImmoStorage = peripherystorage.NewImmobilizerStorage()
	dev.LockStorage = peripherystorage.NewElectricLockStorage()
	dev.LocationProcess = location.New()
	createdDevice = dev
	return dev
}

//Device mock device
type Device struct {
	device.Device
	LastSentMessage       string
	LastPublishedToRabbit string
}

//Parser ...
func (device *Device) Parser() parser.IParser {
	if device.DeviceParser == nil {
		file := NewFile("..", "/ReportConfiguration.xml")
		xmlProvider := serviceConfiguration.ConstructXMLProvider(file)
		device.DeviceParser = parser.NewGenxBinaryReportParser(device.Parameter24, xmlProvider)
	}
	return device.DeviceParser
}

//Send ...
func (device *Device) Send(message interface{}) error {
	device.LastSentMessage = message.(string)
	return nil
}

//PushToRabbit ...
func (device *Device) PushToRabbit(message, destination string) {
	logger.Logger().WriteToLog(logger.Info, "Pushed ", message, "to ", destination)
	device.LastPublishedToRabbit = message
}
