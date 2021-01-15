package adaptors

import (
	"fmt"
	"genx-go/core/sensors"
	"genx-go/types"
	"strings"
	"testing"
)

func checkSensorsState(_sensors map[string]sensors.ISensor, _template *Template, t *testing.T) map[string]sensors.ISensor {
	for index, sensor := range _sensors {
		switch index {
		case "GPS":
			{
				template := _template.Template("GPS").(*GPSTemplate)
				gps := sensor.(*sensors.GPSSensor)
				compareValues("GPSSensor", "Latitude", template.Latitude, gps.Latitude, t)
				compareValues("GPSSensor", "Longitude", template.Longitude, gps.Longitude, t)
				compareValues("GPSSensor", "Speed", template.Speed, gps.Speed, t)
				compareValues("GPSSensor", "Heading", template.Heading, gps.Heading, t)
				compareValues("GPSSensor", "Satellites", template.Satellites, gps.Satellites, t)
				compareValues("GPSSensor", "GpsValidity", template.GpsValidity, gps.Status, t)
				delete(_sensors, index)
			}
		case "Ignition":
			{
				template := _template.Template("Ignition").(*IgnitionTemplate)
				ignition := sensor.(*sensors.IgnitionSensor)
				compareValues("IgnitionSensor", "State", template.State, ignition.IgnitionState, t)
				delete(_sensors, index)
			}
		case "GPIO":
			{
				template := _template.Template("GPIO").(*GPIOTemplate)
				gpio := sensor.(*sensors.Inputs)
				compareValues("Inputs", "Switch1", template.Switch1, gpio.Switches[1], t)
				compareValues("Inputs", "Switch2", template.Switch2, gpio.Switches[2], t)
				compareValues("Inputs", "Switch3", template.Switch3, gpio.Switches[3], t)
				compareValues("Inputs", "Switch4", template.Switch4, gpio.Switches[4], t)
				delete(_sensors, index)
			}
		case "Relay":
			{
				template := _template.Template("Relay").(*RelaysTemplate)
				relays := sensor.(*sensors.Outputs)
				compareValues("Relay", "Relay1", template.Relay1, relays.Relays[1], t)
				compareValues("Relay", "Relay2", template.Relay2, relays.Relays[2], t)
				compareValues("Relay", "Relay3", template.Relay3, relays.Relays[3], t)
				compareValues("Relay", "Relay4", template.Relay4, relays.Relays[4], t)
				delete(_sensors, index)
			}
		case "Power":
			{
				template := _template.Template("Power").(*PowerTemplate)
				power := sensor.(*sensors.PowerSensor)
				compareValues("Power", "Supply", template.Power, power.Supply, t)
				compareValues("Power", "State", template.State, power.State, t)
				delete(_sensors, index)
			}
		case "LocId":
			{
				template := _template.Template("LocId").(*LocIDTemplate)
				queue := sensor.(*sensors.QueueSensor)
				compareValues("LocId", "LockID", template.LockID, queue.LockID, t)
				delete(_sensors, index)
			}
		case "Time":
			{
				timeSensor := sensor.(*sensors.TimeSensor)
				template := _template.Template("Time").(*TimeTemplate)
				if template.TimeStamp != "" {
					jTime := &types.JSONTime{Time: timeSensor.TimeStamp}
					bTime, mErr := jTime.MarshalJSON()
					if mErr != nil {
						t.Error("Error while time marshaling")
					}
					compareValues("timeSensor", "TimeStamp", template.TimeStamp, strings.ReplaceAll(string(bTime), "\"", ""), t)
				}
				if template.EventTime != "" {
					jTime := &types.JSONTime{Time: timeSensor.EventTimeGMT}
					bTime, mErr := jTime.MarshalJSON()
					if mErr != nil {
						t.Error("Error while time marshaling")
					}
					compareValues("timeSensor", "EventTimeStamp", template.EventTime, strings.ReplaceAll(string(bTime), "\"", ""), t)
				}
				delete(_sensors, index)
			}
		case "Odometer":
			{
				template := _template.Template("Odometer").(*TripTemplate)
				tripSensor := sensor.(*sensors.TripSensor)
				compareValues("TripSensor", "Odometer", template.Odometer, tripSensor.Odometer, t)
				delete(_sensors, index)
			}
		case "Network":
			{
				template := _template.Template("Network").(*NetworkTemplate)
				networkSensor := sensor.(*sensors.NetworkSensor)
				compareValues("NetworkSensor", "RSSI", template.RSSI, networkSensor.RSSI, t)
				compareValues("TripSensor", "CSID", template.CSID, networkSensor.CSID, t)
				delete(_sensors, index)
			}
		case "Firmware":
			{
				template := _template.Template("Firmware").(*FirmwareTemplate)
				firmwareSensor := sensor.(*sensors.Firmware)
				compareValues("Firmware", "Version", template.Version, firmwareSensor.Version, t)
				delete(_sensors, index)
			}
		default:
			{
				t.Error("Unexpected index ", index)
			}
		}
	}
	return _sensors
}

func compareValues(sensorName, field string, shouldValue, value interface{}, t *testing.T) {
	if value != shouldValue {
		t.Error(fmt.Sprintf("For sensor %v field %v has incorrect value. Should : %v, current: %v", sensorName, field, shouldValue, value))
	}
}
