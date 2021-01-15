package repository

import (
	"genx-go/core/device/interfaces"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//NewDeviceActivityRepository ...
func NewDeviceActivityRepository(_connectionString string) *DeviceActivityRepository {
	_connection, err := gorm.Open(mysql.Open(_connectionString), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Error connecting to raw database:" + err.Error())
	}
	return &DeviceActivityRepository{
		Connection: _connection,
	}
}

//DeviceActivityRepository ...
type DeviceActivityRepository struct {
	Connection *gorm.DB
}

//Save ...
func (r *DeviceActivityRepository) Save(device ...interfaces.IDevice) error {
	return nil
}

//Load ...
func (r *DeviceActivityRepository) Load(identity string) interface{} {
	return nil
}
