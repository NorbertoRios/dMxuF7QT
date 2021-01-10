package unitofwork

import "genx-go/core/device/interfaces"

//IDeviceUnitOfWork ...
type IDeviceUnitOfWork interface {
	Commit() error
	Device(string) interfaces.IDevice
	Delete(string)
}
