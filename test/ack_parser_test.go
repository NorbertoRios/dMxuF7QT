package test

import (
	"genx-go/message"
	"genx-go/parser"
	"testing"
)

func TestAckMessageParsing(t *testing.T) {
	packet := []byte("000003870006 ACK < SETPARAM;7=216.187.77.150;ENDPARAM;BACKUPNVRAM;>")
	rm := factory.BuildRawMessage(packet)
	parser := parser.ConstructAckMesageParser()
	deviceMessage := parser.Parse(rm)
	if deviceMessage == nil {
		t.Error("Message cant be null")
	}
	switch deviceMessage.(type) {
	case *message.AckMessage:
		{
			if deviceMessage.(*message.AckMessage).Value != "SETPARAM;7=216.187.77.150;ENDPARAM;BACKUPNVRAM;" {
				t.Error("Wrong parameter")
			}
		}
	}
}
