package mock

import (
	"genx-go/core/device"
	"genx-go/core/device/interfaces"
	"genx-go/core/peripherystorage"
	"genx-go/core/sensors"
	"genx-go/logger"
	"sync"
	"time"
)

var createdDevice interfaces.IDevice

//NewDevice ...
func NewDevice() interfaces.IDevice {
	dev := &Device{}
	dev.Param24 = []string{}
	dev.CurrentState = make(map[sensors.ISensor]time.Time)
	dev.UDPChannel = &UDPChannel{}
	dev.SerialNumber = "000003870006"
	dev.Mutex = &sync.Mutex{}
	dev.DeviceObservable = device.NewObservable()
	dev.ImmoStorage = peripherystorage.NewImmobilizerStorage()
	dev.LockStorage = peripherystorage.NewElectricLockStorage()
	createdDevice = dev
	return dev
}

//Device mock device
type Device struct {
	device.Device
	LastSentMessage       string
	LastPublishedToRabbit string
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
