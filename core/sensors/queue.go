package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
	"time"
)

//QueueSensor represents queue data
type QueueSensor struct {
	Base
	LockID uint32
}

//BuildAdaptedQueueSensor ...
func BuildAdaptedQueueSensor(_lockId uint32) ISensor {
	sensor := &QueueSensor{LockID: _lockId}
	sensor.symbol = "LocId"
	sensor.createdAt = time.Now().UTC()
	return sensor
}

//BuildQueueSensor returns new gps sensor
func BuildQueueSensor(data map[string]interface{}) ISensor {
	if v, f := data[core.LocId]; f {
		return nil
	} else {
		lockID := &columns.LockID{RawValue: v}
		sensor := &QueueSensor{LockID: lockID.Value()}
		posibleReasons := map[byte]byte{
			6: 1, // 1- Periodical
			8: 1,
		}
		sensor.Trigered(data, posibleReasons)
		sensor.symbol = "LocId"
		sensor.createdAt = time.Now().UTC()
		return sensor
	}
}

//ToDTO ..
func (s *QueueSensor) ToDTO() map[string]interface{} {
	hash := make(map[string]interface{})
	hash[s.symbol] = s.LockID
	return hash
}
