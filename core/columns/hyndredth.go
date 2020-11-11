package columns

//Hundredth column
type Hundredth struct {
	RawValue interface{}
}

//Value value
func (column *Hundredth) Value() float32 {
	return float32(column.RawValue.(int32) / 100.0)
}
