package device

import (
	"genx-go/core/sensors"
	"genx-go/repository/models"
	"log"
	"sync"
)

const (
	CurrentConfig = "current"
	Unsended      = "unsended"
)

//Storage godfather of devices
type Storage struct {
	UnknownDevices    map[string]*BaseDevice
	Devices           map[string]*Device
	Mutex             *sync.Mutex
	DeviceUpdateState func(device *Device)
	LoadDeviceConfig  func(string, string) *models.ConfigurationModel
	LoadDeviceState   func(identity string) []sensors.ISensor
}

//ConstructStorage return new storage for devices
func ConstructStorage(onDeviceUpdateState func(device *Device), onNeedLoadDeviceConfig func(string, string) *models.ConfigurationModel, onNeedLoadLastDeviceState func(identity string) []sensors.ISensor) *Storage {
	storage := &Storage{
		Devices:           make(map[string]*Device),
		UnknownDevices:    make(map[string]*BaseDevice),
		Mutex:             &sync.Mutex{},
		DeviceUpdateState: onDeviceUpdateState,
		LoadDeviceConfig:  onNeedLoadDeviceConfig,
		LoadDeviceState:   onNeedLoadLastDeviceState,
	}
	//defer storage.start()
	return storage
}

//Device returns device by identity
func (storage *Storage) Device(identity string) *Device {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	if device, f := storage.Devices[identity]; f {
		return device
	}
	return nil
}

func (storage *Storage) createDevice(identity string) *Device {
	_ = storage.LoadDeviceState(identity)
	_ = storage.LoadDeviceConfig(identity, CurrentConfig)
	return nil
}

//SaveDevice save device to device storage
func (storage *Storage) SaveDevice(device *Device) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	storage.Devices[device.Identity] = device
	log.Println("[Storage] Device ", device.Identity, " has been added. Total device count:", len(storage.Devices))
}

func (storage *Storage) removeDevice(identity string) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	delete(storage.Devices, identity)
	log.Println("[Storage] Device ", identity, " removed. Total device count:", len(storage.Devices))
}

// func (storage *Storage) start() {
// 	ticker := time.NewTicker(300 * time.Second)
// 	go func() {
// 		defer func() {
// 			if r := recover(); r != nil {
// 				log.Println("[Storage] Recovered in watchdog function:", r)
// 			}
// 		}()
// 		for {
// 			select {
// 			case <-ticker.C:
// 				for identity, device := range storage.Devices {
// 					if time.Now().UTC().Sub(device.LastActivityTS).Seconds() >= 3600 {
// 						storage.removeDevice(identity)
// 					}
// 				}
// 			}
// 		}
// 	}()
// }
