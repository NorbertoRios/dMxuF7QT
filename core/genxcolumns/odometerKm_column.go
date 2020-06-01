package genxcolumns

//OdometerKmColumn odometer km
type OdometerKmColumn struct {
	RawValue interface{}
}

//Value returns value in m
func (column *OdometerKmColumn) Value() int32 {
	vKm := column.RawValue.(int32)
	return int32(vKm * 1000)
}
