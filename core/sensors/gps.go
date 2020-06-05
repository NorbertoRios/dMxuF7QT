package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
)

//GPSSensorRequiredFields required fields for GPS sensor
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
	Base
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
	speedColumn := &columns.Speed{RawValue: rData[core.Speed]}
	latitudeColumn := &columns.Coordinate{RawValue: rData[core.Latitude]}
	longitudeColumn := &columns.Coordinate{RawValue: rData[core.Longitude]}
	headingColumn := &columns.Tenth{RawValue: rData[core.Heading]}
	satelitesColumn := &columns.Byte{RawValue: rData[core.Satellites]}
	gStatusColumn := &columns.Byte{RawValue: rData[core.GPSStat]}
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
