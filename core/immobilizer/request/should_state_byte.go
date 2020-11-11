package request

import (
	"strings"
)

//ShouldStateByte state in byte
type ShouldStateByte struct {
	Data *ChangeImmoStateRequest
}

//State returns 1 or 0 (should ralay state)
func (act *ShouldStateByte) State() byte {
	statePattern := map[string]byte{
		"mobilehigh": 0,
		"mobilelow":  1,
		"armedhigh":  1,
		"armedlow":   0,
	}
	key := strings.TrimSpace(act.Data.State) + strings.TrimSpace(act.Data.Trigger)
	if state, f := statePattern[key]; f {
		return state
	}
	return 2
}
