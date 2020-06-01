package genxparser

import (
	"genx-go/core/sensors"
	"genx-go/genxmessage"
	"genx-go/genxutils"
	"regexp"
)

// 1WIRE:10/NoDevice
// 1:0000000000000000
// 2:0000000000000000
// 3:0000000000000000
// 4:0000000000000000
// 5:0000000000000000
// 000003912835 3912835

// 1:10B8993E01080099 TS-65.300003
// 2:109F6108000000AE TS-60.799999
// 3:288D19BD01000096 TS-65.862503
// 4:0000000000000000
// 5:01AB16430F00004E ID

// 1:10B8993E01080099 TS-65.300003
// 2:102E5A080000006A TS-56.299999
// 3:109F6108000000AE TS-60.799999
// 4:288D19BD01000096 TS-66.199997
// 5:01AB16430F00004E ID

//BuildOneWireMessageParser returns OneWireMessage parser
func BuildOneWireMessageParser() *OneWireMessageParser {
	oneWireExpr, _ := regexp.Compile(`1WIRE:\w+\/(.*)`)
	tempSensorsExpr, _ := regexp.Compile(`(\d):(.*) TS-(.*)`)
	iButtonSesnorExpr, _ := regexp.Compile(`\d:\w+(.{8}) ID`)
	return &OneWireMessageParser{
		OneWireExpression:            oneWireExpr,
		TemperatureSensorsExpression: tempSensorsExpr,
		IButtonSesnorExpression:      iButtonSesnorExpr,
	}

}

//OneWireMessageParser returns message parser
type OneWireMessageParser struct {
	OneWireExpression            *regexp.Regexp
	TemperatureSensorsExpression *regexp.Regexp
	IButtonSesnorExpression      *regexp.Regexp
}

//Parse parse 1WireMessage
func (parser *OneWireMessageParser) Parse(rawMessage *genxmessage.RawMessage) *genxmessage.GenxMessage {
	if !parser.OneWireExpression.Match(rawMessage.RawData) {
		serial := sensors.SerialSensor{RawValue: rawMessage.SerialNumber}
		return &genxmessage.GenxMessage{
			Identity:    serial.ToIdentity(),
			MessageType: rawMessage.MessageType,
			Sensors:     nil,
		}
	}
	wireState := parser.OneWireExpression.FindAllStringSubmatch(string(rawMessage.RawData), -1)[0][1]
	if wireState != "Present" {
		serial := sensors.SerialSensor{RawValue: rawMessage.SerialNumber}
		return &genxmessage.GenxMessage{
			Identity:    serial.ToIdentity(),
			MessageType: rawMessage.MessageType,
			Sensors:     nil,
		}
	}
	finalSensors := make([]sensors.ISensor, 0)
	if parser.TemperatureSensorsExpression.Match(rawMessage.RawData) {
		tSensors := parser.parseTemperatureSensors(rawMessage.RawData)
		finalSensors = append(finalSensors, tSensors...)
	}
	if parser.IButtonSesnorExpression.Match(rawMessage.RawData) {
		iButtonSensors := parser.parseIButtonSensors(rawMessage.RawData)
		finalSensors = append(finalSensors, iButtonSensors)
	}
	serial := &sensors.SerialSensor{RawValue: rawMessage.SerialNumber}
	return &genxmessage.GenxMessage{
		Identity:    serial.ToIdentity(),
		Sensors:     finalSensors,
		MessageType: rawMessage.MessageType,
	}
}

func (parser *OneWireMessageParser) parseTemperatureSensors(rawData []byte) []sensors.ISensor {
	messageSensors := make([]sensors.ISensor, 0)
	subMatch := parser.TemperatureSensorsExpression.FindAllStringSubmatch(string(rawData), -1)
	for i := range subMatch {
		sID := &genxutils.StringUtils{Data: subMatch[i][1]}
		sValue := &genxutils.StringUtils{Data: subMatch[i][3]}
		sensor := &sensors.TemperatureSensor{
			ID:    sID.Byte(),
			Imei:  subMatch[i][2],
			Value: sValue.Float32(),
		}
		messageSensors = append(messageSensors, sensor)
	}
	return messageSensors
}

func (parser *OneWireMessageParser) parseIButtonSensors(rawData []byte) sensors.ISensor {
	if !parser.IButtonSesnorExpression.Match(rawData) {
		return nil
	}
	subMatch := parser.IButtonSesnorExpression.FindAllStringSubmatch(string(rawData), -1)[0][1]
	sValue := &genxutils.StringUtils{Data: subMatch}
	return &sensors.IButtonSensor{BtnID: sValue.Int32(16)}
}
