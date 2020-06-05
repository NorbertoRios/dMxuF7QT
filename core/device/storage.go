package device

import (
	"log"
	"sync"
	"time"
)

//Storage godfather of devices
type Storage struct {
	Devices           map[string]*Device
	Mutex             *sync.Mutex
	DeviceUpdateState func(device *Device)
	DeviceCreated     func(device *Device)
}

//InitStorage return new storage for devices
func InitStorage(onDeviceUpdateState func(device *Device), onDeviceCreated func(device *Device)) *Storage {
	storage := &Storage{
		Devices:           make(map[string]*Device),
		Mutex:             &sync.Mutex{},
		DeviceUpdateState: onDeviceUpdateState,
	}
	defer storage.start()
	return storage
}

func (storage *Storage) removeDevice(identity string) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	delete(storage.Devices, identity)
	log.Println("[Storage] Device ", identity, " removed. Total device count:", len(storage.Devices))
}

func (storage *Storage) start() {
	ticker := time.NewTicker(300 * time.Second)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("[Storage] Recovered in watchdog function:", r)
			}
		}()
		for {
			select {
			case <-ticker.C:
				for identity, device := range storage.Devices {
					if time.Now().UTC().Sub(device.LastActivityTS).Seconds() >= 3600 {
						storage.removeDevice(identity)
					}
				}
			}
		}
	}()
}
