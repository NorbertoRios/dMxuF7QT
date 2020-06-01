package sensors

import (
	"genx-go/core"
	"genx-go/core/genxcolumns"
	"genx-go/genxutils"
)

//RelaySensor represents relay sensor
type RelaySensor struct {
	Relay byte
}

//BuildRelaySensor returns relay sensor
func BuildRelaySensor(data map[string]interface{}) ISensor {
	if v, f := data[core.Relay]; f {
		relays := &genxcolumns.RelayColumn{RawValue: v}
		return &RelaySensor{Relay: relays.Value()}
	}
	return nil
}

//BuildRelaySensorFromString returns switches from string
func BuildRelaySensorFromString(value string) ISensor {
	sU := &genxutils.StringUtils{Data: value}
	relays := &genxcolumns.RelayColumn{RawValue: sU.Byte()}
	return &RelaySensor{
		Relay: relays.Value(),
	}
}
