package genxcolumns

//SerialColumn 1
type SerialColumn struct {
	RawValue interface{}
}

//Value returns speed value
func (column *SerialColumn) Value() string {
	return column.RawValue.(string)
}
