package usecase

import (
	"genx-go/core/device/interfaces"
	"genx-go/logger"
	"genx-go/message"
	"genx-go/message/messagetype"
	"genx-go/parser"
)

//NewMessageArrivedUseCase ...
func NewMessageArrivedUseCase(_device interfaces.IDevice, _request []byte) *MessageArrivedUseCase {
	useCase := &MessageArrivedUseCase{
		caseRequest: _request,
	}
	useCase.device = _device
	return useCase
}

//MessageArrivedUseCase ...
type MessageArrivedUseCase struct {
	BaseUseCase
	caseRequest []byte
}

//Launch ...
func (useCase *MessageArrivedUseCase) Launch() {
	deviceMessage := useCase.prepareMessage()
	commands := useCase.device.MessageArrived(deviceMessage)
	useCase.execute(commands)
}

func (useCase *MessageArrivedUseCase) prepareMessage() interface{} {
	factory := message.Factory()
	rMessage := factory.BuildRawMessage(useCase.caseRequest)
	switch rMessage.MessageType {
	case messagetype.BinaryLocation:
		{
			// Binary location message
			return nil
		}
	case messagetype.Ack:
		{
			messageParser := parser.ConstructAckMesageParser()
			return messageParser.Parse(rMessage)
		}
	case messagetype.Nack:
		{
			//NackMessage
			return nil
		}
	case messagetype.Parameter:
		{
			messageParser := parser.ConstructParametersMessageParser()
			return messageParser.Parse(rMessage)
		}
	case messagetype.Poll:
		{
			//Poll message
			return nil
		}
	case messagetype.DiagHardware:
		{
			messageParser := parser.BuildGenxHardwareMessageParser()
			return messageParser.Parse(rMessage)
		}
	case messagetype.Diag1Wire:
		{
			messageParser := parser.BuildOneWireMessageParser()
			return messageParser.Parse(rMessage)
		}
	case messagetype.GarminMessage:
		{
			//Garmin
			return nil
		}
	case messagetype.DiagCAN:
		{
			messageParser := parser.BuildCANMessageParser()
			return messageParser.Parse(rMessage)
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
			logger.Logger().WriteToLog(logger.Error, "[MessageArrivedUseCase | prepareMessage] Unexpected packet : \"", useCase.caseRequest, "\" message type ", rMessage.MessageType)
			return nil
		}
	}
}
