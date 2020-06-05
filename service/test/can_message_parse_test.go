package test

import (
	"genx-go/core/sensors"
	"genx-go/parser"
	"testing"
)

// J1939:250kHz
// CE,0,0,0,0,0,2,0,0,2,0
// 60415,60671,61183,61440,61441,61443,61444,65170,
// Rev 1 Mar.17 17 41110183
// 65217,65226,65244,65248,65253,65257,65261,65262,
// 65263,65264,65265,65266,65269,65270,65271,65272,
// 65274,65276,65279
// Age:0 Spd:53.7 FL:97.2 RPM:1326.5/1326/550 Odo:175093.1
// CT:183.2 TH:3199.4 VIN:2C4RDGBGXDR742233
// J1708 ENABLED(NO DATA)

// CE,0,0,0,134,3,0,0,0,0,136
// OBD(1/1/FD) BFBEB993-8007E019-FED00001 VIN:2C4RDGBGXDR742233
// MIL:4 RPM:754(151936990/1062/31003) Spd:0 Odo:125147575/0
// DSCC:1324/125143
// MAP:35(34) IAT:318 O2:156(244) FL:62 CT:88
// OBDDTC:P0128,P0673,P0676,P0675,OBDPDTC:P0128,P0673,P0676,P0675,

func TestDiagCanResponce(t *testing.T) {
	packet := []byte("J1939:250kHz\nCE,0,0,0,0,0,2,0,0,2,0\n60415,60671,61183,61440,61441,61443,61444,65170,\nRev 1 Mar.17 17 41110183\n65217,65226,65244,65248,65253,65257,65261,65262,\n65263,65264,65265,65266,65269,65270,65271,65272,\n65274,65276,65279\nAge:0 Spd:53.7 FL:97.2 RPM:1326.5/1326/550 Odo:175093.1\nCT:183.2 TH:3199.4 VIN:2C4RDGBGXDR742233\nJ1708 ENABLED(NO DATA)")
	rm := factory.BuildRawMessage(packet)
	parser := parser.BuildCANMessageParser()
	message := parser.Parse(rm)
	if len(message.Sensors) != 2 {
		t.Error("Sensors count error")
	}
	checkCanSensors(message.Sensors, t)
}

func TestDiagCanResponceWithDTCCodes(t *testing.T) {
	packet := []byte("CE,0,0,0,134,3,0,0,0,0,136\nOBD(1/1/FD) BFBEB993-8007E019-FED00001 VIN:2C4RDGBGXDR742233\nMIL:4 RPM:754(151936990/1062/31003) Spd:0 Odo:125147575/0\nDSCC:1324/125143\nMAP:35(34) IAT:318 O2:156(244) FL:97.2 CT:88\nOBDDTC:P0128,P0673,P0676,P0675,OBDPDTC:P0128,P0673,P0676,P0675,")
	rm := factory.BuildRawMessage(packet)
	parser := parser.BuildCANMessageParser()
	message := parser.Parse(rm)
	if len(message.Sensors) != 3 {
		t.Error("Sensors count error")
	}
	checkCanSensors(message.Sensors, t)
}

func checkCanSensors(sensorsArr []sensors.ISensor, t *testing.T) {
	for _, sens := range sensorsArr {
		switch sens.(type) {
		case *sensors.VINSensor:
			{
				assert("VIN", sens.(*sensors.VINSensor).VIN, "2C4RDGBGXDR742233", t) //VIN
				break
			}
		case *sensors.Fuel:
			{
				assert("Fuel level", sens.(*sensors.Fuel).FuelLevel, float32(97.2), t)
				break
			}
		case *sensors.DTCCodes:
			{
				assertStringSlice("DTCCodes", sens.(*sensors.DTCCodes).Codes, []string{"P0128", "P0673", "P0676", "P0675"}, t)
				break
			}
		default:
			{
				t.Error("Unexpected sensor")
			}
		}
	}
}

func assertStringSlice(name string, value []string, should []string, t *testing.T) {
	if !testSliceEq(value, should) {
		t.Error("[", name, "]Invalid value. Should:", should, "; Current:", value)
	}
}

func testSliceEq(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
