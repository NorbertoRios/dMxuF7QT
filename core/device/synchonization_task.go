package device

import (
	"genx-go/message"
	"log"
	"strings"
	"time"
)

//ConstructSynchronizarionTask return new SynchronizarionTask
func ConstructSynchronizarionTask(device IDevice) *SynchronizarionTask {
	return &SynchronizarionTask{
		device: device,
		state:  Opened,
	}
}

//SynchronizarionTask task for synchronization
type SynchronizarionTask struct {
	currentParameterMessage      *message.ParametersMessage
	diagRequiredParametersSendAt time.Time
	device                       IDevice
	state                        int
	needSynchConfig              bool
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
			requiredParameters := []string{"500", "24", "18"}
			deviceConfig := task.device.Config()
			for _, key := range requiredParameters {
				if !strings.Contains(deviceConfig, task.currentParameterMessage.Parameters[key]) {
					task.needSynchConfig = true
					break
				}
			}
			task.state = Completed
			task.Execute()
		}
	case Completed:
		{
			task.Complete()
		}
	}
}

//OnParametersRecieved when parameters recieved
func (task *SynchronizarionTask) OnParametersRecieved(message *message.ParametersMessage) {
	task.currentParameterMessage = message
	task.state = ParametersReceived
	task.Execute()
}

//Complete calls on task complete
func (task *SynchronizarionTask) Complete() {
	defer func() {
		task.device = nil
	}()
	task.device.OnSynchronizationTaskCompleted(task.needSynchConfig)
}

func (task *SynchronizarionTask) compareRequiredParameters() {

}
