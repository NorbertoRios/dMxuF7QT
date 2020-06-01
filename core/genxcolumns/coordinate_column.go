package genxcolumns

//CoordinateColumn 3;4
type CoordinateColumn struct {
	RawValue interface{}
}

//Value returns speed value
func (column *CoordinateColumn) Value() float32 {
	return float32(float32(column.RawValue.(int32)) / 3600000.0)
}
