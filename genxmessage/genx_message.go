package genxmessage

import (
	"genx-go/core/sensors"
)

//GenxMessage genx parsed message
type GenxMessage struct {
	MessageType string
	Identity    string
	Sensors     []sensors.ISensor
}
