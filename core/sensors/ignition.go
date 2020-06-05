package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
)

//IgnitionSensor represents inputs
type IgnitionSensor struct {
	Base
	IgnitionState byte
}

//BuildIgnitionSensor returns new gps sensor
func BuildIgnitionSensor(data map[string]interface{}) ISensor {
	posibleReasons := map[byte]bool{
		byte(3): true,
		byte(4): true,
	}
	if v, f := data[core.Ignition]; f {
		ignitionState := &columns.Byte{RawValue: v}
		sensor := &IgnitionSensor{IgnitionState: ignitionState.Value()}
		sensor.Trigered = Trigered(data, posibleReasons)
		return sensor
	}
	return nil
}

//BuildIgnitionSensorFromString returns new ignition sensor
func BuildIgnitionSensorFromString(data string) ISensor {
	var state byte
	switch data {
	case "ON":
		state = 1
		break
	default:
		state = 0
		break
	}
	return &IgnitionSensor{IgnitionState: state}
}
