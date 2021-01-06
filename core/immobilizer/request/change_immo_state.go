package request

import (
	"encoding/json"
	"genx-go/core/request"
	"genx-go/logger"
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

//Marshal ...
func (data *ChangeImmoStateRequest) Marshal() string {
	jTask, jerr := json.Marshal(data)
	if jerr != nil {
		logger.Logger().WriteToLog(logger.Error, "[ChangeImmoStateRequest | Marshal] Error while marshaling ChangeImmoStateRequest. Error:", jerr)
		return ""
	}
	return string(jTask)
}
