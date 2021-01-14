package dto

//DtoMessage ...
type DtoMessage struct {
	Parameter24 []string                      `json:"Param24"`
	Data        map[string]interface{}        `json:"Data"`
	Sensors     map[string]*TemperatureSensor `json:"ts"`
	SID         uint64                        `json:"sid"`
}

//GetValue from Data field
func (m *DtoMessage) GetValue(key string) (value interface{}, found bool) {
	value, found = m.Data[key]
	return value, found
}
