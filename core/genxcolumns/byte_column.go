package genxcolumns

//ByteColumn switches and relays
type ByteColumn struct {
	RawValue interface{}
}

//Value value
func (column *ByteColumn) Value() byte {
	return column.RawValue.(byte)
}
