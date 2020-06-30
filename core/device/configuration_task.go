package device

import (
	"container/list"
	"genx-go/logger"
	"genx-go/message"
	"genx-go/repository/models"
	"genx-go/utils"
	"strings"
	"sync"
	"time"
)

//BuildConfigurationTask create configuration task
func BuildConfigurationTask(callbackID string, device IDevice, config *models.ConfigurationModel, onTaskCompleted func(string)) *ConfigurationTask {
	if config.ID == 0 || len(config.Command) == 0 {
		logger.Error("[BuildConfigurationTask] Error in configuration.")
		return nil
	}
	return &ConfigurationTask{
		id:                 callbackID,
		Device:             device,
		mutex:              &sync.Mutex{},
		ConfigurationItems: devideConfiguration(config.Command),
		OnTaskCompleted:    onTaskCompleted,
	}
}

func devideConfiguration(config string) *list.List {
	items := list.New()
	cfgUtils := &utils.ConfigurationUtils{Config: config}
	for key, value := range cfgUtils.ConfigParameters() {
		item := &ConfigItem{
			Name:  key,
			Value: value,
			State: byte(0),
		}
		items.PushBack(item)
	}
	return items
}

//CallbackID for tasks which need send responce to facade
func (task *ConfigurationTask) CallbackID() string {
	return task.id
}

//ConfigurationTask represents task for send config to device
type ConfigurationTask struct {
	id                 string
	TaskType           string
	mutex              *sync.Mutex
	Device             IDevice
	CurrentItem        *list.Element
	ConfigurationItems *list.List
	SendAt             time.Time
	OnTaskCompleted    func(string)
}

//Execute execute task
func (task *ConfigurationTask) Execute() {
	task.mutex.Lock()
	defer task.mutex.Unlock()
	if task.CurrentItem.Value.(*ConfigItem).State == 3 { //Acked
		task.ConfigurationItems.Remove(task.CurrentItem)
		task.CurrentItem = nil
	}

	if task.CurrentItem == nil {
		if ci := task.ConfigurationItems.Front(); ci == nil {
			task.Complete()
			return
		} else {
			task.CurrentItem = ci
		}
	}

	if time.Now().UTC().Sub(task.CurrentItem.Value.(*ConfigItem).SendtAt).Seconds() < 10 {
		return
	}

	stringParameter := task.CurrentItem.Value.(*ConfigItem).Parameter()
	if err := task.Device.Send(stringParameter); err != nil {
		logger.Error("[ConfigurationTask] Cant send configuration. Error: ", err)
		return
	}
	logger.Info("[ConfigurationTask] Config :", stringParameter, " sent.")
	if task.CurrentItem.Value.(*ConfigItem).State == 0 {
		task.CurrentItem.Value.(*ConfigItem).State = 2
	}
	task.CurrentItem.Value.(*ConfigItem).SendtAt = time.Now().UTC()
}

//DeviceResponce to task
func (task *ConfigurationTask) DeviceResponce(responce interface{} /**message.AckMessage*/) {
	switch responce.(type) {
	case *message.AckMessage:
		{
			task.processAckMessageFromDevice(responce.(*message.AckMessage))
		}
	default:
		{
			return
		}
	}
}

func (task *ConfigurationTask) processAckMessageFromDevice(message *message.AckMessage) {
	task.mutex.Lock()
	defer task.mutex.Unlock()
	if task.CurrentItem == nil {
		task.Execute()
		return
	}
	currentCfgItem := task.CurrentItem.Value.(*ConfigItem)
	if strings.ToUpper(currentCfgItem.Parameter()) == strings.ToUpper(message.Value) {
		task.CurrentItem.Value.(*ConfigItem).State = 3
		if currentCfgItem.Name == "24" || currentCfgItem.Name == "500" {
			task.Device.NewRequiredParameter(currentCfgItem.Name, currentCfgItem.Value)
		}
	}
	task.Execute()
}

//Complete call when task is completed
func (task *ConfigurationTask) Complete() {
	defer func() {
		task.Device = nil
	}()
	task.OnTaskCompleted(task.TaskType)
}
