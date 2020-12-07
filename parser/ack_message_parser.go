package parser

import (
	"genx-go/logger"
	"genx-go/message"
	"regexp"
	"strings"
)

//ConstructAckMesageParser returns new AckMessageParser
func ConstructAckMesageParser() *AckMessageParser {
	ackExpr, _ := regexp.Compile(`ACK\s?<(.*)>`)
	return &AckMessageParser{
		AckExpression: ackExpr,
	}
}

//AckMessageParser parse ack message
type AckMessageParser struct {
	AckExpression *regexp.Regexp
}

//Parse parse ack message
func (parser *AckMessageParser) Parse(rMessage *message.RawMessage) interface{} {
	if parser.AckExpression.Match(rMessage.RawData) {
		if value := parser.parseStringValue(rMessage.RawData); value != "" {
			return &message.AckMessage{
				Identity:    rMessage.Identity(),
				Value:       value,
				MessageType: rMessage.MessageType,
			}
		}
	}
	logger.Logger().WriteToLog(logger.Error, "[AckNackMessageParser] Cant parse ", rMessage.MessageType, " message. Message : ", string(rMessage.RawData), "Serial : ", rMessage.SerialNumber)
	return nil
}

func (parser *AckMessageParser) parseStringValue(rawData []byte) string {
	subMatch := parser.AckExpression.FindAllStringSubmatch(string(rawData), -1)
	if len(subMatch) == 0 {
		return ""
	}
	value := strings.ReplaceAll(subMatch[0][1], "<", "")
	value = strings.ReplaceAll(value, ">", "")
	return strings.TrimSpace(value)
}
