package usecase

import (
	"genx-go/core/device/interfaces"
	"genx-go/logger"
	"genx-go/message"
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

//Launch ...
func (miu *MessageIncomeUseCase) Launch() {
	process := miu.device.LocationMessageProcess()
	if process == nil {
		logger.Logger().WriteToLog(logger.Info, "[MessageIncomeUseCase | locationMessageIncome] Cant find location message process on device ", miu.rawMessage.Identity())
		return
	}
	preparedMessage := miu.prepareMessage()
	switch preparedMessage.(type) {
	case *message.RawMessage:
		{
			rm := preparedMessage.(*message.RawMessage)
			parser := miu.device.Parser()
			if parser == nil {
				logger.Logger().WriteToLog(logger.Info, "[MessageIncomeUseCase | Launch] Parser is nul, location message wont be processed. Device: ", rm.Identity())
				process.MessageIncome(rm, miu.device)
				return
			}
			msg := parser.Parse(rm).(*message.LocationMessage)
			for _, m := range msg.Messages {
				process.MessageIncome(m, miu.device)
				miu.uow.UpdateState(rm.Identity(), miu.device)
			}
		}
	default:
		{
			process.MessageIncome(preparedMessage, miu.device)
		}
	}
	miu.uow.UpdateActivity(miu.rawMessage.Identity(), miu.device)
}

func (miu *MessageIncomeUseCase) prepareMessage() interface{} {
	prepare := NewPrepareMessage(miu.rawMessage)
	return prepare.PreparedMessage()
}
