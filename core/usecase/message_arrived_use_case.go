package usecase

import (
	"genx-go/core/device/interfaces"
	"genx-go/logger"
	"genx-go/message"
	"genx-go/message/messagetype"
	"genx-go/parser"
)

//NewMessageArrivedUseCase ...
func NewMessageArrivedUseCase(_device interfaces.IDevice, _rMessage *message.RawMessage) *MessageArrivedUseCase {
	useCase := &MessageArrivedUseCase{
		rMessage: _rMessage,
	}
	useCase.device = _device
	return useCase
}

//MessageArrivedUseCase ...
type MessageArrivedUseCase struct {
	BaseUseCase
	rMessage *message.RawMessage
}

//Launch ...
func (useCase *MessageArrivedUseCase) Launch() {
	deviceMessage := useCase.prepareMessage()
	commands := useCase.device.MessageArrived(deviceMessage)
	useCase.execute(commands)
}

func (useCase *MessageArrivedUseCase) prepareMessage() interface{} {
	switch useCase.rMessage.MessageType {
	case messagetype.BinaryLocation:
		{
			messageParser := useCase.device.Parser()
			return messageParser.Parse(useCase.rMessage)
		}
	case messagetype.Ack:
		{
			messageParser := parser.ConstructAckMesageParser()
			return messageParser.Parse(useCase.rMessage)
		}
	case messagetype.Nack:
		{
			//NackMessage
			return nil
		}
	case messagetype.Parameter:
		{
			messageParser := parser.ConstructParametersMessageParser()
			return messageParser.Parse(useCase.rMessage)
		}
	case messagetype.Poll:
		{
			//Poll message
			return nil
		}
	case messagetype.DiagHardware:
		{
			messageParser := parser.BuildGenxHardwareMessageParser()
			return messageParser.Parse(useCase.rMessage)
		}
	case messagetype.Diag1Wire:
		{
			messageParser := parser.BuildOneWireMessageParser()
			return messageParser.Parse(useCase.rMessage)
		}
	case messagetype.GarminMessage:
		{
			//Garmin
			return nil
		}
	case messagetype.DiagCAN:
		{
			messageParser := parser.BuildCANMessageParser()
			return messageParser.Parse(useCase.rMessage)
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
			logger.Logger().WriteToLog(logger.Error, "[MessageArrivedUseCase | prepareMessage] Unexpected packet : \"", useCase.rMessage.RawData, "\" message type ", useCase.rMessage.MessageType)
			return nil
		}
	}
}
