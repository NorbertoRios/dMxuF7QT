package models

import (
	"genx-go/core/device/interfaces"
	"time"
)

//NewMessageHistory ...
func NewMessageHistory(_device interfaces.IDevice) *MessageHistory {
	return &MessageHistory{}
}

//MessageHistory struct
type MessageHistory struct {
	ID              uint64    `gorm:"column:ID;primary_key"`
	DevID           string    `gorm:"column:DevId"`
	EntryData       []byte    `gorm:"column:EntryData"`
	ParsedEntryData []byte    `gorm:"column:ParsedEntryData"`
	Time            time.Time `gorm:"column:Time"`
	RecievedTime    time.Time `gorm:"column:RecievedTime"`
	ReportClass     string    `gorm:"column:ReportClass"`
	ReportType      int32     `gorm:"column:ReportType"`
	Reason          int32     `gorm:"column:Reason"`
	Latitude        float32   `gorm:"column:Latitude"`
	Longitude       float32   `gorm:"column:Longitude"`
	Speed           float32   `gorm:"column:Speed"`
	ValidFix        byte      `gorm:"column:ValidFix"`
	Altitude        float32   `gorm:"column:Altitude"`
	Heading         float32   `gorm:"column:Heading"`
	IgnitionState   byte      `gorm:"column:IgnitionState"`
	Odometer        int32     `gorm:"column:Odometer"`
	Satellites      byte      `gorm:"column:Satellites"`
	Supply          int32     `gorm:"column:Supply"`
	GPIO            byte      `gorm:"column:GPIO"`
	Relay           byte      `gorm:"column:Relay"`
}
