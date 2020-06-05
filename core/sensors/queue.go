package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
)

//QueueSensor represents queue data
type QueueSensor struct {
	Base
	LockID uint32
}

//BuildQueueSensor returns new gps sensor
func BuildQueueSensor(data map[string]interface{}) ISensor {
	if v, f := data[core.LocId]; f {
		return nil
	} else {
		lockID := &columns.LockID{RawValue: v}
		return &QueueSensor{LockID: lockID.Value()}
	}
}
