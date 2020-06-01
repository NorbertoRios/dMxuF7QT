package genxcolumns

//OdometerMiColumn odometer km
type OdometerMiColumn struct {
	RawValue interface{}
}

//Value returns value in m
func (column *OdometerMiColumn) Value() int32 {
	vMi := column.RawValue.(int32)
	return int32(float32(vMi) / float32(0.00062137))
}
