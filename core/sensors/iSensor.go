package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
)

//ISensor sensor's intergace
type ISensor interface {
}

//Trigered triger for sensor
func Trigered(rData map[string]interface{}, posibleReasons map[byte]byte) byte {
	r, f := rData[core.Reason]
	if !f {
		return byte(0)
	}
	reason := columns.BuildReasonColumn(r)
	if reason == nil {
		return byte(0)
	}
	if v, f := posibleReasons[reason.Code]; !f {
		return byte(0)
	} else {
		return v
	}
}
