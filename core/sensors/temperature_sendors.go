package sensors

import (
	"encoding/json"
	"fmt"
	"genx-go/logger"
	"time"
)

//NewTemperatureSensors ...
func NewTemperatureSensors(tSensors []*TemperatureSensor) *TemperatureSensors {
	sensorsMap := make(map[string]*TemperatureSensor)
	for _, s := range tSensors {
		sensorsMap[fmt.Sprintf("Sensor%v", s.ID+1)] = s
	}
	sensor := &TemperatureSensors{}
	sensor.symbol = "ts"
	sensor.createdAt = time.Now().UTC()
	return sensor
}

//TemperatureSensors ...
type TemperatureSensors struct {
	Base
	Sensors map[string]*TemperatureSensor `json:",omitempty"`
}

//ToDTO ...
func (s *TemperatureSensors) ToDTO() map[string]interface{} {
	hash := make(map[string]interface{})
	jSens, jErr := json.Marshal(s.Sensors)
	if jErr != nil {
		logger.Logger().WriteToLog(logger.Error, "[TemperatureSensor | ToDTO] Error while marshaling. ", jErr)
		return hash
	}
	hash["ts"] = string(jSens)
	return hash
}
