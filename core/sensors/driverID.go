package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
	"log"
	"strconv"
)

//IButton driver id sensor
type IButton struct {
	Base
	BtnID int32
}

//BuildIButtonSensor returns new driver id sensor
func BuildIButtonSensor(data map[string]interface{}) ISensor {
	if v, f := data[core.IBID]; f {
		iBiD := &columns.Tenth{RawValue: v}
		return &IButton{BtnID: iBiD.Value()}
	}
	return nil
}

//BuildIButtonSensorFromString returns new driver id sensor from string
func BuildIButtonSensorFromString(value string) ISensor {
	intValue, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		log.Println("[BuildDriverIDSensorFromString] Cant parse DriverID sensor from : ", value)
		return nil
	}
	return &IButton{BtnID: int32(intValue)}
}
