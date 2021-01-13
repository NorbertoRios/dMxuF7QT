package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
	"genx-go/logger"
	"strconv"
	"time"
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
		iButton := &IButton{BtnID: iBiD.Value()}
		iButton.symbol = "IBID"
		iButton.createdAt = time.Now().UTC()
		return iButton
	}
	return nil
}

//ToDTO returns sensor implemetation in DTO type
func (s *IButton) ToDTO() map[string]interface{} {
	hash := make(map[string]interface{})
	hash[s.symbol] = s.BtnID
	return hash
}

//BuildIButtonSensorFromString returns new driver id sensor from string
func BuildIButtonSensorFromString(value string) ISensor {
	intValue, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		logger.Logger().WriteToLog(logger.Error, "[BuildDriverIDSensorFromString] Cant parse DriverID sensor from : ", value)
		return nil
	}
	iButton := &IButton{BtnID: int32(intValue)}
	iButton.symbol = "IBID"
	iButton.createdAt = time.Now().UTC()
	return iButton
}
