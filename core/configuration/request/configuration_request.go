package request

import (
	"container/list"
	"genx-go/core/request"
	"reflect"
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

//Equal ...
func (request *ConfigurationRequest) Equal(req request.IRequest) bool {
	if _, v := req.(*ConfigurationRequest); v {
		return reflect.DeepEqual(request, req)
	}
	return false
}
