package columns

//Tenth column
type Tenth struct {
	RawValue interface{}
}

//Value value
func (column *Tenth) Value() int32 {
	return column.RawValue.(int32)
}
