package sensors

import (
	"genx-go/core"
	"genx-go/core/genxcolumns"
)

//NetworkSensor network sensor
type NetworkSensor struct {
	BaseSensor
	RSSI int8
	CSID int32
}

//BuildNetworkSensor returns new gps sensor
func BuildNetworkSensor(data map[string]interface{}) ISensor {
	rv, rf := data[core.RSSI]
	cv, cf := data[core.CSID]
	if !rf && !cf {
		return nil
	}
	var RSSI *genxcolumns.RSSIColumn
	var CSID *genxcolumns.TenthColumn
	if cf && !rf {
		CSID = &genxcolumns.TenthColumn{RawValue: cv}
		return &NetworkSensor{CSID: CSID.Value()}
	}
	if !cf && rf {
		RSSI = &genxcolumns.RSSIColumn{RawValue: rv}
		return &NetworkSensor{RSSI: RSSI.Value()}
	}
	RSSI = &genxcolumns.RSSIColumn{RawValue: rv}
	CSID = &genxcolumns.TenthColumn{RawValue: cv}
	return &NetworkSensor{
		RSSI: RSSI.Value(),
		CSID: CSID.Value(),
	}
}
