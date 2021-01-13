package dto

//TemperatureSensor ....
type TemperatureSensor struct {
	ID                   string  `json:"Id"`
	Value                float32 `json:"TemperatureValue"`
	Event                int     `json:"Event"`
	Events               []int   `json:"Events"`
	TemperatureThreshold float32 `json:"TemperatureThreshold"`
	LogID                int     `json:"LogID"`
}
