package request

import (
	"genx-go/logger"
	"strings"
)

//NewConfig ...
func NewConfig(_config []string) *Config {
	return &Config{
		config: _config,
	}
}

//Config Config for configuration
type Config struct {
	config []string
}

//Parameters returns configs as map[string]string
func (Config *Config) Parameters() map[string]string {
	responce := make(map[string]string)
	for _, cfg := range Config.config {
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
func (Config *Config) ParameterByName(parameters ...string) map[string]string {
	config := Config.Parameters()
	response := make(map[string]string)
	for _, param := range parameters {
		if _, f := config[param]; !f {
			logger.Logger().WriteToLog(logger.Error, "[Config | ParameterByName] Сould not find \"", param, "\" in configuration")
			continue
		}
		response[param] = config[param]
	}
	return response
}
