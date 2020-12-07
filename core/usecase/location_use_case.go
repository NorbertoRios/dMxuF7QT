package usecase

import (
	"genx-go/core/device/interfaces"
	"genx-go/core/request"
)

//NewLocationUseCase ...
func NewLocationUseCase(_device interfaces.IDevice, _req *request.BaseRequest) *LocationUseCase {
	useCase := &LocationUseCase{
		caseRequest: _req,
	}
	useCase.device = _device
	return useCase
}

//LocationUseCase ...
type LocationUseCase struct {
	BaseUseCase
	caseRequest *request.BaseRequest
}

//Launch ...
func (lCase *LocationUseCase) Launch() {
	lRequest := lCase.device.LocationRequest()
	commands := lRequest.NewRequest(lCase.caseRequest)
	lCase.execute(commands)
}
