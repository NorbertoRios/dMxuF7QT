package genxmessage

import (
	"genx-go/core/sensors"
	"regexp"
)

//GenxTextMessageParser sensors builder for text message
type GenxTextMessageParser struct {
	Map map[*regexp.Regexp]func(string) sensors.ISensor
}

//Parse parse
func (hms *GenxTextMessageParser) Parse(rawMessage *RawMessage) []sensors.ISensor {
	returnedValue := make([]sensors.ISensor, 0)
	for expr, function := range hms.Map {
		if expr.Match(rawMessage.RawData) {
			substr := expr.FindAllStringSubmatch(string(rawMessage.RawData), -1)[0][1]
			if substr != "" {
				sensor := function(substr)
				returnedValue = append(returnedValue, sensor)
			}
		}
	}
	return returnedValue
}
