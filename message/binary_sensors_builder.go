package message

import "genx-go/core/sensors"

//BuildBinaryMessageSensors returns binary message sensors
func BuildBinaryMessageSensors() *BinaryMessageSensors {
	return &BinaryMessageSensors{
		SingleSensorBuilders: []func(map[string]interface{}) sensors.ISensor{
			sensors.BuildTemperatureValueSensor,
			sensors.BuildIButtonSensor,
			sensors.BuildGpsSensor,
			sensors.BuildFuelSensor,
			sensors.BuildIgnitionSensor,
			sensors.BuildNetworkSensor,
			sensors.BuildPowerSensor,
			sensors.BuildQueueSensor,
			sensors.BuildTimeSensor,
			sensors.BuildTripSensor,
			sensors.BuildInputs,
			sensors.BuildOutputs,
		},
	}
}

//BinaryMessageSensors represents location binary message sensors
type BinaryMessageSensors struct {
	SingleSensorBuilders []func(map[string]interface{}) sensors.ISensor
}

//Build build sensors for message
func (bms *BinaryMessageSensors) Build(rData map[string]interface{}) []sensors.ISensor {
	returnedValue := make([]sensors.ISensor, 0)
	for _, builder := range bms.SingleSensorBuilders {
		sensor := builder(rData)
		if sensor != nil {
			returnedValue = append(returnedValue, sensor)
		}
	}
	return returnedValue
}
