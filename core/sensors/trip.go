package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
	"time"
)

//TripSensor represents trip data
type TripSensor struct {
	Base
	Odometer int32
}

//BuildAdaptedTripSensor ...
func BuildAdaptedTripSensor(_odometer int32) ISensor {
	sensor := &TripSensor{Odometer: _odometer}
	sensor.symbol = "Odometer"
	sensor.createdAt = time.Now().UTC()
	return sensor
}

//BuildTripSensor returns new gps sensor
func BuildTripSensor(data map[string]interface{}) ISensor {
	if v, f := data[core.OdometerKm]; f {
		odometer := &columns.OdometerKm{RawValue: v}
		sensor := &TripSensor{Odometer: odometer.Value()}
		sensor.symbol = "Odometer"
	}
	return nil
}

//ToDTO ...
func (s *TripSensor) ToDTO() map[string]interface{} {
	hash := make(map[string]interface{})
	hash[s.symbol] = s.Odometer
	return hash
}
