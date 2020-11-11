package observers

import (
	"fmt"
	"genx-go/core/immobilizer/request"
	bRequest "genx-go/core/request"
)

//NewSetRelayDrive ...
func NewSetRelayDrive(_request *request.ChangeImmoStateRequest) *SetRelayDrive {
	return &SetRelayDrive{
		request: _request,
	}
}

//SetRelayDrive ...
type SetRelayDrive struct {
	request *request.ChangeImmoStateRequest
}

//Command ...
func (srd *SetRelayDrive) Command() string {
	action := &request.ActionString{Data: srd.request}
	output := &bRequest.OutputNumber{Data: srd.request.Port}
	return fmt.Sprintf("SETRELAYDRIVE%v%v SERIALFILTER %v;BACKUPNVRAM", output.Index(), action.Command(), srd.request.Serial())
}
