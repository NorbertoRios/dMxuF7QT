package types

//Byte byte utility
type Byte struct {
	Data byte
}

//BitIsSet true if bit set
func (b *Byte) BitIsSet(index int) bool {
	return (b.Data & (1 << index)) != 0
}

//ToBool convert byte to bool
func (b *Byte) ToBool() bool {
	return b.Data == byte(1)
}
