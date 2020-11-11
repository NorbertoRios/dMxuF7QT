package models

import "time"

//ConfigurationModel model of device configuration
type ConfigurationModel struct {
	ID        int32     `gorm:"column:cfgId;primary_key"`
	Identity  string    `gorm:"column:devIdentity"`
	Command   string    `gorm:"column:cfgCommand"`
	CreatedAt time.Time `gorm:"column:cfgCreated_at"`
}