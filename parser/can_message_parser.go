package parser

import (
	"genx-go/core/sensors"
	"genx-go/message"
	"regexp"
	"strings"
)

//BuildCANMessageParser returns new CAN Message parser
func BuildCANMessageParser() *CANMessageParser {
	canBaseExpression, _ := regexp.Compile(`(J1939|J1708|OBD)`)
	canFuelExpression, _ := regexp.Compile(`FL\:([^\s]+)`)
	canVinNumberExpression, _ := regexp.Compile(`VIN:([^\s]+)`)
	canOBDDTCExpression, _ := regexp.Compile(`OBDDTC:(.*)`)
	return &CANMessageParser{
		CANBaseExpression:      canBaseExpression,
		CANFuelValueExpression: canFuelExpression,
		CANVinExpression:       canVinNumberExpression,
		CANOBDDTCExpression:    canOBDDTCExpression,
	}
}

//CANMessageParser parser for can messages
type CANMessageParser struct {
	CANBaseExpression      *regexp.Regexp
	CANFuelValueExpression *regexp.Regexp
	CANVinExpression       *regexp.Regexp
	CANOBDDTCExpression    *regexp.Regexp
}

//Parse parse CAN message
func (parser *CANMessageParser) Parse(rawMessage *message.RawMessage) interface{} {
	if parser.CANBaseExpression.Match(rawMessage.RawData) {
		sensorsArr := make([]sensors.ISensor, 0)
		if fl := parser.parseFuelLevel(rawMessage.RawData); fl != nil {
			sensorsArr = append(sensorsArr, fl)
		}
		if vin := parser.parseVIN(rawMessage.RawData); vin != nil {
			sensorsArr = append(sensorsArr, vin)
		}
		if dtcCodes := parser.parseDTC(rawMessage.RawData); dtcCodes != nil {
			sensorsArr = append(sensorsArr, dtcCodes)
		}
		return &message.Message{
			MessageType: rawMessage.MessageType,
			Identity:    rawMessage.Identity(),
			Sensors:     sensorsArr,
		}
	}
	return nil
}

func (parser *CANMessageParser) parseVIN(rawMessage []byte) sensors.ISensor {
	if !parser.CANVinExpression.Match(rawMessage) {
		return nil
	}
	vin := parser.CANVinExpression.FindAllStringSubmatch(string(rawMessage), -1)[0][1]
	return &sensors.VINSensor{VIN: vin}
}

func (parser *CANMessageParser) parseFuelLevel(rawMessage []byte) sensors.ISensor {
	if !parser.CANFuelValueExpression.Match(rawMessage) {
		return nil
	}
	subm := parser.CANFuelValueExpression.FindAllStringSubmatch(string(rawMessage), -1)
	return sensors.BuildFuelSensorFromString(subm[0][1])
}

func (parser *CANMessageParser) parseDTC(rawMessage []byte) sensors.ISensor {
	if !parser.CANOBDDTCExpression.Match(rawMessage) {
		return nil
	}
	subm := strings.Replace(parser.CANOBDDTCExpression.FindAllStringSubmatch(string(rawMessage), -1)[0][1], "OBDPDTC:", "", 1)
	return sensors.BuildDTCCodesSensorFromString(subm)
}
