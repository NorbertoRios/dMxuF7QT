package locationmessage

import (
	"container/list"
	"fmt"
	"genx-go/core/device/interfaces"
	"genx-go/core/locationmessage/request"
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
func (p *Process) MessageIncome(req *request.MessageRequest, IUow interface{}) {
	uow := IUow.(unitofwork.IDeviceUnitOfWork)
	requestData := req.Data()
	switch requestData.(type) {
	case *message.RawMessage:
		{
			if _, s := p.ProcessCurrentTask.(*task.LocationMessageTask); s {
				locationMessage := p.messageParser.Parse(requestData.(*message.RawMessage)).(*message.LocationMessage)
				for _, location := range locationMessage.Messages {
					p.device.MessageArrived(location)
					uow.UpdateState(location.Identity, p.device)
					uow.UpdateActivity(location.Identity, p.device)
				}
			} else {
				logger.Logger().WriteToLog(logger.Info, "[SyncProcess | MessageIncome] Receipt location message from unsynchronized device: ", req.Identity)
				p.device.MessageArrived(requestData)
			}
		}
	default:
		{
			logger.Logger().WriteToLog(logger.Info, "[SyncProcess | MessageIncome] Receipt message from unsynchronized device: ", req.Identity)
			p.device.MessageArrived(requestData)
		}
	}
}
