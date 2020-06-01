package genxmessage

import "genx-go/core/sensors"

//BuildBinaryMessageSensors returns binary message sensors
func BuildBinaryMessageSensors() *BinaryMessageSensors {
	return &BinaryMessageSensors{
		SensorBuilders: []func(map[string]interface{}) sensors.ISensor{
			sensors.BuildTemperatureValueSensor,
			sensors.BuildDriverIDSensor,
			sensors.BuildGPIOSensor,
			sensors.BuildGpsSensor,
			sensors.BuildIgnitionSensor,
			sensors.BuildNetworkSensor,
			sensors.BuildPowerSensor,
			sensors.BuildQueueSensor,
			sensors.BuildRelaySensor,
			sensors.BuildTimeSensor,
			sensors.BuildTripSensor,
		},
	}
}

//BinaryMessageSensors represents location binary message sensors
type BinaryMessageSensors struct {
	SensorBuilders []func(map[string]interface{}) sensors.ISensor
}

//Build build sensors for message
func (bms *BinaryMessageSensors) Build(rData map[string]interface{}) []sensors.ISensor {
	returnedValue := make([]sensors.ISensor, 0)
	for _, builder := range bms.SensorBuilders {
		sensor := builder(rData)
		if sensor != nil {
			returnedValue = append(returnedValue, sensor)
		}
	}
	return returnedValue
}
