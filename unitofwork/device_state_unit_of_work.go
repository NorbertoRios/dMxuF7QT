package unitofwork

import (
	"genx-go/repository"
)

//DeviceStateUnitOfWork ...
type DeviceStateUnitOfWork struct {
	BaseUnitOfWork
	repository repository.IRepository
}
