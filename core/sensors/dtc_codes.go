package sensors

import (
	"genx-go/utils"
	"strings"
)

//DTCCodes DTCCodes sensors
type DTCCodes struct {
	Codes []string
}

//BuildDTCCodesSensorFromString returns new DTCCodes sensor
func BuildDTCCodesSensorFromString(value string) ISensor {
	strArr := &utils.StringArrayUtils{Data: strings.Split(value, ",")}
	return &DTCCodes{Codes: strArr.Unique()}
}
