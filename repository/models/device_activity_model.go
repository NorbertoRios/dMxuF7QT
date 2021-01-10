package models

import (
	"genx-go/core/device/interfaces"
	"time"
)

//NewDeviceActivity ...
func NewDeviceActivity(_device interfaces.IDevice) *DeviceActivity {
	return &DeviceActivity{
		// Identity:           _device.Identity(),
		// MessageTime:        _device.LastLocationMessage().Time(),
		// LastUpdateTime:     _device.LastUpdateTime(),
		// LastMessageID:      _device.LastMessageID(),
		// LastMessage:        _device.LastLocationMessage().DTO(),
		// Serializedsoftware: _device.Software(),
	}
}

//DeviceActivity device activity model
type DeviceActivity struct {
	Identity           string    `gorm:"column:daiDeviceIdentity"`
	MessageTime        time.Time `gorm:"column:daiLastMessageTime"`
	LastUpdateTime     time.Time `gorm:"column:daiLastUpdateTime"`
	LastMessageID      uint64    `gorm:"column:daiLastMessageId"`
	LastMessage        string    `gorm:"column:daiLastMessage"`
	Serializedsoftware string    `gorm:"column:daiSoftware"`
}
