package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
)

//NewDetachObserverCommand ..
func NewDetachObserverCommand(_observer interfaces.IObserver) *DetachObserverCommand {
	return &DetachObserverCommand{
		observer: _observer,
	}
}

//DetachObserverCommand ..
type DetachObserverCommand struct {
	observer interfaces.IObserver
}

//Execute ..
func (c *DetachObserverCommand) Execute(device interfaces.IDevice) *list.List {
	device.GetObservable().Detach(c.observer)
	return nil
}