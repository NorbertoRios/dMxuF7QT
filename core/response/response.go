package response

import (
	"encoding/json"
	"genx-go/logger"
)

//NewResponse ...
func NewResponse(callbackID string, success bool, code string) *Response {
	return &Response{
		CallbackID: callbackID,
		Success:    success,
		Code:       code,
	}
}

//Response immobilizer facade response
type Response struct {
	CallbackID string `json:"CallbackId"`
	Success    bool   `json:"Success"`
	Code       string `json:"Code"`
}

//Marshal responce
func (resp *Response) Marshal() string {
	jmess, jerr := json.Marshal(resp)
	if jerr != nil {
		logger.Logger().WriteToLog(logger.Error, "[Response | Marshal] Error while marshaling responce ", jerr)
		return ""
	}
	return string(jmess)
}
