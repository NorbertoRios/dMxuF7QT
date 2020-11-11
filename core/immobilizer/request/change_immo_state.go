package request

import (
	"genx-go/core/request"
	"reflect"
)

//ChangeImmoStateRequest ...
type ChangeImmoStateRequest struct {
	request.ChangeStateRequest
	State        string `json:"ExpectedState"`
	Trigger      string `json:"Trigger"`
	SafetyOption bool   `json:"Safety"`
}

//Equal ...
func (data *ChangeImmoStateRequest) Equal(req request.IRequest) bool {
	if _, v := req.(*ChangeImmoStateRequest); v {
		return reflect.DeepEqual(data, req)
	}
	return false
}
