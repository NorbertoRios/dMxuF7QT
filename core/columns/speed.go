package columns

//Speed 13
type Speed struct {
	RawValue interface{}
}

//Value returns speed value
func (column *Speed) Value() float32 {
	iV := column.RawValue.(int32)
	return float32(iV) / 3.6
}
