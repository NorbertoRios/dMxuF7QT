package peripherystorage

import (
	"genx-go/core/device/interfaces"
	"genx-go/core/immobilizer"
	"sync"
)

//NewImmobilizerStorage ...
func NewImmobilizerStorage() *ImmobilizerStorage {
	return &ImmobilizerStorage{
		mutex:        &sync.Mutex{},
		immobilizers: make(map[int]interfaces.IImmobilizer, 0),
	}
}

//ImmobilizerStorage storage
type ImmobilizerStorage struct {
	mutex        *sync.Mutex
	immobilizers map[int]interfaces.IImmobilizer
}

//Immobilizer ...
func (storage *ImmobilizerStorage) Immobilizer(index int, trigger string, device interfaces.IDevice) interfaces.IImmobilizer {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	var immo interfaces.IImmobilizer
	var f bool
	if immo, f = storage.immobilizers[index]; f && immo.Trigger() == trigger {
		return immo
	}
	immo = immobilizer.NewImmobilizer(index, trigger)
	storage.immobilizers[index] = immo
	return immo
}
