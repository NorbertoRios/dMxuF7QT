package lock

import (
	"genx-go/core/lock/observers"
	"genx-go/core/lock/request"
	"testing"
)

func TestElectricLockRequest(t *testing.T) {
	data := &request.UnlockRequest{
		ExpirationTime: "2020-01-01T10:20:15Z",
		TimeToPulse:    3,
	}
	data.Port = "OUT0"
	data.Identity = "genx_000003870006"
	command := observers.NewElectricLockSetRelayDrive(data)
	expectedCommand := "SETRELAYDRIVE1X3FFFFFFF SERIALFILTER 000003870006"
	cmd := command.Command()
	if cmd != expectedCommand {
		t.Error("[TestImmobilizerCommandGeneration] Error in command generation")
	}
}
