package watchdog

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/usecase"
	"time"
)

//NewWatchdog ...
func NewWatchdog(_device interfaces.IDevice, _commands *list.List, _duration int) *Watchdog {
	return &Watchdog{
		device:      _device,
		commands:    _commands,
		duration:    _duration,
		stopChannel: make(chan struct{}),
	}
}

//Watchdog ...
type Watchdog struct {
	device      interfaces.IDevice
	commands    *list.List
	duration    int
	stopChannel chan struct{}
}

//Start ...
func (w *Watchdog) Start() {
	go func() {
		ticker := time.NewTicker(time.Duration(w.duration) * time.Second)
		for {
			select {
			case <-ticker.C:
				{
					ticker.Stop()
					usecase.NewBaseUseCase(w.device, w.commands).Launch()
					return
				}
			case <-w.stopChannel:
				{
					ticker.Stop()
					return
				}
			}
		}
	}()
}

//Stop ...
func (w *Watchdog) Stop() {
	w.stopChannel <- struct{}{}
}
