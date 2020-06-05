package device

import (
	"genx-go/core/sensors"
	"genx-go/message"
	"time"
)

//Device struct
type Device struct {
	Identity         string
	Sensors          []sensors.ISensor
	LastActivityTS   time.Time
	onMessageArrived func(msg *message.Message)
}

//BuildDevice device intialize
func BuildDevice(identity string) *Device {
	device := &Device{}
	device.Identity = identity
	return device
}

//CurrentState returns device current state
func (device *Device) CurrentState() []sensors.ISensor {
	return nil
}
