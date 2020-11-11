package columns

//Coordinate 3;4
type Coordinate struct {
	RawValue interface{}
}

//Value returns speed value
func (column *Coordinate) Value() float32 {
	return float32(float32(column.RawValue.(int32)) / 3600000.0)
}
