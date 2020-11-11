package device_storage

import (
	"genx-go/connection"
	"genx-go/core/device"
	"genx-go/core/immostorage"
	"genx-go/core/observers"
	"genx-go/core/sensors"
	"strings"
	"sync"
	"time"
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
	d = &device.Device{
		Param24:             param24Columns,
		CurrentState:        make(map[sensors.ISensor]time.Time),
		SerialNumber:        serial,
		LastStateUpdateTime: time.Now().UTC(),
		Mutex:               &sync.Mutex{},
		Observable:          device.NewObservable(),
		UDPChannel:          channel,
		ImmoStorage:         immostorage.NewImmobilizerStorage(),
	}
	d.Observable.Attach(&observers.ConsoleTestObserver{})
	storage.DeviceCollection[serial] = d
	return d
}
