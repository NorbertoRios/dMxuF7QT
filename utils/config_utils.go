package utils

import (
	"genx-go/logger"
	"regexp"
	"strings"
)

//ConfigurationUtils utils for configuration
type ConfigurationUtils struct {
	Config string
}

//ConfigParameters returns configs as map[string]string
func (utils *ConfigurationUtils) ConfigParameters() map[string]string {
	responce := make(map[string]string)
	re := regexp.MustCompile(`(\n)|(\r\n)`)
	configurations := re.Split(utils.Config, -1)
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
func (utils *ConfigurationUtils) ParameterByName(parameters ...string) map[string]string {
	config := utils.ConfigParameters()
	responce := make(map[string]string)
	for _, param := range parameters {
		if _, f := config[param]; !f {
			logger.Error("[ConfigurationUtils | ParameterByName] Сould not find \"", param, "\" in configuration")
			continue
		}
		responce[param] = config[param]
	}
	return responce
}
