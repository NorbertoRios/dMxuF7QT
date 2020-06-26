package utils

import (
	"fmt"
	"strconv"
	"strings"
)

//StringUtils utils for string
type StringUtils struct {
	Data string
}

//Identity returns identity from serial
func (utils *StringUtils) Identity() string {
	serial := utils.Data
	for l := len(utils.Data); l < 12; l++ {
		serial = fmt.Sprintf("0%v", serial)
	}
	return fmt.Sprintf("genx_%v", serial)
}

//BitmaskStringToByte string value to byte
func (utils *StringUtils) BitmaskStringToByte() byte {
	result := byte(0)
	bitsArr := strings.Split(utils.Data, "")
	lngth := len(bitsArr) - 1
	for i := lngth; i > -1; i-- {
		if bitsArr[i] == "1" {
			result = 1<<(lngth-i) | result
		}
	}
	return result
}

//Byte returns byte from string
func (utils *StringUtils) Byte(base int) byte {
	value, _ := strconv.ParseUint(utils.Data, base, 8)
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

//UInt16 string value to UInt16
func (utils *StringUtils) UInt16(base int) uint16 {
	value, _ := strconv.ParseInt(utils.Data, base, 16)
	return uint16(value)
}
