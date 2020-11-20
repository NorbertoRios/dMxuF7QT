package configuration

import (
	"bytes"
	"container/list"
	"encoding/json"
	"genx-go/logger"
	"genx-go/types"
	"net/http"
)

//NewFacadeRequest ...
func NewFacadeRequest(_identity string, _sentConfigs, _newConfigs *list.List) *FacadeRequest {
	sentArray := types.NewList(_sentConfigs)
	newArray := types.NewList(_newConfigs)
	return &FacadeRequest{
		Identity:    _identity,
		SentConfigs: sentArray.StringArray(),
		NewConfigs:  newArray.StringArray(),
	}
}

//FacadeRequest ...
type FacadeRequest struct {
	Identity    string
	SentConfigs []string
	NewConfigs  []string
}

//Request ...
func (request *FacadeRequest) Request() *http.Request {
	requestBody, err := json.Marshal(request)
	if err != nil {
		logger.Logger().WriteToLog("[FacadeRequest | Request] Error while FacadeRequest marshaling. Error: ", err)
		return nil
	}
	facadeRequest, err := http.NewRequest("POST", "facade url", bytes.NewBuffer(requestBody))
	if err != nil {
		logger.Logger().WriteToLog("[FacadeRequest | Request] Error while request creating. Error: ", err)
		return nil
	}
	return facadeRequest
}
