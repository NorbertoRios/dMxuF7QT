package sensors

import (
	"genx-go/core/columns"
	"genx-go/utils"
)

//Switch switch
type Switch struct {
	Base
	ID    int
	State byte
}

//BuildSwitch return new switch
func BuildSwitch(index int, swMask interface{}) *Switch {
	bValue := columns.Byte{RawValue: swMask}
	utilsSwitchMask := &utils.ByteUtility{Data: bValue.Value()}
	boolState := &utils.BoolUtils{Data: utilsSwitchMask.BitIsSet(index)}
	return &Switch{ID: index, State: boolState.ToByte()}
}

//BuildSwitchFromByte return new switch from byte
func BuildSwitchFromByte(index int, mask byte) *Switch {
	utilsSwitchMask := &utils.ByteUtility{Data: mask}
	boolState := &utils.BoolUtils{Data: utilsSwitchMask.BitIsSet(index)}
	return &Switch{ID: index, State: boolState.ToByte()}
}
