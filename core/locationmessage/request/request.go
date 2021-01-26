package request

import (
	"genx-go/logger"
	"genx-go/message"
	"genx-go/message/messagetype"
	"genx-go/parser"
)

//NewMessageRequest ...
func NewMessageRequest(_rawMessage *message.RawMessage) *MessageRequest {
	return &MessageRequest{
		Identity:   _rawMessage.Identity(),
		rawMessage: _rawMessage,
	}
}

//MessageRequest ...
type MessageRequest struct {
	Identity   string
	rawMessage *message.RawMessage
}

//Data ...
func (mr *MessageRequest) Data() interface{} {
	switch mr.rawMessage.MessageType {
	case messagetype.BinaryLocation:
		{
			return mr.rawMessage
		}
	case messagetype.Ack:
		{
			messageParser := parser.ConstructAckMesageParser()
			return messageParser.Parse(mr.rawMessage)
		}
	case messagetype.Nack:
		{
			//NackMessage
		}
	case messagetype.Parameter:
		{
			messageParser := parser.ConstructParametersMessageParser()
			return messageParser.Parse(mr.rawMessage)
		}
	case messagetype.Poll:
		{
			//Poll message
		}
	case messagetype.DiagHardware:
		{
			messageParser := parser.BuildGenxHardwareMessageParser()
			return messageParser.Parse(mr.rawMessage)
		}
	case messagetype.Diag1Wire:
		{
			messageParser := parser.BuildOneWireMessageParser()
			return messageParser.Parse(mr.rawMessage)
		}
	case messagetype.GarminMessage:
		{
			//Garmin
		}
	case messagetype.DiagCAN:
		{
			messageParser := parser.BuildCANMessageParser()
			return messageParser.Parse(mr.rawMessage)
		}
	case messagetype.DiagJBUS:
		{
			//JBus
		}
	case messagetype.Diag:
		{
			//Diag
		}
	default:
		{
			logger.Logger().WriteToLog(logger.Error, "[MessageArrivedUseCase | prepareMessage] Unexpected packet : \"", mr.rawMessage.RawData, "\" message type ", mr.rawMessage.MessageType)
		}
	}
	return struct{}{}
}
