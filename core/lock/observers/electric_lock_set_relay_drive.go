package observers

import (
	"fmt"
	"genx-go/core/lock/request"
	bReq "genx-go/core/request"
)

//NewElectricLockSetRelayDrive ...
func NewElectricLockSetRelayDrive(_request *request.UnlockRequest) *ElectricLockSetRelayDrive {
	return &ElectricLockSetRelayDrive{
		request: _request,
	}
}

//ElectricLockSetRelayDrive ...
type ElectricLockSetRelayDrive struct {
	request *request.UnlockRequest
}

//Command ...
func (srd *ElectricLockSetRelayDrive) Command() string {
	output := &bReq.OutputNumber{Data: srd.request.Port}
	return fmt.Sprintf("SETRELAYDRIVE%vX%v SERIALFILTER %v", output.Index(), srd.request.Pulse(), srd.request.Serial())
}
