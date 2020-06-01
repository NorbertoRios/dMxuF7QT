package genxcolumns

//LockIDColumn column for lockId
type LockIDColumn struct {
	RawValue interface{}
}

//Value returns reedable value
func (column *LockIDColumn) Value() uint32 {
	return column.RawValue.(uint32)
}
