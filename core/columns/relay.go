package columns

//Relay represents relay column
type Relay struct {
	RawValue interface{}
}

//Value returns value of array
func (column *Relay) Value() byte {
	rv := column.RawValue.(byte)
	var relays byte
	for i := 0; i < 4; i++ {
		relays <<= 1
		relays |= rv & 1
		rv >>= 1
	}
	return relays
}
