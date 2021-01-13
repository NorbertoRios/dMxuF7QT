package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
	"time"
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

//ToDTO returns sensor implemetation in DTO type
func (s *GPSSensor) ToDTO() map[string]interface{} {
	hash := make(map[string]interface{})
	hash["Speed"] = s.Speed
	hash["Latitude"] = s.Latitude
	hash["Longitude"] = s.Longitude
	hash["Heading"] = s.Heading
	hash["Satellites"] = s.Satellites
	hash["GPSStat"] = s.Status
	return hash
}

//BuildGpsSensor returns new gps sensor
func BuildGpsSensor(rData map[string]interface{}) ISensor {
	if !validateGpsFields(rData) {
		return nil
	}
	posibleReasons := map[byte]byte{
		60: 1, //1 - Periodical
		10: 2, //2-GPS_LOST
		11: 3, //2-GPS_Found
	}
	speedColumn := &columns.Speed{RawValue: rData[core.Speed]}
	latitudeColumn := &columns.Coordinate{RawValue: rData[core.Latitude]}
	longitudeColumn := &columns.Coordinate{RawValue: rData[core.Longitude]}
	headingColumn := &columns.Tenth{RawValue: rData[core.Heading]}
	satelitesColumn := &columns.Byte{RawValue: rData[core.Satellites]}
	gStatusColumn := &columns.Byte{RawValue: rData[core.GPSStat]}
	sensor := &GPSSensor{
		Speed:      speedColumn.Value(),
		Latitude:   latitudeColumn.Value(),
		Longitude:  longitudeColumn.Value(),
		Heading:    float32(headingColumn.Value()),
		Satellites: satelitesColumn.Value(),
		Status:     gStatusColumn.Value(),
	}
	sensor.Trigered(rData, posibleReasons)
	sensor.symbol = "GPS"
	sensor.createdAt = time.Now().UTC()
	return sensor
}

func validateGpsFields(data map[string]interface{}) bool {
	for _, key := range GPSSensorRequiredFields {
		if _, found := data[key]; !found {
			return false
		}
	}
	return true
}
