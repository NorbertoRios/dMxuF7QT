package sensors

import (
	"genx-go/core/columns"
)

//TemperatureValueSensor values from binary report
type TemperatureValueSensor struct {
	Base
	Values map[int]int
}

//BuildTemperatureValueSensor returns temp sensors values
func BuildTemperatureValueSensor(rData map[string]interface{}) ISensor {
	if v, f := rData["TemperatureSensors"]; f {
		column := &columns.Temperature{RawValue: v}
		values := make(map[int]int)
		sNum := 4
		for i := 0; i < 8; i += 2 {
			values[sNum] = column.Value(i)
			sNum--
		}
		//delete(rData, "IBID")
		return &TemperatureValueSensor{Values: values}
	}
	return nil
}

//ToDTO ..
func (s *TemperatureValueSensor) ToDTO() map[string]interface{} {
	// hash := make(map[string]interface{})
	// for i, v := range s.Values {
	// 	hash[fmt.Sprintf("TemperatureValueSensor%v", i)] = v
	// }
	// return hash
	return make(map[string]interface{})
}
