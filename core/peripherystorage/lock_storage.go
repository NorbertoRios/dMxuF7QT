package peripherystorage

import (
	"genx-go/core/device/interfaces"
	"genx-go/core/lock"
	"sync"
)

//NewElectricLockStorage ...
func NewElectricLockStorage() *ElectricLockStorage {
	return &ElectricLockStorage{
		mutex: &sync.Mutex{},
		locks: make(map[int]interfaces.IProcess, 0),
	}
}

//ElectricLockStorage storage
type ElectricLockStorage struct {
	mutex *sync.Mutex
	locks map[int]interfaces.IProcess
}

//ElectricLock ...
func (storage *ElectricLockStorage) ElectricLock(index int, device interfaces.IDevice) interfaces.IProcess {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	var eLock interfaces.IProcess
	var f bool
	if eLock, f = storage.locks[index]; f {
		return eLock
	}
	eLock = lock.NewElectricLock(device, index)
	storage.locks[index] = eLock
	return eLock
}
