package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
	"time"
)

//IgnitionSensor represents inputs
type IgnitionSensor struct {
	Base
	IgnitionState byte
}

//BuildIgnitionSensor returns new gps sensor
func BuildIgnitionSensor(data map[string]interface{}) ISensor {
	posibleReasons := map[byte]byte{
		3: 1, // 1- means IgnitionOn
		2: 2, // 2 - mean IgnitionOff
	}
	if v, f := data[core.Ignition]; f {
		ignitionState := &columns.Byte{RawValue: v}
		sensor := &IgnitionSensor{IgnitionState: ignitionState.Value()}
		sensor.Trigered(data, posibleReasons)
		sensor.symbol = "Ignition"
		sensor.createdAt = time.Now().UTC()
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
	sensor := &IgnitionSensor{IgnitionState: state}
	sensor.symbol = "Ignition"
	sensor.createdAt = time.Now().UTC()
	return sensor

}

//ToDTO ..
func (s *IgnitionSensor) ToDTO() map[string]interface{} {
	hash := make(map[string]interface{})
	hash[s.symbol] = s.IgnitionState
	return hash
}
