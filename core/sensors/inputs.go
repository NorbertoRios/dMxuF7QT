package sensors

import (
	"genx-go/core/columns"
	"genx-go/types"
	"time"
)

//Inputs ...
type Inputs struct {
	Base
	Switches map[int]byte
}

//BuildInputs ...
func BuildInputs(data map[string]interface{}) ISensor {
	inputs := newInputs(data)
	swOnCode := byte(20)
	swOffCode := byte(24)
	posibleReasons := map[byte]byte{
		swOnCode:  1, //switch on
		swOffCode: 2, //switch off
	}
	for i := 0; i < 4; i++ {
		inputs.Trigered(data, posibleReasons)
		swOnCode++
		swOffCode++
	}
	return inputs
}

func newInputs(data map[string]interface{}) *Inputs {
	v, f := data["Switches"]
	if !f {
		return nil
	}
	switches := make(map[int]byte)
	for i := 0; i < 4; i++ {
		bValue := columns.Byte{RawValue: v}
		switchMask := &types.Byte{Data: bValue.Value()}
		boolState := &types.Bool{Data: switchMask.BitIsSet(i)}
		switches[i+1] = boolState.ToByte()
	}
	s := &Inputs{}
	s.symbol = "GPIO"
	s.createdAt = time.Now().UTC()
	s.Switches = switches
	return s
}

//BuildInputsFromString returns switches from string
func BuildInputsFromString(bitMask string) ISensor {
	sType := &types.String{Data: bitMask}
	byteValue := sType.Byte(10)
	resultSwitches := make(map[int]byte)
	for i := 0; i < 4; i++ {
		switchMask := &types.Byte{Data: byteValue}
		boolState := &types.Bool{Data: switchMask.BitIsSet(i)}
		resultSwitches[i+1] = boolState.ToByte()
	}
	s := &Inputs{Switches: resultSwitches}
	s.symbol = "GPIO"
	s.createdAt = time.Now().UTC()
	return s
}

//ToDTO ...
func (s *Inputs) ToDTO() map[string]interface{} {
	hash := make(map[string]interface{})
	data := byte(0)
	for i, r := range s.Switches {
		data = s.append(i, data, r)
	}
	hash[s.symbol] = data
	return hash
}

func (Inputs) append(id int, data, value byte) byte {
	return data | value<<id - 1
}
