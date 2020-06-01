package genxcolumns

//RSSIColumn represents RSSI column
type RSSIColumn struct {
	RawValue interface{}
}

//Value returns rssi value
func (column *RSSIColumn) Value() int8 {
	return int8(column.RawValue.(byte))
}
