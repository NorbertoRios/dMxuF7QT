package unitofwork

import (
	"container/list"
	"genx-go/core/device"
	"genx-go/core/device/interfaces"
	"genx-go/logger"
	"genx-go/repository"
	"sync"
)

//BaseUnitOfWork ...
type BaseUnitOfWork struct {
	dirty      map[string]*list.List
	mutex      *sync.Mutex
	repository repository.IRepository
}

//Update ...
func (uow *BaseUnitOfWork) Update(identity string, _device device.Device) {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	var l *list.List
	l, f := uow.dirty[identity]
	if !f {
		l = list.New()
	}
	l.PushBack(_device)
	uow.dirty[identity] = l
}

func (uow *BaseUnitOfWork) update() bool {
	devices := []interfaces.IDevice{}
	for _, list := range uow.dirty {
		for element := list.Front(); element != nil; element = element.Next() {
			d, _ := element.Value.(interfaces.IDevice)
			devices = append(devices, d)
		}
	}
	err := uow.repository.Save(devices...)
	if err == nil {
		uow.dirty = make(map[string]*list.List)
		return true
	}
	logger.Logger().WriteToLog(logger.Info, "[DeviceStateUnitOfWork | Commit] Error while commit. Error: ", err.Error())
	return false
}

//Commit ...
func (uow *BaseUnitOfWork) Commit() bool {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	return uow.update()
}
