package genxmessage

import (
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
func (message *RawMessage) Value(startIndex int, size int) interface{} {
	if len(message.RawData) < startIndex+size {
		log.Println("Data array out of bound")
		return nil
	}
	switch size {
	case 1:
		return byte(message.RawData[startIndex])
	case 2:
		return int32(message.RawData[startIndex])<<8 | int32(message.RawData[startIndex+1])
	case 4:
		return int32(message.RawData[startIndex])<<24 | int32(message.RawData[startIndex+1])<<16 | int32(message.RawData[startIndex+2])<<8 | int32(message.RawData[startIndex+3])
	default:
		return message.RawData[startIndex : startIndex+size]
	}
}
