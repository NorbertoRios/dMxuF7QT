package test

import (
	"fmt"
	"genx-go/message"
	"genx-go/parser"
	"testing"
)

func TestParametersMessageParsing(t *testing.T) {
	packet := []byte("PARAMETERS\n56=;\n24=1.2.3.4.5.6.7.89.86.4.2.3;\n500=GFM1212312;\n000003912835 3912835\n")
	rm := factory.BuildRawMessage(packet)
	parser := parser.ConstructParametersMessageParser()
	deviceMessage := parser.Parse(rm)
	if deviceMessage == nil {
		t.Error("Message cant be null")
	}
	shouldValues := map[string]string{
		"56":  "56=;",
		"24":  "24=1.2.3.4.5.6.7.89.86.4.2.3;",
		"500": "500=GFM1212312;",
	}
	switch deviceMessage.(type) {
	case *message.ParametersMessage:
		{
			EqualSlice(deviceMessage.(*message.ParametersMessage).Parameters, shouldValues, t)
		}
	default:
		{
			t.Error("Error in message type")
		}
	}
}

func TestAllParametersMessageParsing(t *testing.T) {
	packet := []byte("ALL-PARAMETERS\n9=3870006;\n500=GFM121232;\n501=259200;\n503=100;\n505=0;\n")
	rm := factory.BuildRawMessage(packet)
	parser := parser.ConstructParametersMessageParser()
	deviceMessage := parser.Parse(rm)
	if deviceMessage == nil {
		t.Error("Message cant be null")
	}
	shouldValues := map[string]string{
		"9":   "9=3870006;",
		"501": "501=259200;",
		"503": "503=100;",
		"505": "505=0;",
		"500": "500=GFM121232;",
	}
	switch deviceMessage.(type) {
	case *message.ParametersMessage:
		{
			EqualSlice(deviceMessage.(*message.ParametersMessage).Parameters, shouldValues, t)
		}
	default:
		{
			t.Error("Error in message type")
		}
	}
}

func EqualSlice(current, should map[string]string, t *testing.T) bool {
	if len(current) != len(should) {
		t.Error("Maps are not equal. Maps are different lengths")
	}
	for key, value := range current {
		if value != should[key] {
			t.Error(fmt.Sprintf("Maps are not equal. Current %v . Should %v", value, should[key]))
		}
	}
	return true
}
