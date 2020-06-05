package parser

import (
	"fmt"
	"genx-go/configuration"
	"genx-go/core"
	"genx-go/message"
	"log"
	"strings"
)

//GenxBinaryReportParser parse message from genx
type GenxBinaryReportParser struct {
	reportFields []*configuration.Field
}

//BuildGenxBinaryReportParser returns new report parser
func BuildGenxBinaryReportParser(param24 string, reportConfiguration *configuration.ReportConfiguration) *GenxBinaryReportParser {
	param24 = strings.ReplaceAll(param24, ";", "")
	param24_columns := strings.Split(param24, ".")
	fields := reportConfiguration.GetFieldsByIds(param24_columns)
	return &GenxBinaryReportParser{
		reportFields: fields,
	}
}


//Parse parser for location message
func (parser *GenxBinaryReportParser) Parse(rawMessage *message.RawMessage) ([]*message.Message, string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic:Recovered in ParseLocationMessage:", r)
		}
	}()
	messages := make([]*message.Message, 0)
	if len(rawMessage.RawData) == 0 {
		log.Println("[ParseLocationMessage] Cant parse empty  packet")
		return nil, ""
	}
	position := 0
	firstLen := 0
	log.Println(len(rawMessage.RawData))
	for position < len(rawMessage.RawData) && position+firstLen < len(rawMessage.RawData) {
		data := make(map[string]interface{})
		for _, f := range parser.reportFields {
			value, count := parser.readField(rawMessage, position, f)
			position = position + count
			data[f.Name] = value
		}
		data[core.RawData] = rawMessage.RawData[firstLen:position]
		if firstLen == 0 {
			firstLen = position - 1
		}
		msg := message.BuildMessage(data, rawMessage.MessageType)
		messages = append(messages, msg)
	}
	return messages, parser.buildAck(rawMessage.RawData)
}

func (parser *GenxBinaryReportParser) buildAck(packet []byte) string {
	return fmt.Sprintf("UDPACK %v", len(packet))
}

func (*GenxBinaryReportParser) readField(message *message.RawMessage, position int, field *configuration.Field) (interface{}, int) {
	if field.Size == -1 {
		return message.String(position)
	}
	return message.Value(position, field)
}
