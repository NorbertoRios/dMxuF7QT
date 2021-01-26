package repository

import (
	"errors"
	"genx-go/core/device"
	genxLogger "genx-go/logger"
	"genx-go/repository/models"

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
func (r *DeviceActivityRepository) Save(device ...device.Device) error {
	return nil
}

//Load ...
func (r *DeviceActivityRepository) Load(identity string) interface{} {
	d := &models.DeviceActivity{}
	err := r.Connection.Where("daiDeviceIdentity=?", identity).Find(d).Error
	bErr := errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil && bErr {
		genxLogger.Logger().WriteToLog(genxLogger.Info, "[DeviceActivityRepository | Load] Activity for device ", identity, " not found")
	} else if err != nil && !bErr {
		genxLogger.Logger().WriteToLog(genxLogger.Fatal, "[DeviceActivityRepository | Load] Error while loading activity. Device ", identity, ". Error: ", err)
	}
	return d
}
