package genxcolumns

//TenthColumn column
type TenthColumn struct {
	RawValue interface{}
}

//Value value
func (column *TenthColumn) Value() int32 {
	return column.RawValue.(int32)
}
