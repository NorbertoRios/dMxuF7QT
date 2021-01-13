package repository

import "genx-go/core/device/interfaces"

//IRepository interface for all repositories
type IRepository interface {
	Save(...interfaces.IDevice) error
	Load(string) interface{}
}
