package locationmessage

import (
	"container/list"
	"fmt"
	"genx-go/core/device/interfaces"
	"genx-go/core/locationmessage/task"
	"genx-go/core/process"
	"genx-go/logger"
	"genx-go/message"
	"genx-go/parser"
	"genx-go/unitofwork"
)

//NewSynchProcess ...
func NewSynchProcess(_device interfaces.IDevice) *Process {
	return &Process{
		device: _device,
	}
}

//NewLocationProcess ...
func NewLocationProcess(_device interfaces.IDevice, parameter24 []string) *Process {
	return &Process{
		device:        _device,
		messageParser: parser.NewGenxBinaryReportParser(parameter24),
	}
}

//Process ...
type Process struct {
	process.BaseProcess
	device        interfaces.IDevice
	messageParser parser.IParser
}

//Start ...
func (p *Process) Start() *list.List {
	return p.ProcessCurrentTask.Commands()
}

//Param24Arriver ...
func (p *Process) Param24Arriver(param24 []string) *list.List {
	cmdList := list.New()
	p.messageParser = parser.NewGenxBinaryReportParser(param24)
	cmdList.PushBackList(p.ProcessCurrentTask.Invoker().CanselTask(p.ProcessCurrentTask, fmt.Sprintf("New 24 parameter arrived : %v", param24)))
	newTask := task.NewLocationMessageTask(p, p.device)
	cmdList.PushBackList(newTask.Commands())
	return cmdList
}

//MessageIncome ...
func (p *Process) MessageIncome(msg *message.RawMessage, IUow interface{}) {
	uow := IUow.(unitofwork.IDeviceUnitOfWork)
	switch p.ProcessCurrentTask.(type) {
	case *task.LocationMessageTask:
		{
			locationMessage := p.messageParser.Parse(msg).(*message.LocationMessage)
			for _, location := range locationMessage.Messages {
				p.device.MessageArrived(location)
				uow.UpdateState(msg.Identity(), p.device)
				uow.UpdateActivity(msg.Identity(), p.device)
			}
			p.device.Send(locationMessage.Ack)
		}
	case *task.SyncTask:
		{
			logger.Logger().WriteToLog(logger.Info, "[Process | MessageIncome] Cant handle message from device. Device aint synchronized.")
		}
	}
}
