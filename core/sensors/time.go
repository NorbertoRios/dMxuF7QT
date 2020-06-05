package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
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
	if tsF {
		timeStamp = &columns.Time{RawValue: tsV}
	}
	if etF {
		eventTime = &columns.Time{RawValue: etV}
	}
	if timeStamp == nil && eventTime == nil {
		return nil
	} else if timeStamp != nil && eventTime == nil {
		return &TimeSensor{TimeStamp: timeStamp.Value()}
	} else if timeStamp == nil && eventTime != nil {
		return &TimeSensor{EventTimeGMT: eventTime.Value()}
	} else {
		return &TimeSensor{
			EventTimeGMT: eventTime.Value(),
			TimeStamp:    timeStamp.Value(),
		}
	}
}
