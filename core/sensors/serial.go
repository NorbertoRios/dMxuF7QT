package sensors

import "fmt"

//SerialSensor 1
type SerialSensor struct {
	RawValue interface{}
}

//Value returns speed value
func (column *SerialSensor) Value() string {
	return column.RawValue.(string)
}

//ToIdentity returns identity value
func (column *SerialSensor) ToIdentity() string {
	serial := column.RawValue.(string)
	for l := len(column.RawValue.(string)); l < 12; l++ {
		serial = fmt.Sprintf("0%v", serial)
	}
	return fmt.Sprintf("genx_%v", serial)
}
