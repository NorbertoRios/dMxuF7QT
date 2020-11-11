package sensors

import (
	"genx-go/core"
	"genx-go/types"
)

//BuildOutputs returns array of switches
func BuildOutputs(data map[string]interface{}) []ISensor {
	v, f := data[core.Relay]
	if !f {
		return nil
	}
	resultRelays := make([]ISensor, 0)
	for i := 0; i < 4; i++ {
		sw := BuildRelay(i, v)
		resultRelays = append(resultRelays, sw)
	}
	return resultRelays
}

//BuildOutputsFromString returns switches from string
func BuildOutputsFromString(bitMask string) []ISensor {
	sType := &types.String{Data: bitMask}
	byteValue := sType.BitmaskStringToByte()
	resultSwitches := make([]ISensor, 0)
	for i := 0; i < 4; i++ {
		sw := BuildRelayFromByte(i, byteValue)
		resultSwitches = append(resultSwitches, sw)
	}
	return resultSwitches
}
