package sensors

import (
	"genx-go/core/columns"
	"genx-go/types"
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
	switchMask := &types.Byte{Data: bValue.Value()}
	boolState := &types.Bool{Data: switchMask.BitIsSet(index)}
	return &Relay{ID: index, State: boolState.ToByte()}
}

//BuildRelayFromByte return new relay from byte
func BuildRelayFromByte(index int, mask byte) *Relay {
	switchMask := &types.Byte{Data: mask}
	boolState := &types.Bool{Data: switchMask.BitIsSet(index)}
	return &Relay{ID: index, State: boolState.ToByte()}
}
