package test

import (
	"genx-go/types"
	"testing"
)

func TestByteArray(t *testing.T) {
	byteArray := &types.ByteArray{Data: []byte{1, 2, 3, 255}}
	hexString := byteArray.String()
	if hexString == "" {
		t.Error("Hex string should not be emptty")
	}
	if hexString != "010203FF" {
		t.Error("Expected 010203FF. Foud ", hexString)
	}
}

func TestByte(t *testing.T) {
	bType := &types.Byte{Data: 16}
	if !bType.BitIsSet(4) {
		t.Error("Expect true")
	}
}

func TestStringArray(t *testing.T) {
	stringArray := &types.StringArray{Data: []string{"One", "Two", "Three", "Four", "Five", "Bob", "Five", "One"}}
	index, found := stringArray.IndexOf("Bob")
	if !found || index != 5 {
		t.Error("Expected true and 5.")
	}
	uniq := stringArray.Unique()
	assertStringSlice("StringArayUtils", uniq, []string{"One", "Two", "Three", "Four", "Five", "Bob"}, t)
}

func TestStringType(t *testing.T) {
	typeBits := &types.String{Data: "10001001"}
	if typeBits.BitmaskStringToByte() != byte(137) {
		t.Error("Expected 137")
	}
	typeByte := &types.String{Data: "1C"}
	if typeByte.Byte(16) != byte(28) {
		t.Error("Expexted 28. Found ", typeByte.Byte(16))
	}
	typeFloat := &types.String{Data: "23.45678"}
	if typeFloat.Float32() != float32(23.45678) {
		t.Error("Expexted 23.45678. Found ", typeFloat.Float32())
	}
	typeUint16 := &types.String{Data: "15000"}
	if typeUint16.UInt16(10) != uint16(15000) {
		t.Error("Expexted 64000. Found ", typeUint16.UInt16(10))
	}
	typeInt32 := &types.String{Data: "101204"}
	if typeInt32.Int32(10) != int32(101204) {
		t.Error("Expexted 101204. Found ", typeUint16.Int32(10))
	}
}
