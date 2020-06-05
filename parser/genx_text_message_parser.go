package parser

import "genx-go/message"

//GenxTextMessageParser  parser for text messages
type GenxTextMessageParser struct {
}

//Parse returns Message
func (parser *GenxTextMessageParser) Parse(rawMessage *message.RawMessage) *message.Message {
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
