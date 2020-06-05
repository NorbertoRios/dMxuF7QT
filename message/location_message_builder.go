package message

import (
	"genx-go/core"
	"genx-go/core/sensors"
)

//BuildMessage build new message
func BuildMessage(rData map[string]interface{}, messageType string) *Message {
	sensorsBuilder := BuildBinaryMessageSensors()
	serial := &sensors.SerialSensor{RawValue: rData[core.UnitName]}
	message := &Message{
		Sensors:     sensorsBuilder.Build(rData),
		Identity:    serial.ToIdentity(),
		MessageType: messageType,
	}
	return message
}
