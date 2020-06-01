package sensors

import (
	"genx-go/core"
	"genx-go/core/genxcolumns"
)

var GPSSensorRequiredFields = []string{
	core.Speed,
	core.Latitude,
	core.Longitude,
	core.Heading,
	core.Satellites,
	core.GPSStat,
}

//GPSSensor gps sensor
type GPSSensor struct {
	BaseSensor
	Speed      float32
	Latitude   float32
	Longitude  float32
	Heading    float32
	Satellites byte
	Status     byte
}

//BuildGpsSensor returns new gps sensor
func BuildGpsSensor(rData map[string]interface{}) ISensor {
	if !validateGpsFields(rData) {
		return nil
	}
	speedColumn := &genxcolumns.SpeedColumn{RawValue: rData[core.Speed]}
	latitudeColumn := &genxcolumns.CoordinateColumn{RawValue: rData[core.Latitude]}
	longitudeColumn := &genxcolumns.CoordinateColumn{RawValue: rData[core.Longitude]}
	headingColumn := &genxcolumns.TenthColumn{RawValue: rData[core.Heading]}
	satelitesColumn := &genxcolumns.ByteColumn{RawValue: rData[core.Satellites]}
	gStatusColumn := &genxcolumns.ByteColumn{RawValue: rData[core.GPSStat]}
	return &GPSSensor{
		Speed:      speedColumn.Value(),
		Latitude:   latitudeColumn.Value(),
		Longitude:  longitudeColumn.Value(),
		Heading:    float32(headingColumn.Value()),
		Satellites: satelitesColumn.Value(),
		Status:     gStatusColumn.Value(),
	}
}

func validateGpsFields(data map[string]interface{}) bool {
	for _, key := range GPSSensorRequiredFields {
		if _, found := data[key]; !found {
			return false
		}
	}
	return true
}
