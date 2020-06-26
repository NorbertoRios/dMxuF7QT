package test

import (
	"genx-go/parser"
	"testing"
)

func TestAckMessageParsing(t *testing.T) {
	packet := []byte("000003870006 ACK < SETPARAM;7=216.187.77.150;ENDPARAM;BACKUPNVRAM;>")
	rm := factory.BuildRawMessage(packet)
	parser := parser.ConstructAckMesageParser()
	message := parser.Parse(rm)
	if message == nil {
		t.Error("Message cant be null")
	}
	if message.Value != "SETPARAM;7=216.187.77.150;ENDPARAM;BACKUPNVRAM;" {
		t.Error("Wrong parameter")
	}	
}
