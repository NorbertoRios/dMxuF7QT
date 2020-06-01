package genxcolumns

//HundredthColumn column
type HundredthColumn struct {
	RawValue interface{}
}

//Value value
func (column *HundredthColumn) Value() float32 {
	return float32(column.RawValue.(int32) / 100.0)
}
