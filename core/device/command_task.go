package device

import (
	"genx-go/logger"
	"genx-go/message"
	"genx-go/message/messagetype"
	"time"
)

var taskTypesMap map[string]string = map[string]string{
	Diag1Wire: messagetype.Diag1Wire,
	DiagCAN:   messagetype.DiagCAN,
	DiagJBUS:  messagetype.DiagJBUS,
}

//CommandTask single command task
type CommandTask struct {
	id                      string
	TaskType                string
	Command                 string
	Device                  IDevice
	SendAt                  time.Time
	OnTaskCompleted         func(string)
	State                   int
	currentParameterMessage *message.ParametersMessage
}

func command(taskType string) string {
	switch taskType {
	case Diag1Wire:
		return "DIAG 1WIRE"
	case DiagCAN:
		return "DIAG CAN"
	case DiagJBUS:
		return "DIAG JBUSSTAT VBRIEF"
	default:
		return ""
	}
}

//BuildCommandTask new command task
func BuildCommandTask(taskType, callbackID string, device IDevice, onTaskCompleted func(string)) *CommandTask {
	command := command(taskType)
	if command == "" {
		logger.Error("[CommandTask | BuildCommandTask] Cant create command task. Command is empty. Task type :", taskType)
		return nil
	}
	return &CommandTask{
		id:              callbackID,
		TaskType:        taskType,
		Device:          device,
		OnTaskCompleted: onTaskCompleted,
		State:           Opened,
		Command:         command,
	}
}

//CallbackID return task callback id
func (task *CommandTask) CallbackID() string {
	return task.id
}

//Execute task
func (task *CommandTask) Execute() {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("[CommandTask] panic:Recovered in Execute:", r)
		}
	}()
	switch task.State {
	case Opened:
		{
			if err := task.Device.Send(task.Command); err != nil {
				logger.Error("[CommandTask] Error while command send. ", err)
				return
			}
			task.State = Sended
			return
		}
	case Sended:
		{
			if time.Now().UTC().Sub(task.SendAt).Seconds() > 60 {
				task.State = Opened
				task.Execute()
				return
			}
			return
		}
	case ParametersReceived:
		{
			task.Complete()
		}
	}
}

//DeviceResponce on device responce
func (task *CommandTask) DeviceResponce(responce interface{}) {
	if task.State != Sended {
		return
	}
	switch responce.(type) {
	case *message.ParametersMessage:
		{
			if v, f := responce.(*message.ParametersMessage); f {
				task.currentParameterMessage = v
				task.State = ParametersReceived
				task.Execute()
			}
		}
	}
}

func (task *CommandTask) onParametersReceivedState() {
	if v, f := taskTypesMap[task.TaskType]; f {
		if task.currentParameterMessage.MessageType == v {
			task.State = ParametersReceived
			task.Execute()
		}
	}
	return
}

//Complete calls on task complete
func (task *CommandTask) Complete() {
	defer func() {
		task.Device = nil
	}()
	task.OnTaskCompleted(task.TaskType)
}
