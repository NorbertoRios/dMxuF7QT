package sensors

import (
	"genx-go/core/columns"
	"genx-go/types"
	"time"
)

//Outputs ...
type Outputs struct {
	Base
	Relays map[int]byte
}

//BuildOutputs returns array of switches
func BuildOutputs(data map[string]interface{}) ISensor {
	v, f := data["Relay"]
	if !f {
		return nil
	}
	relays := make(map[int]byte)
	for i := 0; i < 4; i++ {
		bValue := columns.Byte{RawValue: v}
		mask := &types.Byte{Data: bValue.Value()}
		boolState := &types.Bool{Data: mask.BitIsSet(i)}
		relays[i+1] = boolState.ToByte()
	}
	o := &Outputs{Relays: relays}
	o.symbol = "Relay"
	o.createdAt = time.Now().UTC()
	return o
}

//BuildOutputsFromString returns switches from string
func BuildOutputsFromString(bitMask string) ISensor {
	sType := &types.String{Data: bitMask}
	byteValue := sType.BitmaskStringToByte()
	resultRelays := make(map[int]byte)
	for i := 0; i < 4; i++ {
		mask := &types.Byte{Data: byteValue}
		boolState := &types.Bool{Data: mask.BitIsSet(i)}
		resultRelays[i+1] = boolState.ToByte()
	}
	o := &Outputs{Relays: resultRelays}
	o.symbol = "Relay"
	o.createdAt = time.Now().UTC()
	return o
}

//ToDTO ...
func (s *Outputs) ToDTO() map[string]interface{} {
	hash := make(map[string]interface{})
	data := byte(0)
	for i, r := range s.Relays {
		data = s.append(i, data, r)
	}
	hash[s.symbol] = data
	return hash
}

func (Outputs) append(id int, data, value byte) byte {
	return data | value<<id - 1
}
