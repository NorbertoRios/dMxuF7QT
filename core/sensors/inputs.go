package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
	"genx-go/utils"
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
	bValue := columns.Byte{RawValue: v}
	utilsSwitchMask := &utils.ByteUtility{Data: bValue.Value()}
	for i := 0; i < 4; i++ {
		boolState := &utils.BoolUtils{Data: utilsSwitchMask.BitIsSet(i)}
		sw := &Switch{ID: i, State: boolState.ToByte()}
		posibleReasons := map[byte]bool{
			swOnCode:  true,
			swOffCode: true,
		}
		sw.Trigered = Trigered(data, posibleReasons)
		resultSwitches = append(resultSwitches, sw)
		swOnCode--
		swOffCode--
	}
	return resultSwitches
}
