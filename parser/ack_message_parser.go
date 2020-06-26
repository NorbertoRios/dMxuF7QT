package parser

import (
	"genx-go/message"
	"genx-go/utils"
	"log"
	"regexp"
	"strings"
)

//ConstructAckMesageParser returns new AckMessageParser
func ConstructAckMesageParser() *AckNackMessageParser {
	ackExpr, _ := regexp.Compile(`\d+ ACK(.*)`)
	return &AckNackMessageParser{
		AckExpression: ackExpr,
	}
}

//AckNackMessageParser parse ack message
type AckNackMessageParser struct {
	AckExpression *regexp.Regexp
}

//Parse parse ack message
func (parser *AckNackMessageParser) Parse(rMessage *message.RawMessage) *message.AckMessage {
	if parser.AckExpression.Match(rMessage.RawData) {
		if value := parser.parseStringValue(rMessage.RawData); value != "" {
			sUtils := &utils.StringUtils{Data: rMessage.SerialNumber}
			return &message.AckMessage{
				Identity:    sUtils.Identity(),
				Value:       value,
				MessageType: rMessage.MessageType,
			}
		}
	}
	log.Println("[AckNackMessageParser] Cant parse ", rMessage.MessageType, " message. Message : ", string(rMessage.RawData), "Serial : ", rMessage.SerialNumber)
	return nil
}

func (parser *AckNackMessageParser) parseStringValue(rawData []byte) string {
	subMatch := parser.AckExpression.FindAllStringSubmatch(string(rawData), -1)
	if len(subMatch) == 0 {
		return ""
	}
	value := strings.ReplaceAll(subMatch[0][1], "<", "")
	value = strings.ReplaceAll(value, ">", "")
	return strings.TrimSpace(value)
}
