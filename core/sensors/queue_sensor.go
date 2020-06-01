package sensors

import (
	"genx-go/core"
	"genx-go/core/genxcolumns"
)

//QueueSensor represents queue data
type QueueSensor struct {
	BaseSensor
	LockID uint32
}

//BuildQueueSensor returns new gps sensor
func BuildQueueSensor(data map[string]interface{}) ISensor {
	if v, f := data[core.LocId]; f {
		return nil
	} else {
		lockID := &genxcolumns.LockIDColumn{RawValue: v}
		return &QueueSensor{LockID: lockID.Value()}
	}
}
