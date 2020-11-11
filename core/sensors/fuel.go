package sensors

import "genx-go/types"

//Fuel fuel sensor
type Fuel struct {
	Base
	FuelLevel float32
}

//BuildFuelSensor returns new gps sensor
func BuildFuelSensor(data map[string]interface{}) ISensor {
	return nil
}

//BuildFuelSensorFromString returns new fuel sensor
func BuildFuelSensorFromString(fuelLevel string) ISensor {
	fLevel := types.String{Data: fuelLevel}
	return &Fuel{FuelLevel: fLevel.Float32()}
}
