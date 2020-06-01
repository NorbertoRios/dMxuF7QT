package sensors

import "genx-go/genxutils"

//FuelSensor fuel sensor
type FuelSensor struct {
	BaseSensor
	FuelLevel float32
}

//BuildFuelSensor returns new gps sensor
func BuildFuelSensor(data map[string]interface{}) ISensor {
	return nil
}

//BuildFuelSensorFromString returns new fuel sensor
func BuildFuelSensorFromString(fuelLevel string) ISensor {
	fLevel := genxutils.StringUtils{Data: fuelLevel}
	return &FuelSensor{FuelLevel: fLevel.Float32()}
}
