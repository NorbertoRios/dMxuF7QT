package usecase

import (
	"genx-go/core/device/interfaces"
	"genx-go/logger"
	"genx-go/message"
	"genx-go/message/messagetype"
	"genx-go/parser"
	"genx-go/unitofwork"
)

//NewMessageIncomeUseCase ...
func NewMessageIncomeUseCase(_rawMessage *message.RawMessage, _device interfaces.IDevice, _uow unitofwork.IDeviceUnitOfWork) *MessageIncomeUseCase {
	return &MessageIncomeUseCase{
		rawMessage: _rawMessage,
		device:     _device,
		uow:        _uow,
	}
}

//MessageIncomeUseCase ...
type MessageIncomeUseCase struct {
	device     interfaces.IDevice
	rawMessage *message.RawMessage
	uow        unitofwork.IDeviceUnitOfWork
}

func (miu *MessageIncomeUseCase) locationMessageIncome() {
	process := miu.device.LocationMessageProcess()
	if process == nil {
		logger.Logger().WriteToLog(logger.Info, "[MessageIncomeUseCase | locationMessageIncome] Cant find location message process on device ", miu.rawMessage.Identity())
		return
	}
	process.MessageIncome(miu.rawMessage, miu.uow)
}

//Launch ...
func (miu *MessageIncomeUseCase) Launch() {
	switch miu.rawMessage.MessageType {
	case messagetype.BinaryLocation:
		{
			miu.locationMessageIncome()
		}
	case messagetype.Ack:
		{
			messageParser := parser.ConstructAckMesageParser()
			miu.device.MessageArrived(messageParser.Parse(miu.rawMessage))
		}
	case messagetype.Nack:
		{
			//NackMessage
		}
	case messagetype.Parameter:
		{
			messageParser := parser.ConstructParametersMessageParser()
			miu.device.MessageArrived(messageParser.Parse(miu.rawMessage))
		}
	case messagetype.Poll:
		{
			//Poll message
		}
	case messagetype.DiagHardware:
		{
			messageParser := parser.BuildGenxHardwareMessageParser()
			miu.device.MessageArrived(messageParser.Parse(miu.rawMessage))
		}
	case messagetype.Diag1Wire:
		{
			messageParser := parser.BuildOneWireMessageParser()
			miu.device.MessageArrived(messageParser.Parse(miu.rawMessage))
		}
	case messagetype.GarminMessage:
		{
			//Garmin
		}
	case messagetype.DiagCAN:
		{
			messageParser := parser.BuildCANMessageParser()
			miu.device.MessageArrived(messageParser.Parse(miu.rawMessage))
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
			logger.Logger().WriteToLog(logger.Error, "[MessageArrivedUseCase | prepareMessage] Unexpected packet : \"", miu.rawMessage.RawData, "\" message type ", miu.rawMessage.MessageType)
		}
	}
}
