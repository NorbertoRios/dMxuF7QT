package models

import (
	"genx-go/core/device/interfaces"
	"time"
)

//NewDeviceActivity ...
func NewDeviceActivity(_device interfaces.IDevice) *DeviceActivity {
	state := _device.CurrentDeviceState()
	for _, deviceSensor := range state {
		
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
