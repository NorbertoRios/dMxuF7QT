package watchdog

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"time"
)

//NewWatchdog ...
func NewWatchdog(_commands *list.List, _device interfaces.IDevice, _duration int) *Watchdog {
	return &Watchdog{
		duration:    _duration,
		device:      _device,
		commands:    _commands,
		stopChannel: make(chan struct{}),
	}
}

//Watchdog ...
type Watchdog struct {
	duration    int
	device      interfaces.IDevice
	commands    *list.List
	stopChannel chan struct{}
}

//Stop ...
func (wd *Watchdog) Stop() {
	wd.stopChannel <- struct{}{}
}

//Start ...
func (wd *Watchdog) Start() {
	go func() {
		ticker := time.NewTicker(time.Duration(wd.duration) * time.Second)
		for {
			select {
			case <-ticker.C:
				{
					ticker.Stop()
					wd.device.ProcessCommands(wd.commands)
					return
				}
			case <-wd.stopChannel:
				{
					ticker.Stop()
					return
				}
			}
		}
	}()
}
