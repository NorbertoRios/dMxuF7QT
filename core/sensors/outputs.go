package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
	"genx-go/utils"
)

//BuildOutputs returns relays
func BuildOutputs(data map[string]interface{}) []ISensor {
	v, f := data[core.Relay]
	if !f {
		return nil
	}
	resultRelays := make([]ISensor, 0)
	bValues := columns.Byte{RawValue: v}
	utilsRelayMask := &utils.ByteUtility{Data: bValues.Value()}
	for i := 0; i < 4; i++ {
		boolState := &utils.BoolUtils{Data: utilsRelayMask.BitIsSet(i)}
		relay := &Relay{ID: i, State: boolState.ToByte()}
		resultRelays = append(resultRelays, relay)
	}
	return resultRelays
}
