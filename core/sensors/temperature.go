package sensors

//TemperatureSensor temp sensor
type TemperatureSensor struct {
	Base
	Imei  string
	ID    byte
	Value float32
}
