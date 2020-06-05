package test

import (
	"genx-go/utils"
	"testing"
)

func TestByteArrayUtils(t *testing.T) {
	byteArrayUtils := &utils.ByteArrayUtility{Data: []byte{1, 2, 3, 255}}
	hexString := byteArrayUtils.String()
	if hexString == "" {
		t.Error("Hex string should not be emptty")
	}
	if hexString != "010203FF" {
		t.Error("Expected 010203FF. Foud ", hexString)
	}
}

func TestByteUtils(t *testing.T) {
	byteUtils := &utils.ByteUtility{Data: 16}
	if !byteUtils.BitIsSet(4) {
		t.Error("Expect true")
	}
}

func TestStringArrayUtils(t *testing.T) {
	stringArrayUtils := &utils.StringArrayUtils{Data: []string{"One", "Two", "Three", "Four", "Five", "Bob", "Five", "One"}}
	index, found := stringArrayUtils.IndexOf("Bob")
	if !found || index != 5 {
		t.Error("Expected true and 5.")
	}
	uniq := stringArrayUtils.Unique()
	assertStringSlice("StringArayUtils", uniq, []string{"One", "Two", "Three", "Four", "Five", "Bob"}, t)
}

func TestStringUtils(t *testing.T) {
	stringUtilsBits := &utils.StringUtils{Data: "10001001"}
	if stringUtilsBits.BitmaskStringToByte() != byte(137) {
		t.Error("Expected 137")
	}
	stringUtilsByte := &utils.StringUtils{Data: "1C"}
	if stringUtilsByte.Byte(16) != byte(28) {
		t.Error("Expexted 28. Found ", stringUtilsByte.Byte(16))
	}
	stringUtilsFloat := &utils.StringUtils{Data: "23.45678"}
	if stringUtilsFloat.Float32() != float32(23.45678) {
		t.Error("Expexted 23.45678. Found ", stringUtilsFloat.Float32())
	}
	stringUtilsUint16 := &utils.StringUtils{Data: "15000"}
	if stringUtilsUint16.UInt16(10) != uint16(15000) {
		t.Error("Expexted 64000. Found ", stringUtilsUint16.UInt16(10))
	}
	stringUtilsInt32 := &utils.StringUtils{Data: "101204"}
	if stringUtilsInt32.Int32(10) != int32(101204) {
		t.Error("Expexted 101204. Found ", stringUtilsUint16.Int32(10))
	}
}
