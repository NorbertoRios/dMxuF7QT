package request

import (
	"genx-go/logger"
	"regexp"
	"strings"
)

//NewStringConfigByArray ...
func NewStringConfigByArray(_configArr []string) *StringConfig {
	_config := strings.Join(_configArr, "\n")
	return NewStringConfig(_config)
}

//NewStringConfig ...
func NewStringConfig(_config string) *StringConfig {
	return &StringConfig{
		config: _config,
	}
}

//StringConfig stringConfig for configuration
type StringConfig struct {
	config string
}

//ParametersArray ...
func (stringConfig *StringConfig) ParametersArray() []string {
	re := regexp.MustCompile(`(\n)|(\r\n)`)
	return re.Split(stringConfig.config, -1)
}

//Parameters returns configs as map[string]string
func (stringConfig *StringConfig) Parameters() map[string]string {
	responce := make(map[string]string)
	configurations := stringConfig.ParametersArray()
	for _, cfg := range configurations {
		if len(cfg) == 0 ||
			strings.Contains(strings.ToUpper(cfg), "SETPARAM") ||
			strings.Contains(strings.ToUpper(cfg), "ENDPARAM") ||
			strings.Contains(strings.ToUpper(cfg), "BACKUPNVRAM") {
			continue
		}
		if strings.Contains(strings.ToUpper(cfg), "SETBOUNDARY") {
			boundaries := strings.Split(cfg, " ")
			responce[boundaries[0]+boundaries[1]] = cfg
			continue
		}
		cfgName := strings.Split(cfg, "=")[0]
		if len(cfgName) == 0 {
			continue
		}
		responce[cfgName] = cfg
	}
	return responce
}

//ParameterByName returns parameters by name
func (stringConfig *StringConfig) ParameterByName(parameters ...string) map[string]string {
	config := stringConfig.Parameters()
	response := make(map[string]string)
	for _, param := range parameters {
		if _, f := config[param]; !f {
			logger.Logger().WriteToLog(logger.Error, "[StringConfig | ParameterByName] Сould not find \"", param, "\" in configuration")
			continue
		}
		response[param] = config[param]
	}
	return response
}
