package genxcolumns

//SpeedColumn 13
type SpeedColumn struct {
	RawValue interface{}
}

//Value returns speed value
func (column *SpeedColumn) Value() float32 {
	iV := column.RawValue.(int32)
	return float32(iV) / 3.6
}
