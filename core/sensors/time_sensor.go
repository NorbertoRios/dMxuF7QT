package sensors

import (
	"genx-go/core"
	"genx-go/core/genxcolumns"
	"time"
)

//TimeSensor time sensor
type TimeSensor struct {
	BaseSensor
	EventTimeGMT time.Time
	TimeStamp    time.Time
}

//BuildTimeSensor returns new gps sensor
func BuildTimeSensor(data map[string]interface{}) ISensor {
	tsV, tsF := data[core.TimeStamp]
	etV, etF := data[core.EventTime]
	var timeStamp *genxcolumns.TimeColumn
	var eventTime *genxcolumns.TimeColumn
	if tsF {
		timeStamp = &genxcolumns.TimeColumn{RawValue: tsV}
	}
	if etF {
		eventTime = &genxcolumns.TimeColumn{RawValue: etV}
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
