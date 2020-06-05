package sensors

import (
	"genx-go/core/columns"
	"genx-go/utils"
)

//Relay relay
type Relay struct {
	Base
	ID    int
	State byte
}

//BuildRelay return new switch
func BuildRelay(index int, swMask interface{}) *Relay {
	bValue := columns.Byte{RawValue: swMask}
	utilsSwitchMask := &utils.ByteUtility{Data: bValue.Value()}
	boolState := &utils.BoolUtils{Data: utilsSwitchMask.BitIsSet(index)}
	return &Relay{ID: index, State: boolState.ToByte()}
}

//BuildRelayFromByte return new relay from byte
func BuildRelayFromByte(index int, mask byte) *Relay {
	utilsSwitchMask := &utils.ByteUtility{Data: mask}
	boolState := &utils.BoolUtils{Data: utilsSwitchMask.BitIsSet(index)}
	return &Relay{ID: index, State: boolState.ToByte()}
}
