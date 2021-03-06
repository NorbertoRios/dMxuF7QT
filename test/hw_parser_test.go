package test

import (
	"genx-go/core/sensors"
	"genx-go/message"
	"genx-go/parser"
	"testing"
)

// "DevId": "genx_000003912835",
// "Message": "MODEL:GNX-5P
// SN:000003912835
// FW:G699.06.78kX 12:59:45 May 25 2012
// HW:656, HWOPTID:0016
// IMEI:357852038210210
// MVER:07.60.00
// GVER:7.03 (45969) 00040007
// On:40:09:45(48)
// Ign-ON,Volt-12131,Switch-0000,Relay-0000,A2D-4150
// V-12131/12133/12133 Temp-319
// Mallocs 0
// CRC:6a10.6ce4.88e8.88e8.60c0.948.537a.537a.9bfd.9bfd.

// 000003912835 3912835
// "

func TestHWMessageParsing(t *testing.T) {
	packet := []byte("MODEL:GNX-5P\nSN:000003912835\nFW:G699.06.78kX 12:59:45 May 25 2012\nHW:656, HWOPTID:0016\nIMEI:357852038210210\nMVER:07.60.00\nGVER:7.03 (45969) 00040007\nOn:40:09:45(48)\nIgn-ON,Volt-12131,Switch-1001,Relay-0110,A2D-4150\nV-12131/12133/12133 Temp-319\nMallocs 0\nCRC:6a10.6ce4.88e8.88e8.60c0.948.537a.537a.9bfd.9bfd.\n000003912835 3912835")
	rm := factory.BuildRawMessage(packet)
	parser := parser.BuildGenxHardwareMessageParser()
	msg := parser.Parse(rm)
	if msg.(*message.HardwareMessage).Firmware != "G699.06.78kX" {
		t.Error("Wrong firmware version")
	}
	checkHWSensors(msg.(*message.HardwareMessage).Sensors, t)
}

func checkHWSensors(sensorsArr []sensors.ISensor, t *testing.T) {
	for _, sens := range sensorsArr {
		switch sens.(type) {
		case *sensors.IgnitionSensor:
			{
				assert("Ignition", sens.(*sensors.IgnitionSensor).IgnitionState, byte(1), t)
				break
			}
		case *sensors.Inputs:
			{
				shouldState := []byte{1, 0, 0, 1}
				inputs := sens.(*sensors.Inputs)
				for i := 0; i < 3; i++ {
					if inputs.Switches[i+1] != shouldState[i] {
						t.Error("Inputs states dont equal. Input id:", i)
					}
				}
			}
		case *sensors.Outputs:
			{
				shouldState := []byte{0, 1, 1, 0}
				relays := sens.(*sensors.Outputs)
				for i := 0; i < 3; i++ {
					if relays.Relays[i+1] != shouldState[i] {
						t.Error("Relays states dont equal. Input id:", i)
					}
				}
			}
		case *sensors.PowerSensor:
			{
				assert("Supply", sens.(*sensors.PowerSensor).Supply, int32(12131), t)
				break
			}
		default:
			{
				t.Error("Unexpected sensor")
				break
			}
		}
	}
}
