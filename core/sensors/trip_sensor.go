package sensors

import (
	"genx-go/core"
	"genx-go/core/genxcolumns"
)

//TripSensor represents trip data
type TripSensor struct {
	BaseSensor
	Odometer int32
}

//BuildTripSensor returns new gps sensor
func BuildTripSensor(data map[string]interface{}) ISensor {
	if v, f := data[core.OdometerKm]; f {
		odometer := &genxcolumns.OdometerKmColumn{RawValue: v}
		return &TripSensor{Odometer: odometer.Value()}
	}
	return nil
}
