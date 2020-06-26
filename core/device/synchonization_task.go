package device

import (
	"fmt"
	"genx-go/message"
	"log"
	"strings"
	"time"
)

//BuildSynchronizarionTask build new SynchronizarionTask
func BuildSynchronizarionTask(device IDevice, storage *TaskStorage, onTaskCompleted func(string)) {
	task := &SynchronizarionTask{
		TaskType:      SynchronizationTask,
		TaskStorage:   storage,
		taskCompleted: onTaskCompleted,
		device:        device,
		state:         Opened,
	}
	task.TaskStorage.NewTask(task.TaskType, task)
}

//SynchronizarionTask task for synchronization
type SynchronizarionTask struct {
	TaskType                     string
	taskCompleted                func(string)
	TaskStorage                  *TaskStorage
	currentParameterMessage      *message.ParametersMessage
	diagRequiredParametersSendAt time.Time
	device                       IDevice
	state                        int
}

//CallbackID for tasks which need send responce to facade
func (task *SynchronizarionTask) CallbackID() string {
	return ""
}

//Execute execute task
func (task *SynchronizarionTask) Execute() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[SynchronizarionTask] panic:Recovered in Execute:", r)
		}
	}()
	switch task.state {
	case Opened:
		{
			if err := task.device.Send("DIAG PARAMS=500,24,18;"); err != nil {
				log.Println("[SynchronizarionTask] Error while command send. ", err)
				return
			}
			task.state = DiagRequiredParametersSended
			return
		}
	case DiagRequiredParametersSended:
		{
			if time.Now().UTC().Sub(task.diagRequiredParametersSendAt).Seconds() > 60 {
				task.state = Opened
				task.Execute()
				return
			}
			return
		}
	case ParametersReceived:
		{
			task.onParametersReceivedState()
		}
	case Completed:
		{
			task.Complete()
		}
	default:
		{

		}
	}
}

func (task *SynchronizarionTask) onSubtaskCompleted(taskName string) {
	task.taskCompleted(taskName)
	task.state = Opened
	task.Execute()
}

func (task *SynchronizarionTask) onParametersReceivedState() {
	requiredParameters := []string{"500", "24", "18"}
	for _, rp := range requiredParameters {
		if _, ok := task.currentParameterMessage.Parameters[rp]; !ok {
			task.state = DiagRequiredParametersSended
			task.Execute()
			return
		}
	}
	deviceConfig := task.device.OnLoadCurrentConfig().Command
	for _, key := range requiredParameters {
		if !strings.Contains(deviceConfig, task.currentParameterMessage.Parameters[key]) {
			BuildConfigurationTask(task.TaskStorage, task.TaskStorage.Device.OnLoadCurrentConfig(), task.onSubtaskCompleted)
			log.Println(fmt.Sprint("[SynchronizarionTask] New subtask for push current config to device is  created"))
			task.state = SubtaskIsActive
			task.Execute()
			return
		}
	}
	task.state = Completed
	task.Execute()
}

//DeviceResponce on device responce
func (task *SynchronizarionTask) DeviceResponce(responce interface{}) {
	if task.state != DiagRequiredParametersSended {
		return
	}
	switch responce.(type) {
	case *message.ParametersMessage:
		{
			task.currentParameterMessage = responce.(*message.ParametersMessage)
			task.state = ParametersReceived
			task.Execute()
		}
	}
}

//Complete calls on task complete
func (task *SynchronizarionTask) Complete() {
	defer func() {
		task.TaskStorage = nil
	}()
	task.TaskStorage.Device.OnSynchronizationTaskCompleted()
	task.taskCompleted(task.TaskType)
}
