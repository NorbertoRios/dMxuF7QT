package columns

//Serial 1
type Serial struct {
	RawValue interface{}
}

//Value returns speed value
func (column *Serial) Value() string {
	return column.RawValue.(string)
}
