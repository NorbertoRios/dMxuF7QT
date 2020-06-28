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
func BuildConfigurationTask(storage *TaskStorage, config *models.ConfigurationModel, onTaskCompleted func(string)) {
	if config.ID == 0 || len(config.Command) == 0 {
		logger.Error("[BuildConfigurationTask] Error in configuration.")
		return
	}
	task := &ConfigurationTask{
		Storage:            storage,
		mutex:              &sync.Mutex{},
		ConfigurationItems: devideConfiguration(config.Command),
		OnTaskCompleted:    onTaskCompleted,
	}
	storage.NewTask(task.TaskType, task)
	go task.Execute()
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
	return ""
}

//ConfigurationTask represents task for send config to device
type ConfigurationTask struct {
	TaskType           string
	mutex              *sync.Mutex
	Storage            *TaskStorage
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
	if err := task.Storage.Device.Send(stringParameter); err != nil {
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
	if strings.ToUpper(task.CurrentItem.Value.(*ConfigItem).Parameter()) == strings.ToUpper(message.Value) {
		task.CurrentItem.Value.(*ConfigItem).State = 3
	}
	task.Execute()
}

//Complete call when task is completed
func (task *ConfigurationTask) Complete() {
	defer func() {
		task.Storage = nil
	}()
	task.OnTaskCompleted(task.TaskType)
}
