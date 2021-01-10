package unitofwork

import (
	"fmt"
	"genx-go/core/device/interfaces"
	"genx-go/logger"
	"sync"
)

//UnitOfWork ...
type UnitOfWork struct {
	register          map[string]struct{}
	clean             map[string]interfaces.IDevice
	dirty             map[string]interfaces.IDevice
	remove            map[string]interfaces.IDevice
	deviceStateBuffer map[string]interfaces.IDevice
	mutex             *sync.Mutex
}

//Device ...
func (uow *UnitOfWork) Device(identity string) interfaces.IDevice {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	var device interfaces.IDevice
	var f bool
	if device, f = uow.clean[identity]; !f {
		logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[UnitOfWork | Device] Device with identity %v not found.", identity))
	}
	return device
}

//Update ...
func (uow *UnitOfWork) Update(identity string, _device interfaces.IDevice) {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[UnitOfWork | Update] Device: %v has become in the queue for updating", identity))
	uow.deviceStateBuffer[identity] = _device
	uow.dirty[identity] = _device
}

//Register ...
func (uow *UnitOfWork) Register(identity string) {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	uow.register[identity] = struct{}{}
	logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[UnitOfWork | Register] Device: %v has become in the queue for registartion", identity))
}

//Remove ...
func (uow *UnitOfWork) Remove(identity string, _device interfaces.IDevice) {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	uow.remove[identity] = _device
	logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[UnitOfWork | Remove] Device: %v has become in the queue for removing", identity))
}
