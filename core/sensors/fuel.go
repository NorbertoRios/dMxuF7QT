package sensors

import (
	"genx-go/core"
	"genx-go/types"
	"time"
)

//Fuel fuel sensor
type Fuel struct {
	Base
	FuelLevel float32
}

//BuildFuelSensor returns new gps sensor
func BuildFuelSensor(data map[string]interface{}) ISensor {
	sensor := &Fuel{}
	sensor.symbol = "FuelLevel"
	return sensor
}

//BuildFuelSensorFromString returns new fuel sensor
func BuildFuelSensorFromString(fuelLevel string) ISensor {
	fLevel := types.String{Data: fuelLevel}
	f := &Fuel{FuelLevel: fLevel.Float32()}
	f.symbol = "FuelLevel"
	f.createdAt = time.Now().UTC()
	return f
}

//ToDTO ..
func (s *Fuel) ToDTO() map[string]interface{} {
	hash := make(map[string]interface{})
	hash[core.FuelLevel] = s.FuelLevel
	return hash
}
