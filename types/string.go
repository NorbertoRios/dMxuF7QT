package types

import (
	"strconv"
	"strings"
)

//String str for string
type String struct {
	Data string
}

//BitmaskStringToInt32 string value to byte
func (str *String) BitmaskStringToInt32() int32 {
	result := int32(0)
	bitsArr := strings.Split(str.Data, "")
	lngth := len(bitsArr) - 1
	for i := lngth; i > -1; i-- {
		if bitsArr[i] == "1" {
			result = 1<<(lngth-i) | result
		}
	}
	return result
}

//BitmaskStringToByte string value to byte
func (str *String) BitmaskStringToByte() byte {
	result := byte(0)
	bitsArr := strings.Split(str.Data, "")
	lngth := len(bitsArr) - 1
	for i := lngth; i > -1; i-- {
		if bitsArr[i] == "1" {
			result = 1<<(lngth-i) | result
		}
	}
	return result
}

//Byte returns byte from string
func (str *String) Byte(base int) byte {
	value, err := strconv.ParseUint(str.Data, base, 8)
	if err != nil {
		return 0
	}
	return byte(value)
}

//Float32 string value to float32
func (str *String) Float32() float32 {
	value, _ := strconv.ParseFloat(str.Data, 32)
	return float32(value)
}

//Int32 string value to int32
func (str *String) Int32(base int) int32 {
	value, _ := strconv.ParseInt(str.Data, base, 32)
	return int32(value)
}

//UInt16 string value to UInt16
func (str *String) UInt16(base int) uint16 {
	value, _ := strconv.ParseInt(str.Data, base, 16)
	return uint16(value)
}
