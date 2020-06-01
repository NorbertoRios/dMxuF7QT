package sensors

//TemperatureSensor temp sensor
type TemperatureSensor struct {
	BaseSensor
	Imei  string
	ID    byte
	Value float32
}
