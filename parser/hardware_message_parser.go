package parser

import (
	"genx-go/core/sensors"
	"genx-go/message"
	"regexp"
)

//BuildGenxHardwareMessageParser returns hardware message sensors
func BuildGenxHardwareMessageParser() *GenxHardwareMessageParser {
	fwExpr, _ := regexp.Compile(`FW:([^\s]+)`)
	ignExpr, _ := regexp.Compile(`Ign-([^,]+)`)
	voltExpr, _ := regexp.Compile(`Volt-([^,]+)`)
	//switchExpr, _ := regexp.Compile(`Switch-([^,]+)`)
	//relayExpr, _ := regexp.Compile(`Relay-([^,]+)`)
	return &GenxHardwareMessageParser{
		SingleSensorsBuilders: map[*regexp.Regexp]func(string) sensors.ISensor{
			fwExpr:   sensors.BuildFirmwareSensor,
			ignExpr:  sensors.BuildIgnitionSensorFromString,
			voltExpr: sensors.BuildPowerSensorFromString,
		},
	}
}

//GenxHardwareMessageParser represents sensors for hardware message
type GenxHardwareMessageParser struct {
	SingleSensorsBuilders map[*regexp.Regexp]func(string) sensors.ISensor
	ArraySensorsBuilders  map[*regexp.Regexp]func(string) []sensors.ISensor
}

//Parse parse genx hardware message
func (parser *GenxHardwareMessageParser) Parse(rawMessage *message.RawMessage) *message.Message {
	messageSensors := make([]sensors.ISensor, 0)
	for expr, builder := range parser.SingleSensorsBuilders {
		if expr.Match(rawMessage.RawData) {
			value := expr.FindAllStringSubmatch(string(rawMessage.RawData), -1)[0][1]
			sens := builder(value)
			if sens != nil {
				messageSensors = append(messageSensors, sens)
			}
		}
	}
	serial := &sensors.SerialSensor{RawValue: rawMessage.SerialNumber}
	message := &message.Message{
		Sensors:     messageSensors,
		MessageType: rawMessage.MessageType,
		Identity:    serial.ToIdentity(),
	}
	return message
}
