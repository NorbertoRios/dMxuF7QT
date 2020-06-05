package message

import (
	"genx-go/configuration"
	"genx-go/core"
	"log"
)

//RawMessage raw genx message
type RawMessage struct {
	SerialNumber string
	MessageType  string
	RawData      []byte
}

//StringSlice return slice for string values
func (message *RawMessage) String(startIndex int) (interface{}, int) {
	for i, v := range message.RawData {
		if v == 0 {
			return string(message.RawData[startIndex:i]), i + 1
		}
	}
	return message.RawData[startIndex:], len(message.RawData[startIndex:])
}

//Ack returns message Ack
// func (message *RawMessage) Ack() string {
// 	return fmt.Sprintf("UDPACK %v", len(message.RawData))
// }

//Value returns value as intreface{}
func (message *RawMessage) Value(startIndex int, field *configuration.Field) (interface{}, int) {
	if len(message.RawData) < startIndex+field.Size {
		log.Println("Data array out of bound")
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
