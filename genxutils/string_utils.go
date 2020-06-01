package genxutils

import (
	"strconv"
)

//StringUtils utils for string
type StringUtils struct {
	Data string
}

//Byte string value to byte
func (utils *StringUtils) Byte() byte {
	value, _ := strconv.Atoi(utils.Data)
	return byte(value)
}

//Float32 string value to float32
func (utils *StringUtils) Float32() float32 {
	value, _ := strconv.ParseFloat(utils.Data, 32)
	return float32(value)
}

//Int32 string value to int32
func (utils *StringUtils) Int32(base int) int32 {
	value, _ := strconv.ParseInt(utils.Data, base, 32)
	return int32(value)
}