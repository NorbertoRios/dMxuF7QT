package columns

//OdometerKm odometer km
type OdometerKm struct {
	RawValue interface{}
}

//Value returns value in m
func (column *OdometerKm) Value() int32 {
	vKm := column.RawValue.(int32)
	return int32(vKm * 1000)
}
