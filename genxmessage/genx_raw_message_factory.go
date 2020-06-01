package genxmessage

import (
	"fmt"
	"genx-go/genxutils"
	"log"
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
	hardwareBriefExpr, _ := regexp.Compile("^(?P<Serial>[0-9]{7,12}):FW:")
	diag1WireExpr, _ := regexp.Compile("(?s)1WIRE:.*((?P<Serial>[0-9]{7,12})) [0-9]{7,12}")
	paramExpr, _ := regexp.Compile("(?s)PARAMETERS.*(?P<Serial>[0-9]{7,12}) [0-9]{7,12}")
	diagExpr, _ := regexp.Compile("(?s).*(?P<Serial>[0-9]{7,12}) [0-9]{7,12}")
	diagCan, _ := regexp.Compile("(?s).*(?P<Serial>[0-9]{7,12}):.*\nCE")
	diagJExpr, _ := regexp.Compile("^(?P<Serial>[0-9]{7,12}):J|O")
	garminMessage, _ := regexp.Compile("(?s).*USER_MESSAGE.*(?P<Serial>[0-9]{7,12}) [0-9]{7,12}")
	return &RawMessageFactory{
		Maps: []ReportMap{
			{
				Type: "breport",
				Reg:  breportExpr,
			},
			{
				Type: "ack",
				Reg:  ackExpr,
			},
			{
				Type: "nack",
				Reg:  nackExpr,
			},
			{
				Type: "param",
				Reg:  allParamsExpr,
			},
			{
				Type: "poll",
				Reg:  pollExpr,
			},
			{
				Type: "poll",
				Reg:  poll1Expr,
			},
			{
				Type: "diag_hardware_brief",
				Reg:  hardwareBriefExpr,
			},
			{
				Type: "diag_1wire",
				Reg:  diag1WireExpr,
			},
			{
				Type: "param",
				Reg:  paramExpr,
			},
			{
				Type: "diag",
				Reg:  diagCan,
			},
			{
				Type: "diag",
				Reg:  diagJExpr,
			},
			{
				Type: "diag",
				Reg:  diagExpr,
			},
			{
				Type: "message",
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
	strArr := &genxutils.StringArrayUtils{Data: names}
	if id, found := strArr.IndexOf(param); found {
		return mapp.Reg.FindAllStringSubmatch(string(packet), -1)[0][id], nil
	}
	bUtil := &genxutils.ByteUtility{Data: packet}
	log.Println("[RawMessageFactory] Cant extract parameter : ", param, " from packet : ", bUtil.String())
	return "", fmt.Errorf(fmt.Sprintf("Cant find %v in message", param))
}

//Create create raw message
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
	bUtil := &genxutils.ByteUtility{Data: packet}
	log.Println("[RawMessageFactory] Cant create raw message for packet : ", bUtil.String())
	return nil
}
