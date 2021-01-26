package repository

import "genx-go/core/device"

//IRepository interface for all repositories
type IRepository interface {
	Save(...device.Device) error
	Load(string) interface{}
}
