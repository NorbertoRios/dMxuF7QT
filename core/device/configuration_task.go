package device

import (
	"container/list"
	"genx-go/message"
	"genx-go/repository/models"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"
)

//BuildConfigurationTask create configuration task
func BuildConfigurationTask(device IDevice, config *models.ConfigurationModel) *ConfigurationTask {
	if config.ID == 0 || len(config.Command) == 0 {
		return nil
	}
	return &ConfigurationTask{
		Device:             device,
		mutex:              &sync.Mutex{},
		ConfigurationItems: devideConfiguration(config.Command),
	}
}

func devideConfiguration(config string) *list.List {
	items := list.New()
	re := regexp.MustCompile(`(\n)|(\r\n)`)
	configurations := re.Split(config, -1)
	for _, cfg := range configurations {
		if len(cfg) == 0 ||
			strings.Contains(strings.ToUpper(cfg), "SETPARAM") ||
			strings.Contains(strings.ToUpper(cfg), "ENDPARAM") ||
			strings.Contains(strings.ToUpper(cfg), "BACKUPNVRAM") {
			continue
		}
		if strings.Contains(strings.ToUpper(cfg), "SETBOUNDARY") {
			cfgItem := &ConfigItem{
				Name:  "SETBOUNDARY",
				Value: cfg,
				State: byte(0),
			}
			items.PushBack(cfgItem)
			continue
		}
		cfgName := strings.Split(cfg, "=")[0]
		if len(cfgName) == 0 {
			continue
		}
		cfgItem := &ConfigItem{
			Name:  cfgName,
			Value: cfg,
			State: byte(0),
		}
		items.PushBack(cfgItem)
	}
	return items
}

//ConfigurationTask represents task for send config to device
type ConfigurationTask struct {
	mutex              *sync.Mutex
	Device             IDevice
	CurrentItem        *list.Element
	ConfigurationItems *list.List
	SendAt             time.Time
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
		log.Println("[ConfigurationTask] Cant send configuration. Error: ", err)
		return
	}
	log.Println("[ConfigurationTask] Config :", stringParameter, " sent.")
	if task.CurrentItem.Value.(*ConfigItem).State == 0 {
		task.CurrentItem.Value.(*ConfigItem).State = 2
	}
	task.CurrentItem.Value.(*ConfigItem).SendtAt = time.Now().UTC()
}

//AckArrived to task
func (task *ConfigurationTask) AckArrived(message *message.AckMessage) {
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
		task.Device = nil
	}()
	task.Device.OnConfigTaskCompleted()
}
