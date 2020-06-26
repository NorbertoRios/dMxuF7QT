package message

import "genx-go/core/sensors"

//HardwareMessage represents hardware message
type HardwareMessage struct {
	Identity    string
	MessageType string
	Firmware    string
	Sensors     []sensors.ISensor
}
