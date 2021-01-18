package unitofwork

import (
	"container/list"
	"fmt"
	"genx-go/adaptors"
	commInterfaces "genx-go/connection/interfaces"
	"genx-go/core/device"
	"genx-go/core/device/interfaces"
	"genx-go/logger"
	"genx-go/repository"
	"genx-go/repository/models"
	"reflect"
	"sync"
)

//NewDeviceUnitOfWork ...
func NewDeviceUnitOfWork(_deviceStateRepository, _activityRepository repository.IRepository) *DeviceUnitOfWork {
	return &DeviceUnitOfWork{
		activityUnitOfWork: NewDeviceActivityUnitOfWork(_activityRepository),
		stateUnitOfWork:    NewDeviceStateUnitOfWork(_deviceStateRepository),
		dirty:              make(map[string]*list.List),
		clean:              make(map[string]interfaces.IDevice),
		remove:             []string{},
		mutex:              &sync.Mutex{},
	}
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
	return uow.update()
}

//Delete ....
func (uow *DeviceUnitOfWork) Delete(identity string) {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	delete(uow.clean, identity)
	logger.Logger().WriteToLog(logger.Info, "[DeviceUnitOfWork | Delete] Device ", identity, " successfully removed.")
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
func (uow *DeviceUnitOfWork) Register(_identity string, _channel commInterfaces.IChannel) {
	activity := uow.activityUnitOfWork.Load(_identity)
	if reflect.DeepEqual(activity, &models.DeviceActivity{}) {
		uow.clean[_identity] = device.NewActivityLessDevice(_channel)
		return
	}
	adaptableActivity := adaptors.NewDeviceActivity(activity)
	sensors := adaptableActivity.Adapt()

	//_device := device.NewDevice(adaptableActivity.DTO().Parameter24, sensors, _channel)
	_device := device.NewDevice([]string{"1", "7", "13", "36", "3", "4", "23", "65", "10", "17", "11", "79", "46", "44", "43", "41", "48", "130"}, sensors, _channel)
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
		return nil
	}
	return d
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
