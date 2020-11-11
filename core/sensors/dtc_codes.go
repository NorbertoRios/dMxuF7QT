package sensors

import (
	"genx-go/types"
	"strings"
)

//DTCCodes DTCCodes sensors
type DTCCodes struct {
	Codes []string
}

//BuildDTCCodesSensorFromString returns new DTCCodes sensor
func BuildDTCCodesSensorFromString(value string) ISensor {
	strArr := &types.StringArray{Data: strings.Split(value, ",")}
	return &DTCCodes{Codes: strArr.Unique()}
}
