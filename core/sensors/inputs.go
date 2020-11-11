package sensors

import (
	"genx-go/core"
	"genx-go/types"
)

//BuildInputs returns array of switches
func BuildInputs(data map[string]interface{}) []ISensor {
	v, f := data[core.Switches]
	if !f {
		return nil
	}
	swOnCode := byte(23)
	swOffCode := byte(27)
	resultSwitches := make([]ISensor, 0)
	for i := 0; i < 4; i++ {
		posibleReasons := map[byte]byte{
			swOnCode:  1, //switch on
			swOffCode: 2, //switch off
		}
		sw := BuildSwitch(i, v)
		sw.Trigered = Trigered(data, posibleReasons)
		resultSwitches = append(resultSwitches, sw)
		swOffCode--
		swOnCode--
	}
	return resultSwitches
}

//BuildInputsFromString returns switches from string
func BuildInputsFromString(bitMask string) []ISensor {
	sType := &types.String{Data: bitMask}
	byteValue := sType.BitmaskStringToByte()
	resultSwitches := make([]ISensor, 0)
	for i := 0; i < 4; i++ {
		sw := BuildSwitchFromByte(i, byteValue)
		resultSwitches = append(resultSwitches, sw)
	}
	return resultSwitches
}
