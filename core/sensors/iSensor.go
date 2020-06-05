package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
)

//ISensor sensor's intergace
type ISensor interface {
}

//Trigered triger for sensor
func Trigered(rData map[string]interface{}, posibleReasons map[byte]bool) bool {
	r, f := rData[core.Reason]
	if !f {
		return false
	}
	reason := columns.BuildReasonColumn(r)
	if reason == nil {
		return false
	}
	return posibleReasons[reason.Code]
}
