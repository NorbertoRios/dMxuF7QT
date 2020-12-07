package usecase

import (
	"container/list"
	"fmt"
	"genx-go/core/device/interfaces"
	"genx-go/logger"
)

//NewBaseUseCase ...
func NewBaseUseCase(_device interfaces.IDevice, _req *list.List) *BaseUseCase {
	return &BaseUseCase{
		device:  _device,
		request: _req,
	}
}

//BaseUseCase ...
type BaseUseCase struct {
	device  interfaces.IDevice
	request *list.List
}

//Launch ...
func (uCase *BaseUseCase) Launch() {
	uCase.execute(uCase.request)
}

func (uCase *BaseUseCase) execute(commands *list.List) {
	for commands.Len() > 0 {
		cmd := commands.Front()
		command, valid := cmd.Value.(interfaces.ICommand)
		if valid {
			nList := command.Execute(uCase.device)
			if nList != nil && nList.Len() > 0 {
				commands.PushFrontList(nList)
			}
		} else {
			logger.Logger().WriteToLog(logger.Error, fmt.Sprintf("[BaseUseCase | BaseUseCase] Command %T doesn't implement interface ICommand", cmd))
		}
		commands.Remove(cmd)
	}
}
