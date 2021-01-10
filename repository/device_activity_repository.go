package repository

import (
	"genx-go/core/device/interfaces"

	"gorm.io/gorm"
)

//NewDeviceActivityRepository ...
func NewDeviceActivityRepository(_connection *gorm.DB) *DeviceActivityRepository {
	return &DeviceActivityRepository{
		Connection: _connection,
	}
}

//DeviceActivityRepository ...
type DeviceActivityRepository struct {
	Connection *gorm.DB
}

//Save ...
func (r *DeviceActivityRepository) Save(device interfaces.IDevice) error {
	return nil
}

//Load ...
func (r *DeviceActivityRepository) Load(identity string) interface{} {
	return nil
}
