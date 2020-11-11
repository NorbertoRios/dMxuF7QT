package observers

import (
	"container/list"
	"genx-go/core/device/interfaces"
)

//NewAttachObserverCommand ..
func NewAttachObserverCommand(_observer interfaces.IObserver) *AttachObserverCommand {
	return &AttachObserverCommand{
		observer: _observer,
	}
}

//AttachObserverCommand ..
type AttachObserverCommand struct {
	observer interfaces.IObserver
}

//Execute ..
func (c *AttachObserverCommand) Execute(device interfaces.IDevice) *list.List {
	device.GetObservable().Attach(c.observer)
	return nil
}
