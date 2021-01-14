package usecase

import (
	"genx-go/core/device/interfaces"
	lockRequest "genx-go/core/lock/request"
	"genx-go/core/request"
)

//NewLockUseCase ..
func NewLockUseCase(_device interfaces.IDevice, _req *lockRequest.UnlockRequest) *LockUseCase {
	lCase := &LockUseCase{
		caseRequest: _req,
	}
	lCase.device = _device
	return lCase
}

//LockUseCase ...
type LockUseCase struct {
	BaseUseCase
	caseRequest *lockRequest.UnlockRequest
}

//Launch ...
func (lCase *LockUseCase) Launch() {
	outNumber := &request.OutputNumber{Data: lCase.caseRequest.Port}
	lock := lCase.device.ElectricLock(outNumber.Index())
	commands := lock.NewRequest(lCase.caseRequest, lCase.device)
	lCase.execute(commands)
}
