package columns

import (
	"time"
)

//Time column
type Time struct {
	RawValue interface{}
}

//Value value
func (column *Time) Value() time.Time {
	return time.Unix(int64(column.RawValue.(int32)), 0).UTC()
}
