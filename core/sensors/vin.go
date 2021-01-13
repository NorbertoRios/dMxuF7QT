package sensors

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
