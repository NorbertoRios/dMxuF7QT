package sensors

import (
	"genx-go/core"
	"genx-go/core/genxcolumns"
	"log"
	"strconv"
)

//DriverIDSensor driver id sensor
type IButtonSensor struct {
	BaseSensor
	BtnID int32
}

//BuildDriverIDSensor returns new driver id sensor
func BuildDriverIDSensor(data map[string]interface{}) ISensor {
	if v, f := data[core.IBID]; f {
		iBiD := &genxcolumns.TenthColumn{RawValue: v}
		return &IButtonSensor{BtnID: iBiD.Value()}
	}
	return nil
}

//BuildDriverIDSensorFromString returns new driver id sensor from string
func BuildDriverIDSensorFromString(value string) ISensor {
	intValue, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		log.Println("[BuildDriverIDSensorFromString] Cant parse DriverID sensor from : ", value)
		return nil
	}
	return &IButtonSensor{BtnID: int32(intValue)}
}
