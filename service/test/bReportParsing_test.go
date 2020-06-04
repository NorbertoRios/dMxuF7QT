package test

import (
	"fmt"
	"genx-go/configuration"
	"genx-go/genxutils"
	"genx-go/core/sensors"
	"genx-go/genxparser"
	"testing"
	"time"
)

// {
//     "DevId": "genx_000036002996",
//     "Data": {
//         "DevId": "genx_000036002996",
//         "GpsValidity": 0,
//         "IgnitionState": 1,
//         "FMI": "Off",
//         "PowerState": "Powered",
//         "PrevSourceId": 90317049,
//         "ReceivedTime": "2020-05-21T09:31:50Z",
//         "storage": "",
//         "Latitude": 48.848717,
//         "Longitude": 37.606094,
//         "Speed": 6.67,
//         "Heading": 313,
//         "TimeStamp": "2020-05-21T09:30:48Z",
//         "Odometer": 324956,
//         "Reason": 6,
//         "LocId": 22470,
//         "Relay": 4,
//         "Supply": 11207,
//         "CSID": 25503,
//         "RSSI": -103,
//         "Satellites": 0,
//         "IBID": 712242433‬,
//         "GPSStat": 56,
//         "GPIO": 0,
//         "Rpm": 0,
//         "BusStatus": "N",
//         "BusOdometer": 0,
//         "DerivedOdometer": 0,
//         "VIN": "WF05XXGCC5GG58765"
//     }
// }

func TestMessageParsing(t *testing.T) {
	param24 := "1.7.13.36.3.4.23.65.10.17.11.79.46.44.43.82.152.41.48.56.70.77.93.130;"
	file := genxutils.ConstructFile("reportConfiguration.xml")
	xmlProvider := configuration.ConstructXmlProvider(file)
	config, err := configuration.ConstructReportConfiguration(xmlProvider)
	if err != nil {
		t.Error("Error while instantation report configuration")
	}
	parser := genxparser.BuildGenxBinaryReportParser(param24, config)
	if parser == nil {
		t.Error("Parser is nil")
	}
	packet := []byte{0x33, 0x36, 0x30, 0x30, 0x32, 0x39, 0x39, 0x36, 0x00, 0x00, 0x00, 0x57, 0xc6, 0x00, 0x18, 0x5e, 0xc6, 0x4a, 0x48, 0x0a, 0x7b, 0x57, 0x16, 0x08, 0x11, 0xc3, 0xac, 0x00, 0x04, 0xf5, 0x5c, 0x00, 0x06, 0x01, 0x39, 0x01, 0x02, 0x00, 0x99, 0x00, 0x00, 0x63, 0x9f, 0x2a, 0x73, 0xf5, 0x01, 0x80, 0x2b, 0xc7, 0x38, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0C, 0x2A, 0x0D, 0x2A, 0x0F, 0x23, 0x0F, 0xE3}
	rm := factory.BuildRawMessage(packet)
	if err != nil {
		t.Error("[TestAckPreParsing] Error in construct new raw message")
	}
	result, ack := parser.Parse(rm)
	if result == nil {
		t.Error("Parsed result is null")
	}
	message := result[0]
	identity := message.Identity
	if ack != "UDPACK 70" {
		t.Error("Failed in ack")
	}
	if identity != "genx_000036002996" {
		t.Error("Failed identity")
	}
	checkBReportSensors(message.Sensors, t)
}

func checkBReportSensors(sensorsArr []sensors.ISensor, t *testing.T) {
	for _, sens := range sensorsArr {
		switch sens.(type) {
		case *sensors.IButtonSensor:
			{
				assert("IBID", sens.(*sensors.IButtonSensor).BtnID, int32(0x2a73f501), t) //Driver Id
				break
			}
		case *sensors.GPSSensor:
			{
				assert("Latitude", sens.(*sensors.GPSSensor).Latitude, float32(48.848717), t)
				assert("Longitude", sens.(*sensors.GPSSensor).Longitude, float32(37.606094), t)
				assert("Speed", sens.(*sensors.GPSSensor).Speed, float32(6.666667), t)
				assert("Heading", sens.(*sensors.GPSSensor).Heading, float32(313), t)
				assert("GPSStatus", sens.(*sensors.GPSSensor).Status, byte(56), t)
				assert("Speed", sens.(*sensors.GPSSensor).Satellites, byte(0), t)
				break
			}
		case *sensors.IgnitionSensor:
			{
				assert("Ignition", sens.(*sensors.IgnitionSensor).IgnitionState, byte(1), t)
				break
			}
		case *sensors.GPIOSensor:
			{
				assert("GPIO", sens.(*sensors.GPIOSensor).Switches, byte(0), t)
				break
			}
		case *sensors.RelaySensor:
			{
				assert("Relay", sens.(*sensors.RelaySensor).Relay, byte(4), t)
				break
			}
		case *sensors.NetworkSensor:
			{
				assert("CSID", sens.(*sensors.NetworkSensor).CSID, int32(25503), t)
				assert("RSSI", sens.(*sensors.NetworkSensor).RSSI, int8(-103), t)
				break
			}
		case *sensors.PowerSensor:
			{
				assert("Supply", sens.(*sensors.PowerSensor).Supply, int32(11207), t)
				break
			}
		case *sensors.QueueSensor:
			{
				assert("LockId", sens.(*sensors.QueueSensor).LockID, uint32(22470), t)
				break
			}
		case *sensors.TimeSensor:
			{
				tShould, _ := time.Parse("2006-01-02T15:04:05Z", "2020-05-21T09:30:48Z")
				assert("TimeStamp", sens.(*sensors.TimeSensor).TimeStamp, tShould, t)
				break
			}
		case *sensors.TripSensor:
			{
				assert("Odometer", sens.(*sensors.TripSensor).Odometer, int32(324956), t)
				break
			}
		case *sensors.TemperatureValueSensor: //0x0C, 0x2A, 0x0D, 0x2A, 0x0F, 0x23, 0x0F, 0xE3 - TValues
			{
				assert("TValue1", sens.(*sensors.TemperatureValueSensor).Values[0], int(0x0C2A), t)
				assert("TValue2", sens.(*sensors.TemperatureValueSensor).Values[1], int(0x0D2A), t)
				assert("TValue3", sens.(*sensors.TemperatureValueSensor).Values[2], int(0x0F23), t)
				assert("TValue4", sens.(*sensors.TemperatureValueSensor).Values[3], int(0x0FE3), t)
				break
			}
		default:
			{
				t.Error("Unexpected sensor")
			}
		}
	}
}

func assert(paramName string, value interface{}, should interface{}, t *testing.T) {
	if value != should {
		t.Error(fmt.Sprintf("%v=%v doesnt equal %v", paramName, value, should))
	}
}
