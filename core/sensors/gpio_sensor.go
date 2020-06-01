package sensors

import (
	"genx-go/core"
	"genx-go/core/genxcolumns"
	"genx-go/genxutils"
)

//GPIOSensor represents gpio sensor
type GPIOSensor struct {
	Switches byte
}

//BuildGPIOSensor build gpio sensor
func BuildGPIOSensor(data map[string]interface{}) ISensor {
	if v, f := data[core.Switches]; f {
		switches := genxcolumns.ByteColumn{RawValue: v}
		return &GPIOSensor{Switches: switches.Value()}
	}
	return nil
}

//BuildGPIOSensorFromString returns switches from string
func BuildGPIOSensorFromString(value string) ISensor {
	sU := &genxutils.StringUtils{Data: value}
	return &GPIOSensor{
		Switches: sU.Byte(),
	}
}
