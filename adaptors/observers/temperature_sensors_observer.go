package observers

import (
	"genx-go/adaptors/dto"
	"genx-go/core/sensors"
	"genx-go/types"
	"strings"
)

//NewTemperatureSensorsObserver ...
func NewTemperatureSensorsObserver() *TemperatureSensorsObserver {
	return &TemperatureSensorsObserver{}
}

//TemperatureSensorsObserver ...
type TemperatureSensorsObserver struct {
}

//Notify ...
func (o *TemperatureSensorsObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	_tSensors := []*sensors.TemperatureSensor{}
	for index, tempSens := range _message.Sensors {
		strIndex := &types.String{Data: strings.ReplaceAll(index, "Sensor", "")}
		id := strIndex.Byte(8)
		imei := tempSens.ID
		value := tempSens.Value
		_tSensors = append(_tSensors, sensors.BuildTemperatureSensor(id, imei, value))
	}
	if len(_tSensors) == 0 {
		return nil
	}
	return sensors.NewTemperatureSensors(_tSensors)
}
