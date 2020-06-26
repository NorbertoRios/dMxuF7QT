package models

import "time"

//DeviceActivity device activity model
type DeviceActivity struct {
	Identity           string    `gorm:"column:daiDeviceIdentity"`
	MessageTime        time.Time `gorm:"column:daiLastMessageTime"`
	LastUpdateTime     time.Time `gorm:"column:daiLastUpdateTime"`
	LastMessageID      uint64    `gorm:"column:daiLastMessageId"`
	LastMessage        string    `gorm:"column:daiLastMessage"`
	Serializedsoftware string    `gorm:"column:daiSoftware"`
}
