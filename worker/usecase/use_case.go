package usecase

import (
	"genx-go/core/device/interfaces"
	"genx-go/logger"
	"genx-go/message"
	"genx-go/message/messagetype"
	"genx-go/parser"
)

//NewMessageIncomeUseCase ...
func NewMessageIncomeUseCase(_rawMessage *message.RawMessage, _device interfaces.IDevice) *MessageIncomeUseCase {
	return &MessageIncomeUseCase{
		rawMessage: _rawMessage,
		device:     _device,
	}
}

//MessageIncomeUseCase ...
type MessageIncomeUseCase struct {
	device     interfaces.IDevice
	rawMessage *message.RawMessage
}

//Launch ...
func (miu *MessageIncomeUseCase) Launch() interface{} {
	switch miu.rawMessage.MessageType {
	case messagetype.BinaryLocation:
		{
			messageParser := miu.device.Parser()
			return messageParser.Parse(miu.rawMessage)
		}
	case messagetype.Ack:
		{
			messageParser := parser.ConstructAckMesageParser()
			return messageParser.Parse(miu.rawMessage)
		}
	case messagetype.Nack:
		{
			//NackMessage
			return nil
		}
	case messagetype.Parameter:
		{
			messageParser := parser.ConstructParametersMessageParser()
			return messageParser.Parse(miu.rawMessage)
		}
	case messagetype.Poll:
		{
			//Poll message
			return nil
		}
	case messagetype.DiagHardware:
		{
			messageParser := parser.BuildGenxHardwareMessageParser()
			return messageParser.Parse(miu.rawMessage)
		}
	case messagetype.Diag1Wire:
		{
			messageParser := parser.BuildOneWireMessageParser()
			return messageParser.Parse(miu.rawMessage)
		}
	case messagetype.GarminMessage:
		{
			//Garmin
			return nil
		}
	case messagetype.DiagCAN:
		{
			messageParser := parser.BuildCANMessageParser()
			return messageParser.Parse(miu.rawMessage)
		}
	case messagetype.DiagJBUS:
		{
			//JBus
			return nil
		}
	case messagetype.Diag:
		{
			//Diag
			return nil
		}
	default:
		{
			logger.Logger().WriteToLog(logger.Error, "[MessageArrivedUseCase | prepareMessage] Unexpected packet : \"", miu.rawMessage.RawData, "\" message type ", miu.rawMessage.MessageType)
			return nil
		}
	}
}
