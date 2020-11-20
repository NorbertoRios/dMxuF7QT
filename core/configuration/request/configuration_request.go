package request

import (
	"container/list"
	"genx-go/core/request"
	"genx-go/types"
)

//ConfigurationRequest ...
type ConfigurationRequest struct {
	request.BaseRequest
	Config []string `json:"Configs"`
}

//Commands ...
func (request *ConfigurationRequest) Commands() *list.List {
	sType := &types.StringArray{Data: request.Config}
	return sType.List()
}
