package types

import "fmt"

//NewFloat64 ...
func NewFloat64(_data float64) *Float64 {
	return &Float64{
		data: _data,
	}
}

//Float64 ...
type Float64 struct {
	data float64
}

func (f64 *Float64) String() string {
	return fmt.Sprintf("%.f", f64.data)
}
