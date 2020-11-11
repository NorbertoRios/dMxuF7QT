package message

import (
	"fmt"
	"genx-go/configuration"
	"genx-go/core"
	"genx-go/logger"
)

//RawMessage raw genx message
type RawMessage struct {
	SerialNumber string
	MessageType  string
	RawData      []byte
}

//Identity ...
func (rm *RawMessage) Identity() string {
	serial := rm.SerialNumber
	for l := len(rm.SerialNumber); l < 12; l++ {
		serial = fmt.Sprintf("0%v", serial)
	}
	return fmt.Sprintf("genx_%v", serial)
}

//Serial ...
func (rm *RawMessage) Serial() string {
	serial := rm.SerialNumber
	for l := len(rm.SerialNumber); l < 12; l++ {
		serial = fmt.Sprintf("0%v", serial)
	}
	return serial
}

//String return slice for string values
func (message *RawMessage) String(startIndex int) (interface{}, int) {
	for i, v := range message.RawData {
		if v == 0 {
			return string(message.RawData[startIndex : startIndex+i]), i + 1
		}
	}
	return message.RawData[startIndex:], len(message.RawData[startIndex:])
}

//Value returns value as intreface{}
func (message *RawMessage) Value(startIndex int, field *configuration.Field) (interface{}, int) {
	if len(message.RawData) < startIndex+field.Size {
		logger.Logger().WriteToLog(logger.Error, "Data array out of bound ", len(message.RawData), startIndex+field.Size)
		return nil, 0
	}
	switch field.Size {
	case 1:
		bValue := byte(message.RawData[startIndex])
		if field.Name != core.Reason {
			return bValue, field.Size
		}
		if bValue == byte(0xff) || bValue == byte(0xfe) {
			return message.String(startIndex + 1)
		}
		return bValue, field.Size
	case 2:
		return int32(message.RawData[startIndex])<<8 | int32(message.RawData[startIndex+1]), field.Size
	case 4:
		return int32(message.RawData[startIndex])<<24 | int32(message.RawData[startIndex+1])<<16 | int32(message.RawData[startIndex+2])<<8 | int32(message.RawData[startIndex+3]), field.Size
	default:
		return message.RawData[startIndex : startIndex+field.Size], field.Size
	}
}

//3338373030303600000057b800045fa3c2510a75baac0810ea6c0001b3bf00060158010207bf000063a22f140000000000000000003338373030303600000057b900025fa3c28d0a75bada0810e84b0001b3bf0006013f010208c1000063a22f130000000000000000003338373030303600000057ba00025fa3c2c80a75ba4e0810ea770001b3f20006004f010205bf000063a22f110000000000000000003338373030303600000057bb00015fa3c2d90a75ba2d0810ea910001b3f200100036010205c1000063a22f0b000000000000000000
