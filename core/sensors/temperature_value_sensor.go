package sensors

import (
	"genx-go/core"
	"genx-go/core/genxcolumns"
)

//TemperatureValueSensor values from binary report
type TemperatureValueSensor struct {
	Values []int
}

//BuildTemperatureValueSensor returns temp sensors values
func BuildTemperatureValueSensor(rData map[string]interface{}) ISensor {
	if v, f := rData[core.TemperatureSensors]; f {
		column := &genxcolumns.TemperatureColumn{RawValue: v}
		values := make([]int, 0)
		for i := 0; i < 8; i += 2 {
			values = append(values, column.Value(i))
		}
		return &TemperatureValueSensor{Values: values}
	}
	return nil
}
