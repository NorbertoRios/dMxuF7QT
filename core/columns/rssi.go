package columns

//RSSI represents RSSI column
type RSSI struct {
	RawValue interface{}
}

//Value returns rssi value
func (column *RSSI) Value() int8 {
	return int8(column.RawValue.(byte))
}
