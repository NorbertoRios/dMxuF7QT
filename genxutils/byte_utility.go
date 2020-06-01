package genxutils

import "fmt"

//ByteUtility byte utility
type ByteUtility struct {
	Data []byte
}

//String returns string from byte
func (utility *ByteUtility) String() string {
	result := ""
	for _, v := range utility.Data {
		result += fmt.Sprintf("%02X", v)
	}
	return result
}