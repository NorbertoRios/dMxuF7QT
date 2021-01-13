package adaptors

import (
	"encoding/json"
	"genx-go/core/sensors"
	"genx-go/logger"
	"genx-go/repository/models"
)

//NewDeviceActivity ...
func NewDeviceActivity(_activity *models.DeviceActivity) *DeviceActivity {
	return &DeviceActivity{
		activity: _activity,
	}
}

//DeviceActivity ...
type DeviceActivity struct {
	activity *models.DeviceActivity
}

//Adapt ...
func (da *DeviceActivity) Adapt() []sensors.ISensor {
	activitySensors := []sensors.ISensor{}
	message := &DtoMessage{}
	err := json.Unmarshal([]byte(str), message)
	if err != nil {
		logger.Logger().WriteToLog(logger.Error, "[DeviceActivity | Adapt] Error while unmarshaling last known device state. Error: ", err, ". LastActivity : ", da.activity)
		return activitySensors
	}
}

func (da *DeviceActivity) adaptLastMessageToSensors(dto *DtoMessage) []sensors.ISensor {
	
}

func (da *DeviceActivity) adaptSoftware() ISensor {

}
