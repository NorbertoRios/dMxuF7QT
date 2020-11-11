package sensors

import (
	"genx-go/core/columns"
	"genx-go/types"
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
	switchMask := &types.Byte{Data: bValue.Value()}
	boolState := &types.Bool{Data: switchMask.BitIsSet(index)}
	return &Switch{ID: index, State: boolState.ToByte()}
}

//BuildSwitchFromByte return new switch from byte
func BuildSwitchFromByte(index int, mask byte) *Switch {
	switchMask := &types.Byte{Data: mask}
	boolState := &types.Bool{Data: switchMask.BitIsSet(index)}
	return &Switch{ID: index, State: boolState.ToByte()}
}
