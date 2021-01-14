package usecase

import (
	"genx-go/core/device/interfaces"
	immoRequest "genx-go/core/immobilizer/request"
	"genx-go/core/request"
)

//NewImmobilizerUseCase ...
func NewImmobilizerUseCase(_device interfaces.IDevice, _req *immoRequest.ChangeImmoStateRequest) *ImmobilizerUseCase {
	uCase := &ImmobilizerUseCase{
		caseRequest: _req,
	}
	uCase.device = _device
	return uCase
}

//ImmobilizerUseCase ...
type ImmobilizerUseCase struct {
	BaseUseCase
	caseRequest *immoRequest.ChangeImmoStateRequest
}

//Launch use case
func (iCase *ImmobilizerUseCase) Launch() {
	outNumber := &request.OutputNumber{Data: iCase.caseRequest.Port}
	immo := iCase.device.Immobilizer(outNumber.Index(), iCase.caseRequest.Trigger)
	commands := immo.NewRequest(iCase.caseRequest, iCase.device)
	iCase.execute(commands)
}
