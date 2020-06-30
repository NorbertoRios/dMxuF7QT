package device

import (
	"genx-go/connection"
	"genx-go/logger"
	"genx-go/message"
	"genx-go/repository/models"
	"sync"
	"time"
)

const (
	CurrentConfig = "current"
	Unsended      = "unsended"
)

//Storage godfather of devices
type Storage struct {
	PublishMessage    func(interface{})
	RawMessageFactory *message.RawMessageFactory
	Devices           map[string]IDevice
	Mutex             *sync.Mutex
	DeviceUpdateState func(IDevice)
	LoadDeviceConfig  func(string, string) *models.ConfigurationModel
	LoadDeviceState   func(identity string) *LastKnownDeviceState
}

//ConstructStorage return new storage for devices
func ConstructStorage(onDeviceUpdateState func(IDevice), onNeedLoadDeviceConfig func(string, string) *models.ConfigurationModel,
	onNeedLoadLastDeviceState func(identity string) *LastKnownDeviceState, publishMessage func(interface{})) *Storage {

	storage := &Storage{
		Devices:           make(map[string]IDevice),
		Mutex:             &sync.Mutex{},
		DeviceUpdateState: onDeviceUpdateState,
		LoadDeviceConfig:  onNeedLoadDeviceConfig,
		LoadDeviceState:   onNeedLoadLastDeviceState,
		PublishMessage:    publishMessage,
	}
	defer storage.start()
	return storage

}

func (storage *Storage) start() {
	ticker := time.NewTicker(300 * time.Second)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("[Storage] Recovered in watchdog function:", r)
			}
		}()
		for {
			select {
			case <-ticker.C:
				for identity := range storage.Devices {
					if time.Now().UTC().Sub(storage.Devices[identity].LastActivityTimeStamp()).Seconds() >= 3600 {
						storage.removeDevice(identity)
					}
				}
			}
		}
	}()
}

//NewMessage on new message from device
func (storage *Storage) NewMessage(channel *connection.UDPChannel, packet []byte) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("[NewMessage] panic:Recovered in new message:", r)
		}
	}()
	rawMessage := storage.RawMessageFactory.BuildRawMessage(packet)
	var device IDevice
	if storage.DeviceExist(rawMessage.Identity) {
		device = storage.Device(rawMessage.Identity)
	} else {
		lastKnownState := storage.LoadDeviceState(rawMessage.Identity)
		device = BuildBaseDevice(rawMessage.Identity, channel, lastKnownState, storage.LoadDeviceConfig,
			storage.createDevice, storage.PublishMessage)
	}
	go device.MessageArrived(rawMessage)
}

//DeviceExist is device exist (facade)
func (storage *Storage) DeviceExist(identity string) bool {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	_, found := storage.Devices[identity]
	return found
}

//Device returns device by identity
func (storage *Storage) Device(identity string) IDevice {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	return storage.Devices[identity]
}

func (storage *Storage) createDevice(device IDevice) {
	baseDevice := device.(*BaseDevice)
	newDevice := BuildDevice(baseDevice, storage.DeviceUpdateState)
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	storage.Devices[newDevice.identity] = newDevice
}

//SaveDevice save device to device storage
func (storage *Storage) SaveDevice(device *Device) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	storage.Devices[device.Identity()] = device
	logger.Info("[Storage | SaveDevice] Device ", device.identity, " has been added. Total device count:", len(storage.Devices))
}

func (storage *Storage) removeDevice(identity string) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	delete(storage.Devices, identity)
	logger.Info("[Storage | RemoveDevice] Device ", identity, " removed. Total device count:", len(storage.Devices))
}
