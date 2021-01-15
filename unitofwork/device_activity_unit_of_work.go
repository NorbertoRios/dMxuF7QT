package unitofwork

import (
	"container/list"
	"genx-go/repository"
	"genx-go/repository/models"
	"sync"
)

//NewDeviceActivityUnitOfWork ...
func NewDeviceActivityUnitOfWork(_repository repository.IRepository) *DeviceActivityUnitOfWork {
	uow := &DeviceActivityUnitOfWork{}
	uow.repository = _repository
	uow.mutex = &sync.Mutex{}
	uow.dirty = make(map[string]*list.List)
	return uow
}

//DeviceActivityUnitOfWork ...
type DeviceActivityUnitOfWork struct {
	BaseUnitOfWork
}

//Load ...
func (uow *DeviceActivityUnitOfWork) Load(_identity string) *models.DeviceActivity {
	return uow.repository.Load(_identity).(*models.DeviceActivity)
}
