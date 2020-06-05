package columns

import (
	"genx-go/utils"
)

//Reason column
type Reason struct {
	Code  byte
	Value string
}

//BuildReasonColumn returns reason column
func BuildReasonColumn(value interface{}) *Reason {
	switch value.(type) {
	case byte:
		{
			return &Reason{Code: value.(byte)}
		}
	case string:
		{
			code, value := decodeReasonAndReasonCode(value)
			return &Reason{Code: code, Value: value}
		}
	}
	return nil
}

func decodeReasonAndReasonCode(value interface{}) (byte, string) {
	sValue, _ := value.(string)
	sUtils := utils.StringUtils{Data: sValue[1:3]}
	return sUtils.Byte(16), sValue[3:]
}
