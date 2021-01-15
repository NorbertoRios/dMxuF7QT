package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
	"genx-go/logger"
	"genx-go/types"
	"time"
)

//TimeSensor time sensor
type TimeSensor struct {
	Base
	EventTimeGMT time.Time
	TimeStamp    time.Time
}

func parseTime(strTime string) time.Time {
	datetime, terr := time.Parse("2006-01-02T15:04:05Z", strTime)
	if terr != nil {
		logger.Logger().WriteToLog(logger.Error, "[BuildAdaptedTimeSensor | BuildAdaptedTimeSensor] Error while parsing time. Error: ", terr)
		return time.Time{}
	}
	return datetime
}

//BuildAdaptedTimeSensor returns new gps sensor
func BuildAdaptedTimeSensor(data map[string]interface{}) ISensor {
	tsV, tsF := data["TimeStamp"]
	etV, etF := data["EventTime"]
	sensor := &TimeSensor{}
	sensor.symbol = "Time"
	sensor.createdAt = time.Now().UTC()
	var timeStamp time.Time
	var eventTime time.Time
	if tsF {
		timeStamp = parseTime(tsV.(string))
		sensor.TimeStamp = timeStamp
	}
	if etF {
		eventTime = parseTime(etV.(string))
		sensor.EventTimeGMT = eventTime
	}
	return sensor
}

//BuildTimeSensor returns new gps sensor
func BuildTimeSensor(data map[string]interface{}) ISensor {
	tsV, tsF := data[core.TimeStamp]
	etV, etF := data[core.EventTime]
	var timeStamp *columns.Time
	var eventTime *columns.Time
	ts := &TimeSensor{}
	ts.symbol = "Time"
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
