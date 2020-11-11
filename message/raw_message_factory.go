package message

import (
	"fmt"
	"genx-go/logger"
	"genx-go/message/messagetype"
	"genx-go/types"
	"regexp"
)

//CounstructRawMessageFactory returns raw message factory
func CounstructRawMessageFactory() *RawMessageFactory {
	ackExpr, _ := regexp.Compile("(?P<Serial>[0-9]{7,12}) ACK <")
	nackExpr, _ := regexp.Compile("(?P<Serial>[0-9]{7,12}) NAK- ")
	allParamsExpr, _ := regexp.Compile("(?s)ALL-PARAMETERS.*9=(?P<Serial>[0-9]{7,12})")
	pollExpr, _ := regexp.Compile("(?s)(?P<Serial>[0-9]{7,12}).*POLL.*(SN:[0-9]{7,12})")
	poll1Expr, _ := regexp.Compile("(?s)^(?P<Serial>[0-9]{7,12}),")
	breportExpr, _ := regexp.Compile("^(?P<Serial>[0-9]{7,12})\x00")
	hardwareBriefExpr, _ := regexp.Compile(`MODEL\:(.*)?\nSN:(?P<Serial>[0-9]{7,12})\nFW\:(.*)?\nHW\:`)
	diag1WireExpr, _ := regexp.Compile("(?s)1WIRE:.*((?P<Serial>[0-9]{7,12})) [0-9]{7,12}")
	paramExpr, _ := regexp.Compile("(?s)PARAMETERS.*(?P<Serial>[0-9]{7,12}) [0-9]{7,12}")
	diagExpr, _ := regexp.Compile("(?s).*(?P<Serial>[0-9]{7,12}) [0-9]{7,12}")
	diagCan, _ := regexp.Compile("(?s).*(?P<Serial>[0-9]{7,12}):.*\nCE")
	diagJExpr, _ := regexp.Compile("^(?P<Serial>[0-9]{7,12}):J|O")
	garminMessage, _ := regexp.Compile("(?s).*USER_MESSAGE.*(?P<Serial>[0-9]{7,12}) [0-9]{7,12}")
	return &RawMessageFactory{
		Maps: []ReportMap{
			{
				Type: messagetype.BinaryLocation,
				Reg:  breportExpr,
			},
			{
				Type: messagetype.Ack,
				Reg:  ackExpr,
			},
			{
				Type: messagetype.Nack,
				Reg:  nackExpr,
			},
			{
				Type: messagetype.Parameter,
				Reg:  allParamsExpr,
			},
			{
				Type: messagetype.Poll,
				Reg:  pollExpr,
			},
			{
				Type: messagetype.Poll,
				Reg:  poll1Expr,
			},
			{
				Type: messagetype.DiagHardware,
				Reg:  hardwareBriefExpr,
			},
			{
				Type: messagetype.Diag1Wire,
				Reg:  diag1WireExpr,
			},
			{
				Type: messagetype.Parameter,
				Reg:  paramExpr,
			},
			{
				Type: messagetype.DiagCAN,
				Reg:  diagCan,
			},
			{
				Type: messagetype.DiagJBUS,
				Reg:  diagJExpr,
			},
			{
				Type: messagetype.Diag,
				Reg:  diagExpr,
			},
			{
				Type: messagetype.GraminMessage,
				Reg:  garminMessage,
			},
		},
	}
}

//RawMessageFactory factory for raw message
type RawMessageFactory struct {
	Maps []ReportMap
}

func (factory *RawMessageFactory) extractParam(index int, param string, packet []byte) (string, error) {
	mapp := factory.Maps[index]
	names := mapp.Reg.SubexpNames()
	strArr := &types.StringArray{Data: names}
	if id, found := strArr.IndexOf(param); found {
		return mapp.Reg.FindAllStringSubmatch(string(packet), -1)[0][id], nil
	}
	bType := &types.ByteArray{Data: packet}
	logger.Logger().WriteToLog(logger.Error, "[RawMessageFactory] Cant extract parameter : ", param, " from packet : ", bType.String())
	return "", fmt.Errorf(fmt.Sprintf("Cant find %v in message", param))
}

//BuildRawMessage create raw message
func (factory *RawMessageFactory) BuildRawMessage(packet []byte) *RawMessage {
	for index, mapp := range factory.Maps {
		if mapp.Reg.Match(packet) {
			serial, sErr := factory.extractParam(index, "Serial", packet)
			if sErr != nil {
				return nil
			}
			return &RawMessage{
				SerialNumber: serial,
				MessageType:  mapp.Type,
				RawData:      packet,
			}
		}
	}
	bArray := &types.ByteArray{Data: packet}
	logger.Logger().WriteToLog(logger.Error, "[RawMessageFactory] Cant create raw message for packet : ", bArray.String())
	return nil
}
