package request

import (
	"fmt"
	"genx-go/core/request"
	"genx-go/logger"
	"genx-go/types"
	"reflect"
	"strings"
	"time"
)

//UnlockRequest ...
type UnlockRequest struct {
	request.ChangeStateRequest
	TimeToPulse    int    `json:"TimeToPulse"`
	ExpirationTime string `json:"ExpirationTime"`
}

//Pulse ...
func (data *UnlockRequest) Pulse() string {
	var bitMask string
	bitCount := data.TimeToPulse * 10
	for i := 0; i < bitCount; i++ {
		bitMask = fmt.Sprintf("1%v", bitMask)
	}
	sType := &types.String{Data: bitMask}
	return strings.ToUpper(fmt.Sprintf("%02x", sType.BitmaskStringToInt32()))
}

// 111111111111111111111111111111
//Equal ...
func (data *UnlockRequest) Equal(req request.IRequest) bool {
	if _, v := req.(*UnlockRequest); v {
		return reflect.DeepEqual(data, req)
	}
	return false
}

//Time ...
func (data *UnlockRequest) Time() time.Time {
	t, err := time.Parse(time.RFC3339, data.ExpirationTime)
	if err == nil {
		return t
	}
	logger.Logger().WriteToLog(logger.Error, "[UnlockRequest | Time] Error while converting string \"", data.ExpirationTime, "\" to date time.")
	return time.Time{}
}
