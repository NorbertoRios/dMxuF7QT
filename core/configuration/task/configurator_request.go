package task

//NewConfiguratorRequest ...
func NewConfiguratorRequest(identity string, oldConfig, newConfig []string) *ConfiguratorRequest {
	return &ConfiguratorRequest{
		Identity:  identity,
		OldConfig: oldConfig,
		NewConfig: newConfig,
	}
}

//ConfiguratorRequest ...
type ConfiguratorRequest struct {
	Identity  string   `json:"identity"`
	OldConfig []string `json:"config_top"`
	NewConfig []string `json:"config_sub"`
}
