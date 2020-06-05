package utils

import "fmt"

//ByteArrayUtility byte utility
type ByteArrayUtility struct {
	Data []byte
}

//String returns string from byte
func (utils *ByteArrayUtility) String() string {
	result := ""
	for _, v := range utils.Data {
		result += fmt.Sprintf("%02X", v)
	}
	return result
}
