package utils

//ByteUtility byte utility
type ByteUtility struct {
	Data byte
}

//BitIsSet true if bit set
func (bu *ByteUtility) BitIsSet(index int) bool {
	return (bu.Data & (1 << index)) != 0
}

//ToBool convert byte to bool
func (bu *ByteUtility) ToBool() bool {
	return bu.Data == byte(1)
}
