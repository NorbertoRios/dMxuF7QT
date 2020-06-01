package sensors

import (
	"genx-go/core"
	"genx-go/core/genxcolumns"
)

//IgnitionSensor represents inputs
type IgnitionSensor struct {
	BaseSensor
	IgnitionState byte
}

//BuildIgnitionSensor returns new gps sensor
func BuildIgnitionSensor(data map[string]interface{}) ISensor {
	if v, f := data[core.Ignition]; f {
		ignitionState := &genxcolumns.ByteColumn{RawValue: v}
		return &IgnitionSensor{
			IgnitionState: ignitionState.Value(),
		}
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
