package unitofwork

import (
	"fmt"
	"genx-go/core/device/interfaces"
	"genx-go/logger"
	"genx-go/repository"
	"reflect"
)

//NewDeviceUnitOfWork ...
func NewDeviceUnitOfWork(_deviceStateRepository, _activityRepository repository.IRepository) *DeviceUnitOfWork {
	uow := &DeviceUnitOfWork{
		deviceStateRepository: _deviceStateRepository,
		activityRepository:    _activityRepository,
	}
	uow.clean = make(map[string]interfaces.IDevice)
	uow.dirty = make(map[string]interfaces.IDevice)
	uow.register = make(map[string]struct{})
	uow.remove = make(map[string]interfaces.IDevice)
	return uow
}

//DeviceUnitOfWork ...
type DeviceUnitOfWork struct {
	UnitOfWork
	deviceStateRepository repository.IRepository
	activityRepository    repository.IRepository
}

//Commit ...
func (uow *DeviceUnitOfWork) Commit() {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	uow.registerDevices()
	uow.updateDevices()
	uow.removeDevices()
}

func (uow *DeviceUnitOfWork) updateDevices() bool {
	devices := []interfaces.IDevice{}
	for identity, device := range uow.dirty {
		if reflect.DeepEqual(uow.deviceStateBuffer[identity], device.LastDeviceMessage()) {
			devices = append(devices, device)
		}
	}
	return uow.deviceStateRepository.Save() == nil && uow.activityRepository.Save() == nil
}

func (uow *DeviceUnitOfWork) removeDevices() {
	for identity := range uow.remove {
		logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[DeviceUnitOfWork | registerDevices] Device %v removed", identity))
		delete(uow.clean, identity)
	}
}

func (uow *DeviceUnitOfWork) registerDevices() {
	for identity := range uow.register {
		_device := uow.createDevice(identity)
		logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[DeviceUnitOfWork | registerDevices] Device %v registered", identity))
		uow.clean[identity] = _device
		delete(uow.register, identity)
	}
}

func (uow *DeviceUnitOfWork) createDevice(identity string) interfaces.IDevice {
	uow.activityRepository.Load(identity)
	uow.deviceStateRepository.Load(identity)
	return nil
}
