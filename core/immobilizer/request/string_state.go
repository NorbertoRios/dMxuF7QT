package request

import (
	"fmt"
	"genx-go/logger"
	"strings"
)

const (
	on  = "on"
	off = "off"
)

//NewImmoStateRelayBased ...
func NewImmoStateRelayBased(bState byte, _trigger string) *ImmoStateRelayBased {
	var _state string
	if bState == 0 {
		_state = off
	} else if bState == 1 {
		_state = on
	}
	return &ImmoStateRelayBased{
		state:   _state,
		trigger: _trigger,
	}
}

//ImmoStateRelayBased ...
type ImmoStateRelayBased struct {
	state        string //On or Off
	trigger      string
	statePattern map[string]string
}

//State ...
func (s *ImmoStateRelayBased) State() string {
	key := strings.TrimSpace(s.state) + strings.TrimSpace(s.trigger)
	return s.mapByteState(key)
}

func (s *ImmoStateRelayBased) mapByteState(_key string) string {
	statePattern := map[string]string{
		"offhigh": "mobile",
		"onlow":   "mobile",
		"onhigh":  "armed",
		"offlow":  "armed",
	}
	if state, f := statePattern[_key]; f {
		return state
	}
	logger.Logger().WriteToLog(logger.Error, fmt.Sprintf("[ImmoStateRelayBased | State] Unexpected state value. Incoming state:%v. Incoming trigger:%v", s.state, s.trigger))
	return ""
}
