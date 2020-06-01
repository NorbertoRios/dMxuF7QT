package genxcolumns

//RelayColumn represents relay column
type RelayColumn struct {
	RawValue interface{}
}

//Value returns value of array
func (column *RelayColumn) Value() byte {
	rv := column.RawValue.(byte)
	var relays byte
	for i := 0; i < 4; i++ {
		relays <<= 1
		relays |= rv & 1
		rv >>= 1
	}
	return relays
}
