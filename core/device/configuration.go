package device

//BuildCurrentDeviceConfig returns *CurrentConfig from config model
// func BuildCurrentDeviceConfig(configModel *models.ConfigurationModel) *CurrentConfig {
// 	if configModel == nil {
// 		loggerPrintln("[BuildCurrentDeviceConfig] Cant parse current device config from database.")
// 		return nil
// 	}
// 	re := regexp.MustCompile(`(\n)|(\r\n)`)
// 	configurations := re.Split(configModel.Command, -1)
// 	cfgItems := make(map[string]string, 0)
// 	for _, cfg := range configurations {
// 		if !strings.Contains(cfg, "=") {
// 			continue
// 		}
// 		splitedCfg := strings.Split(cfg, "=")
// 		cfgName := splitedCfg[0]
// 		cfgValue := splitedCfg[1]
// 		if len(cfgName) == 0 || len(cfgValue) == 0 {
// 			continue
// 		}
// 		cfgItems[cfgName] = strings.ReplaceAll(cfgValue, ";", "")
// 	}
// 	if len(cfgItems) == 0 {
// 		return nil
// 	}
// 	return &CurrentConfig{
// 		Values: cfgItems,
// 	}
// }

//CurrentConfig current device config
// type CurrentConfig struct {
// 	Values map[string]string
// }
