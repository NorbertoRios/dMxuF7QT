package mock

import (
	"genx-go/core/device"
	"genx-go/core/device/interfaces"
	"genx-go/core/immostorage"
	"genx-go/logger"
	"sync"
)

var createdDevice interfaces.IDevice

//NewMockDevice returns new mock device
func NewMockDevice() interfaces.IDevice {
	d := &Device{}
	d.Observable = device.NewObservable()
	d.UDPChannel = &UDPChannel{}
	d.Mutex = &sync.Mutex{}
	d.ImmoStorage = immostorage.NewImmobilizerStorage()
	d.SerialNumber = "000003870006"
	createdDevice = d
	return d
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
