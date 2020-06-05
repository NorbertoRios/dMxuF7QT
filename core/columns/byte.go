package columns

//Byte switches and relays
type Byte struct {
	RawValue interface{}
}

//Value value
func (column *Byte) Value() byte {
	return column.RawValue.(byte)
}
