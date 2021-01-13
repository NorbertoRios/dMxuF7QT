package sensors

//ISensor sensor's intergace
type ISensor interface {
	ToDTO() map[string]interface{}
	Symbol() string
}
