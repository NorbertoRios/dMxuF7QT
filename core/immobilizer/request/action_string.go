package request

import (
	"fmt"
	"genx-go/logger"
	"strings"
)

//ActionString request decorator
type ActionString struct {
	Data *ChangeImmoStateRequest
}

//Command returns action command
func (act *ActionString) Command() string {
	actionPattern := map[string]string{
		"mobilehigh": "OFF",
		"mobilelow":  "ON",
		"armedhigh":  "ON",
		"armedlow":   "OFF",
	}
	key := strings.TrimSpace(act.Data.State) + strings.TrimSpace(act.Data.Trigger)
	if action, f := actionPattern[key]; f {
		return action
	}
	logger.Logger().WriteToLog(logger.Error, fmt.Sprintf("[ActionString | Command] Unexpected action value. Incoming state:%v. Incoming trigger:%v", act.Data.State, act.Data.Trigger))
	return ""
}
