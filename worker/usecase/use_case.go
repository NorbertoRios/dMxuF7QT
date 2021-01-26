package usecase

import (
	"genx-go/core/device/interfaces"
	"genx-go/core/locationmessage/request"
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
	req := request.NewMessageRequest(miu.rawMessage)
	process.MessageIncome(req, miu.uow)
}
