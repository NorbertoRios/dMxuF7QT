package sensors

import (
	"genx-go/genxutils"
	"strings"
)

//DTCCodesSensor DTCCodes sensors
type DTCCodesSensor struct {
	Codes []string
}

//BuildDTCCodesSensorFromString returns new DTCCodes sensor
func BuildDTCCodesSensorFromString(value string) ISensor {
	strArr := &genxutils.StringArrayUtils{Data: strings.Split(value, ",")}
	return &DTCCodesSensor{Codes: strArr.Unique()}
}
