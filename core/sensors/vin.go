package sensors

import (
	"time"
)

//BuildVINSensor ...
func BuildVINSensor(_vin string) *VINSensor {
	sensor := &VINSensor{
		VIN: _vin,
	}
	sensor.symbol = "VIN"
	sensor.createdAt = time.Now().UTC()
	return sensor
}

//VINSensor vin number
type VINSensor struct {
	Base
	VIN string
}

//ToDTO ...
func (s *VINSensor) ToDTO() map[string]interface{} {
	hash := make(map[string]interface{})
	hash["VIN"] = s.VIN
	return hash
}
