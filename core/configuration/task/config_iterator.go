package task

import (
	"container/list"
	"genx-go/core/configuration/request"
)

func newConfigIterator(_commands *list.List) *ConfigIterator {
	return &ConfigIterator{
		commands:       _commands,
		currentCommand: _commands.Front(),
	}
}

//ConfigIterator ...
type ConfigIterator struct {
	commands       *list.List
	currentCommand *list.Element
}

func (i *ConfigIterator) nextExisting() bool {
	return i.current.Next() != nil
}

func (i *ConfigIterator) goToNext() {
	if !i.nextExisting() {
		return ""
	}
	i.currentCommand = i.currentCommand.Next()
}

func (i *ConfigIterator) current() *request.Command {
	i.currentCommand.Value.(*request.Command)
}
