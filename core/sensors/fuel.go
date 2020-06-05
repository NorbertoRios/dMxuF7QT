package sensors

import "genx-go/utils"

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
	fLevel := utils.StringUtils{Data: fuelLevel}
	return &Fuel{FuelLevel: fLevel.Float32()}
}
