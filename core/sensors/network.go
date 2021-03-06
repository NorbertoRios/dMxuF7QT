package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
	"time"
)

//NetworkSensor network sensor
type NetworkSensor struct {
	Base
	RSSI int8
	CSID int32
}

//BuildAdaptedNetworkSensor ...
func BuildAdaptedNetworkSensor(data map[string]interface{}) ISensor {
	sensor := &NetworkSensor{}
	sensor.symbol = "Network"
	sensor.createdAt = time.Now().UTC()
	if v, f := data["RSSI"]; f {
		sensor.RSSI = int8(v.(float64))
	}
	if v, f := data["CSID"]; f {
		sensor.CSID = int32(v.(float64))
	}
	return sensor
}

//BuildNetworkSensor returns new gps sensor
func BuildNetworkSensor(data map[string]interface{}) ISensor {
	rv, rf := data[core.RSSI]
	cv, cf := data[core.CSID]
	if !rf && !cf {
		return nil
	}
	var RSSI *columns.RSSI
	var CSID *columns.Tenth
	if cf && !rf {
		CSID = &columns.Tenth{RawValue: cv}
		return &NetworkSensor{CSID: CSID.Value()}
	}
	if !cf && rf {
		RSSI = &columns.RSSI{RawValue: rv}
		return &NetworkSensor{RSSI: RSSI.Value()}
	}
	RSSI = &columns.RSSI{RawValue: rv}
	CSID = &columns.Tenth{RawValue: cv}
	n := &NetworkSensor{
		RSSI: RSSI.Value(),
		CSID: CSID.Value(),
	}
	n.symbol = "Network"
	n.createdAt = time.Now().UTC()
	return n
}

//ToDTO ..
func (s *NetworkSensor) ToDTO() map[string]interface{} {
	hash := make(map[string]interface{})
	hash["RSSI"] = s.RSSI
	hash["CSID"] = s.CSID
	return hash
}
