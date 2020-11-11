package types

//Bool bool utils
type Bool struct {
	Data bool
}

//ToByte bool to 1 or 0
func (b *Bool) ToByte() byte {
	if b.Data {
		return byte(1)
	}
	return byte(0)
}
