package request

import "container/list"

//NewConfig ...
func NewConfig(_config []string) *Config {
	return &Config{
		rawConfig: _config,
	}
}

//Config ...
type Config struct {
	rawConfig []string
}

//List ...
func (config *Config) List() *list.List {
	cList := list.New()
	for _, s := range config.rawConfig {
		cList.PushBack(NewCommand(s))
	}
	return cList
}
