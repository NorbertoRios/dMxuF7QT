package observers

import (
	"fmt"
	"strings"
)

//NewConfig ...
func NewConfig(_command string) *Config {
	return &Config{
		command: _command,
	}
}

//Config ...
type Config struct {
	command string
}

//Command ...
func (config *Config) Command() string {
	if config.command == "SETBOUNDARY DELETEALL; ENDBOUNDARY;" {
		return config.command
	}
	if strings.Contains(config.command, "SETBOUNDARY") {
		return fmt.Sprintf("%v BACKUPNVRAM;", config.command)
	}
	return fmt.Sprintf("SETPARAMVERIFY;%v ENDPARAM;BACKUPNVRAM;", config.command)
}
