package columns

//LockID column for lockId
type LockID struct {
	RawValue interface{}
}

//Value returns reedable value
func (column *LockID) Value() uint32 {
	return column.RawValue.(uint32)
}
