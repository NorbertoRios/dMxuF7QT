package unitofwork

import (
	"container/list"
	"fmt"
	"genx-go/core/device"
	"genx-go/core/device/interfaces"
	"genx-go/logger"
	"genx-go/repository"
	"sync"
)

//NewDeviceUnitOfWork ...
func NewDeviceUnitOfWork(_deviceStateRepository, _activityRepository repository.IRepository) *DeviceUnitOfWork {
	return &DeviceUnitOfWork{}
}

//DeviceUnitOfWork ...
type DeviceUnitOfWork struct {
	activityUnitOfWork *DeviceActivityUnitOfWork
	stateUnitOfWork    *DeviceStateUnitOfWork
	dirty              map[string]*list.List
	clean              map[string]interfaces.IDevice
	remove             []string
	mutex              *sync.Mutex
}

//Commit ...
func (uow *DeviceUnitOfWork) Commit() bool {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	return uow.removeDevices() && uow.update()
}

func (uow *DeviceUnitOfWork) removeDevices() bool {
	for _, identity := range uow.remove {
		delete(uow.clean, identity)
		logger.Logger().WriteToLog(logger.Info, "[DeviceUnitOfWork | removeDevices] Device :", identity, " removed.")
	}
	return true
}

func (uow *DeviceUnitOfWork) update() bool {
	if !uow.activityUnitOfWork.Commit() || !uow.activityUnitOfWork.Commit() {
		logger.Logger().WriteToLog(logger.Error, "[DeviceUnitOfWork | update] Cant save changes to database")
		return false
	}
	for identity, dList := range uow.dirty {
		uow.clean[identity] = dList.Back().Value.(interfaces.IDevice)
		logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[DeviceUnitOfWork | update] Device : %v added to clean", identity))
	}
	uow.dirty = make(map[string]*list.List)
	return true
}

func (uow *DeviceUnitOfWork) addToDirty(_identity string, _device interfaces.IDevice) {
	l, f := uow.dirty[_identity]
	if !f {
		logger.Logger().WriteToLog(logger.Info, "[DeviceUnitOfWork | update] For device: ", _identity, " collection not found. Creating new collection...")
		l = list.New()
		uow.mutex.Lock()
		uow.dirty[_identity] = l
		uow.mutex.Unlock()
	}
	l.PushBack(*_device.(*device.Device))
}

//UpdateActivity ...
func (uow *DeviceUnitOfWork) UpdateActivity(_identity string, _device interfaces.IDevice) {
	uow.addToDirty(_identity, _device)
	uow.activityUnitOfWork.Update(_identity, *_device.(*device.Device))
}

//UpdateState ...
func (uow *DeviceUnitOfWork) UpdateState(_identity string, _device interfaces.IDevice) {
	uow.addToDirty(_identity, _device)
	uow.stateUnitOfWork.Update(_identity, *_device.(*device.Device))
}

//Register ...
func (uow *DeviceUnitOfWork) Register(_identity string) {
	//activity := uow.activityUnitOfWork.Load(_identity)
	//lastKnownState := uow.stateUnitOfWork.Load(_identity)
	_device := &device.Device{} //create device based on activity and last known state
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	uow.clean[_identity] = _device

}

//Device ...
func (uow *DeviceUnitOfWork) Device(_identity string) interfaces.IDevice {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	d, f := uow.clean[_identity]
	if !f {
		logger.Logger().WriteToLog(logger.Info, "[DeviceUnitOfWork | Device] Device with identity: ", _identity, " does not exist")
		return &device.Device{}
	}
	return d
}
