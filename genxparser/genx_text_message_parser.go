package genxparser

import "genx-go/genxmessage"

//GenxTextMessageParser  parser for text messages
type GenxTextMessageParser struct {
}

//Parse returns GenxMessage
func (parser *GenxTextMessageParser) Parse(rawMessage *genxmessage.RawMessage) *genxmessage.GenxMessage {
	switch rawMessage.MessageType {
	case "ack":
		{

		}
	case "nack":
		{

		}
	case "param":
		{

		}
	case "poll":
		{

		}
	case "diag_hardware_brief":
		{

		}
	case "diag_1wire":
		{

		}
	case "diag":
		{

		}
	case "message":
		{

		}
	}
	return nil
}
