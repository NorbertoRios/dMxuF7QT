package unitofwork

import (
	"genx-go/repository"
	"genx-go/repository/models"
)

//DeviceActivityUnitOfWork ...
type DeviceActivityUnitOfWork struct {
	BaseUnitOfWork
	repository repository.IRepository
}

//Load ...
func (uow *DeviceActivityUnitOfWork) Load(_identity string) *models.DeviceActivity {
	return uow.repository.Load(_identity).(*models.DeviceActivity)
}
