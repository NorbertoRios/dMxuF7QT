package watchdog

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/usecase"
	"genx-go/logger"
	"time"
)

//BaseWatchdog ...
type BaseWatchdog struct {
	task        interfaces.ITask
	duration    int
	stopChannel chan struct{}
}

func (bw *BaseWatchdog) commands() *list.List {
	logger.Logger().WriteToLog(logger.Fatal, "[BaseWatchdog | commands] Unexpected call in base watchdog")
	return list.New()
}

//Start ...
func (bw *BaseWatchdog) Start() {
	go func() {
		ticker := time.NewTicker(time.Duration(bw.duration) * time.Second)
		for {
			select {
			case <-ticker.C:
				{
					ticker.Stop()
					usecase.NewBaseUseCase(bw.task.Device(), bw.commands()).Launch()
					return
				}
			case <-bw.stopChannel:
				{
					ticker.Stop()
					return
				}
			}
		}
	}()
}

//Stop ...
func (bw *BaseWatchdog) Stop() {
	bw.stopChannel <- struct{}{}
}
