package unitofwork

import (
	"container/list"
	"genx-go/repository"
	"sync"
)

//NewDeviceStateUnitOfWork ...
func NewDeviceStateUnitOfWork(_repository repository.IRepository) *DeviceStateUnitOfWork {
	uow := &DeviceStateUnitOfWork{}
	uow.repository = _repository
	uow.mutex = &sync.Mutex{}
	uow.dirty = make(map[string]*list.List)
	return uow
}

//DeviceStateUnitOfWork ...
type DeviceStateUnitOfWork struct {
	BaseUnitOfWork
}
