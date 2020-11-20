package request

//NewRequest ..
func NewRequest(_identity string, _config []string) *ConfigurationRequest {
	return &ConfigurationRequest{
		identity: _identity,
		config:   _config,
	}
}

//ConfigurationRequest ...
type ConfigurationRequest struct {
	identity string
	config   []string
}

//Commands ...
func (request *ConfigurationRequest) Commands() []string {
	return request.config
}

//Identity ...
func (request *ConfigurationRequest) Identity() string {
	return request.identity
}
