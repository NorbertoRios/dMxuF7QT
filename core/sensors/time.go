package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
	"genx-go/types"
	"time"
)

//TimeSensor time sensor
type TimeSensor struct {
	Base
	EventTimeGMT time.Time
	TimeStamp    time.Time
}

//BuildTimeSensor returns new gps sensor
func BuildTimeSensor(data map[string]interface{}) ISensor {
	tsV, tsF := data[core.TimeStamp]
	etV, etF := data[core.EventTime]
	var timeStamp *columns.Time
	var eventTime *columns.Time
	ts := &TimeSensor{}
	ts.createdAt = time.Now().UTC()
	if tsF {
		timeStamp = &columns.Time{RawValue: tsV}
	}
	if etF {
		eventTime = &columns.Time{RawValue: etV}
	}
	if timeStamp == nil && eventTime == nil {
		return nil
	} else if timeStamp != nil && eventTime == nil {
		ts.TimeStamp = timeStamp.Value()
		return ts
	} else if timeStamp == nil && eventTime != nil {
		ts.EventTimeGMT = eventTime.Value()
		return ts
	} else {
		ts.EventTimeGMT = eventTime.Value()
		ts.TimeStamp = timeStamp.Value()
		return ts
	}
}

//ToDTO ...
func (s *TimeSensor) ToDTO() map[string]interface{} {
	hash := make(map[string]interface{})
	hash["TimeStamp"] = &types.JSONTime{Time: s.TimeStamp}
	hash["EventTime"] = &types.JSONTime{Time: s.EventTimeGMT}
	return hash
}
