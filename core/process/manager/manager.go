package manager

import (
	"genx-go/core/device/interfaces"
	"sync"
)

//NewProcessManager ...
func NewProcessManager() *ProcessManager {
	return &ProcessManager{}
}

//ProcessManager ...
type ProcessManager struct {
	processes map[string]interfaces.IProcess
	mutex     *sync.Mutex
}
