package columns

//Temperature represents 1wire temp sensor
type Temperature struct {
	RawValue interface{}
}

//Value returns value (index from 0 to 3)
func (column *Temperature) Value(index int) int {
	value := column.RawValue.([]byte)
	return int(value[index])<<8 | int(value[index+1])
}
