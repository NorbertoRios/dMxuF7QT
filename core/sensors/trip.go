package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
)

//TripSensor represents trip data
type TripSensor struct {
	Base
	Odometer int32
}

//BuildTripSensor returns new gps sensor
func BuildTripSensor(data map[string]interface{}) ISensor {
	if v, f := data[core.OdometerKm]; f {
		odometer := &columns.OdometerKm{RawValue: v}
		return &TripSensor{Odometer: odometer.Value()}
	}
	return nil
}
