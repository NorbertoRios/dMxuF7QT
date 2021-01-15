package unitofwork

import (
	connInterfaces "genx-go/connection/interfaces"
	"genx-go/core/device/interfaces"
)

//IDeviceUnitOfWork ...
type IDeviceUnitOfWork interface {
	Commit() bool
	Device(string) interfaces.IDevice
	Delete(string)
	Register(string, connInterfaces.IChannel)
}
