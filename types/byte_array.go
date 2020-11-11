package types

import "fmt"

//ByteArray byte array
type ByteArray struct {
	Data []byte
}

//String returns string from byte
func (bArray *ByteArray) String() string {
	result := ""
	for _, v := range bArray.Data {
		result += fmt.Sprintf("%02X", v)
	}
	return result
}
