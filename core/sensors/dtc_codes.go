package sensors

import (
	"genx-go/core"
	"genx-go/types"
	"strings"
	"time"
)

//DTCCodes DTCCodes sensors
type DTCCodes struct {
	Base
	Codes []string
}

//ToDTO ...
func (s *DTCCodes) ToDTO() map[string]interface{} {
	hash := make(map[string]interface{})
	hash[core.DTCCode] = s.Codes
	return hash
}

//BuildDTCCodesSensorFromString returns new DTCCodes sensor
func BuildDTCCodesSensorFromString(value string) ISensor {
	strArr := &types.StringArray{Data: strings.Split(value, ",")}
	dtc := &DTCCodes{Codes: strArr.Unique()}
	dtc.symbol = "DTCCode"
	dtc.createdAt = time.Now().UTC()
	return dtc
}
