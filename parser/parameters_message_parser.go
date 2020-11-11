package parser

import (
	"fmt"
	"genx-go/logger"
	"genx-go/message"
	"regexp"
	"strings"
)

//ConstructParametersMessageParser returns parameters message parser
func ConstructParametersMessageParser() *ParametersMessageParser {
	allParamExpr, _ := regexp.Compile(`(?s)ALL-PARAMETERS*.(.*=.*;)`)
	paramExpr, _ := regexp.Compile(`(?s)PARAMETERS*.(.*=.*;)`)
	return &ParametersMessageParser{
		AllParametersExpr: allParamExpr,
		ParametersExpr:    paramExpr,
	}
}

//ParametersMessageParser parse parameters message
type ParametersMessageParser struct {
	AllParametersExpr *regexp.Regexp
	ParametersExpr    *regexp.Regexp
}

//Parse parse message
func (parser *ParametersMessageParser) Parse(rawMessage *message.RawMessage) *message.ParametersMessage {
	var expr *regexp.Regexp
	if parser.ParametersExpr.Match(rawMessage.RawData) {
		expr = parser.ParametersExpr
	}
	if parser.AllParametersExpr.Match(rawMessage.RawData) {
		expr = parser.AllParametersExpr
	}
	if expr == nil {
		logger.Logger().WriteToLog(logger.Error, fmt.Sprintf("[ParametersMessageParser] Can't match parameters message. %v ", string(rawMessage.RawData)))
		return nil
	}

	parameters := expr.FindAllStringSubmatch(string(rawMessage.RawData), -1)[0][1]
	re := regexp.MustCompile(`(\n)|(\r\n)`)
	paramsArr := re.Split(parameters, -1)
	values := make(map[string]string, 0)
	for _, cfg := range paramsArr {
		cfgName := strings.Split(cfg, "=")[0]
		if len(cfgName) == 0 {
			continue
		}
		values[cfgName] = cfg
	}
	return &message.ParametersMessage{
		Identity:    rawMessage.Identity(),
		MessageType: rawMessage.MessageType,
		Parameters:  values,
	}
}
