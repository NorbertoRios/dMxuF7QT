package usecase

import (
	"genx-go/core/device/interfaces"
)

//NewMessageArrivedUseCase ...
func NewMessageArrivedUseCase(_device interfaces.IDevice, _message interface{}) *MessageArrivedUseCase {
	useCase := &MessageArrivedUseCase{
		message: _message,
	}
	useCase.device = _device
	return useCase
}

//MessageArrivedUseCase ...
type MessageArrivedUseCase struct {
	BaseUseCase
	message interface{}
}

//Launch ...
func (useCase *MessageArrivedUseCase) Launch() {
	commands := useCase.device.MessageArrived(useCase.message)
	useCase.execute(commands)
}
