package genxmessage

import (
	"genx-go/core"
	"genx-go/core/sensors"
)

//BuildGenxMessage build new GenxMessage
func BuildGenxMessage(rData map[string]interface{}, messageType string) *GenxMessage {
	sensorsBuilder := BuildBinaryMessageSensors()
	serial := &sensors.SerialSensor{RawValue: rData[core.UnitName]}
	message := &GenxMessage{
		Sensors:     sensorsBuilder.Build(rData),
		Identity:    serial.ToIdentity(),
		MessageType: messageType,
	}
	return message
}
