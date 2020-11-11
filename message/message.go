package message

import (
	"genx-go/core/sensors"
)

//Message genx parsed message
type Message struct {
	MessageType string
	Identity    string
	Sensors     []sensors.ISensor
	SID         uint64
}
