package device_storage

import (
	"genx-go/connection"
	"genx-go/core/device"
	"genx-go/core/observers"
	"strings"
	"sync"
)

//NewDeviceStorage ..
func NewDeviceStorage() *DeviceStorage {
	return &DeviceStorage{
		mutex:            &sync.Mutex{},
		DeviceCollection: make(map[string]*device.Device, 0),
	}
}

//DeviceStorage ..
type DeviceStorage struct {
	mutex            *sync.Mutex
	DeviceCollection map[string]*device.Device
}

//Device ...
func (storage *DeviceStorage) Device(serial string, channel connection.IChannel) *device.Device {
	var d *device.Device
	var f bool
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	if d, f = storage.DeviceCollection[serial]; f {
		d.UDPChannel = channel
		return d
	}
	param24 := "24=1.7.13.36.3.4.23.65.10.17.11.79.46.44.43.41.48.130;"
	param24 = strings.ReplaceAll(strings.Split(param24, "=")[1], ";", "")
	param24Columns := strings.Split(param24, ".")
	d = device.NewDevice(serial, param24Columns, channel).(*device.Device)
	d.Observable().Attach(&observers.ConsoleTestObserver{})
	storage.DeviceCollection[serial] = d
	return d
}
