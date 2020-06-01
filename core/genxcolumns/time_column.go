package genxcolumns

import (
	"time"
)

//TimeColumn column
type TimeColumn struct {
	RawValue interface{}
}

//Value value
func (column *TimeColumn) Value() time.Time {
	return time.Unix(int64(column.RawValue.(int32)), 0).UTC()
}
