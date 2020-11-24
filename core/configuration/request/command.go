package request

import (
	"fmt"
	"strings"
)

//NewCommand ...
func NewCommand(_command string) *Command {
	return &Command{
		command:   _command,
		sentState: false,
	}
}

//Command ...
type Command struct {
	command   string
	sentState bool
}

//Command ...
func (command *Command) Command() string {
	if command.command == "SETBOUNDARY DELETEALL; ENDBOUNDARY;" {
		return command.command
	}
	if strings.Contains(command.command, "SETBOUNDARY") {
		return fmt.Sprintf("%v BACKUPNVRAM;", command.command)
	}
	return fmt.Sprintf("SETPARAMVERIFY;%vENDPARAM;BACKUPNVRAM;", command.command)
}

//Complete ...
func (command *Command) Complete() {
	command.sentState = true
}

//State ...
func (command *Command) State() bool {
	return command.sentState
}
