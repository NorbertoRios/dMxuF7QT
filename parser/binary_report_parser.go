package parser

import (
	"fmt"
	"genx-go/configuration"
	"genx-go/core"
	"genx-go/logger"
	"genx-go/message"
	"genx-go/types"
)

//NewGenxBinaryReportParser ...
func NewGenxBinaryReportParser(param24 []string) *GenxBinaryReportParser {
	file := types.NewFile("/config/initialize/ReportConfiguration.xml")
	provider := configuration.ConstructXMLProvider(file)
	config := configuration.ConstructReportConfiguration(provider)
	return &GenxBinaryReportParser{
		ReportFields: config.GetFieldsByIds(param24),
	}
}

//GenxBinaryReportParser parse message from genx
type GenxBinaryReportParser struct {
	ReportFields []*configuration.Field
}

//Parse parser for location message
func (parser *GenxBinaryReportParser) Parse(rawMessage *message.RawMessage) interface{} {
	defer func() {
		if r := recover(); r != nil {
			logger.Logger().WriteToLog(logger.Error, "panic:Recovered in ParseLocationMessage:", r)
		}
	}()
	messages := make([]*message.Message, 0)
	if len(rawMessage.RawData) == 0 {
		logger.Logger().WriteToLog(logger.Error, "[ParseLocationMessage] Cant parse empty packet")
		return message.NewEmptyLocationMessage()
	}
	position := 0
	firstLen := 0
	for position < len(rawMessage.RawData) && position+firstLen < len(rawMessage.RawData) {
		data := make(map[string]interface{})
		for _, f := range parser.ReportFields {
			value, count := parser.readField(rawMessage, position, f)
			position = position + count
			data[f.Name] = value
		}
		data[core.RawData] = rawMessage.RawData[firstLen:position]
		if firstLen == 0 {
			firstLen = position - 1
		}
		msg := message.BuildMessage(data, rawMessage.MessageType, rawMessage.Identity())
		messages = append(messages, msg)
	}
	return message.NewLocationMessage(messages, parser.buildAck(rawMessage.RawData))
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
