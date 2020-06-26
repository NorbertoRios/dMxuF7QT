package sensors

import (
	"genx-go/utils"
)

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
	uIdentity := &utils.StringUtils{Data: column.RawValue.(string)}
	return uIdentity.Identity()
}
