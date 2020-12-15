package watchdog

import (
	"container/list"
	"genx-go/core/device/interfaces"
)

type ConfigWatchdog struct {
	BaseWatchdog
}

func (cwd *ConfigWatchdog) commands() *list.List {
	cmdList := list.New()
	cmdList.PushBackList(cwd.task.Invoker().(interfaces.IConfigInvoker).SendConfigAfterAnyMessage(awd.task))
	return cmdList
}
