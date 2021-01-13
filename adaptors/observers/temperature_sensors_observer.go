package observers

import (
	"genx-go/adaptors/dto"
	"genx-go/core/sensors"
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
	temperatureSensors := &sensors.TemperatureSensors{}
	for index, tempSens := range _message.TemperatureSensors.Sensors {
		temperatureSensors.Sensors[index] = &sensors.TemperatureSensor{
			ID:tempSens.ID,
			Imei: tempSens.
		}
	}
}
