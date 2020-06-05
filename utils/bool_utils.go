package utils

//BoolUtils bool utils
type BoolUtils struct {
	Data bool
}

//ToByte bool to 1 or 0
func (bu *BoolUtils) ToByte() byte {
	if bu.Data {
		return byte(1)
	}
	return byte(0)
}
