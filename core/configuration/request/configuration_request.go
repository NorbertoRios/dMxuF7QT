package request

import (
	"container/list"
	"genx-go/core/request"
)

//ConfigurationRequest ...
type ConfigurationRequest struct {
	request.BaseRequest
	Config []string `json:"Configs"`
}

//Commands ...
func (request *ConfigurationRequest) Commands() *list.List {
	config := NewConfig(request.Config)
	return config.List()
}
