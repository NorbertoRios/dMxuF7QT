package test

import (
	"genx-go/core/immobilizer/observers"
	"genx-go/core/immobilizer/request"
	"testing"
)

func TestImmobilizerCommandGeneration(t *testing.T) {
	data := &request.ChangeImmoStateRequest{
		SafetyOption: true,
		State:        "armed",
		Trigger:      "high",
	}
	data.Port = "OUT0"
	data.Identity = "genx_000003870006"
	command := observers.NewSetRelayDrive(data)
	expectedCommand := "SETRELAYDRIVE1ON SERIALFILTER 000003870006;BACKUPNVRAM"
	cmd := command.Command()
	if cmd != expectedCommand {
		t.Error("[TestImmobilizerCommandGeneration] Error in command generation")
	}
}
