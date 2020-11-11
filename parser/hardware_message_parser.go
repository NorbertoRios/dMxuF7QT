package parser

import (
	"fmt"
	"genx-go/core/sensors"
	"genx-go/logger"
	"genx-go/message"
	"regexp"
)

//BuildGenxHardwareMessageParser returns hardware message sensors
func BuildGenxHardwareMessageParser() *GenxHardwareMessageParser {
	fwExpr, _ := regexp.Compile(`FW:([^\s]+)`)
	ignExpr, _ := regexp.Compile(`Ign-([^,]+)`)
	voltExpr, _ := regexp.Compile(`Volt-([^,]+)`)
	switchExpr, _ := regexp.Compile(`Switch-([^,]+)`)
	relayExpr, _ := regexp.Compile(`Relay-([^,]+)`)
	return &GenxHardwareMessageParser{
		FirmwareExpression: fwExpr,
		SensorBuilders: map[*regexp.Regexp]interface{}{
			ignExpr:    sensors.BuildIgnitionSensorFromString,
			voltExpr:   sensors.BuildPowerSensorFromString,
			switchExpr: sensors.BuildInputsFromString,
			relayExpr:  sensors.BuildOutputsFromString,
		},
	}
}

//GenxHardwareMessageParser represents sensors for hardware message
type GenxHardwareMessageParser struct {
	FirmwareExpression *regexp.Regexp
	SensorBuilders     map[*regexp.Regexp]interface{}
}

//Parse parse genx hardware message
func (parser *GenxHardwareMessageParser) Parse(rawMessage *message.RawMessage) *message.HardwareMessage {
	messageSensors := make([]sensors.ISensor, 0)
	for expr, builder := range parser.SensorBuilders {
		if expr.Match(rawMessage.RawData) {
			value := expr.FindAllStringSubmatch(string(rawMessage.RawData), -1)[0][1]
			switch builder.(type) {
			case func(string) sensors.ISensor:
				{
					messageSensors = append(messageSensors, builder.(func(string) sensors.ISensor)(value))
				}
			case func(string) []sensors.ISensor:
				{
					messageSensors = append(messageSensors, builder.(func(string) []sensors.ISensor)(value)...)
				}
			}
		}
	}
	var fwVersion string
	if !parser.FirmwareExpression.Match(rawMessage.RawData) {
		logger.Logger().WriteToLog(logger.Error, fmt.Sprintf("[GenxHardwareMessageParser] Can't extract firmwar version from message: %v", string(rawMessage.RawData)))
		fwVersion = ""
	}
	fwVersion = parser.FirmwareExpression.FindAllStringSubmatch(string(rawMessage.RawData), -1)[0][1]
	message := &message.HardwareMessage{
		Firmware:    fwVersion,
		Sensors:     messageSensors,
		MessageType: rawMessage.MessageType,
		Identity:    rawMessage.Identity(),
	}
	return message
}
