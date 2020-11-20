package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"genx-go/core/configuration/request"
	"genx-go/logger"
	"io/ioutil"
	"net/http"
)

//NewRequestToConfiguratorService ...
func NewRequestToConfiguratorService(_request *ConfiguratorRequest) *RequestToConfiguratorService {
	return &RequestToConfiguratorService{
		request: _request,
	}
}

//RequestToConfiguratorService ...
type RequestToConfiguratorService struct {
	request *ConfiguratorRequest
}

//Execute ...
func (configRequest *RequestToConfiguratorService) Execute() []string {
	requestBody, bodyErr := json.Marshal(configRequest.request)
	if bodyErr != nil {
		return configRequest.errorHandler(fmt.Sprintf("[RequestToConfiguratorService | Execute] Error while marshaling. Error : %v", bodyErr), configRequest.request.OldConfig)
	}
	resp, err := http.Post("http://10.128.0.13:7980/config/merge_configs", "application/x-www-form-urlencoded", bytes.NewBuffer(requestBody))
	if err != nil {
		return configRequest.errorHandler(fmt.Sprintf("[RequestToConfiguratorService | Execute] Error while executing request. Error %v: ", err), configRequest.request.OldConfig)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 302 {
		return configRequest.errorHandler(fmt.Sprintf("[RequestToConfiguratorService | Execute] Unexpected status code : %v", resp.StatusCode), configRequest.request.OldConfig)
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return configRequest.errorHandler(fmt.Sprintf("[RequestToConfiguratorService | Execute] Error while reading response body : %v", err), configRequest.request.OldConfig)
	}
	config := request.NewStringConfig(string(responseBody))
	return config.ParametersArray()
}

func (configRequest *RequestToConfiguratorService) errorHandler(description string, returnedValue []string) []string {
	logger.Logger().WriteToLog(logger.Error, description)
	return returnedValue
}
