package device

import (
	"genx-go/connection"
	"genx-go/core/sensors"
)

//BuildDevice build new device
func BuildDevice(channel *connection.UDPChannel, activity []sensors.ISensor, currentConfig map[string]string) {

}

//Device struct
type Device struct {
	BaseDevice
}
