package parser

import (
	"fmt"
	"genx-go/message"
	"genx-go/utils"
	"log"
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
func (parser *ParametersMessageParser) Parse(rawMessage *message.RawMessage) interface{} {
	var expr *regexp.Regexp
	if parser.ParametersExpr.Match(rawMessage.RawData) {
		expr = parser.ParametersExpr
	}
	if parser.AllParametersExpr.Match(rawMessage.RawData) {
		expr = parser.AllParametersExpr
	}
	if expr == nil {
		log.Println(fmt.Sprintf("[ParametersMessageParser] Can't match parameters message. %v ", string(rawMessage.RawData)))
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
	sUtils := &utils.StringUtils{Data: rawMessage.SerialNumber}
	return &message.ParametersMessage{
		Identity:    sUtils.Identity(),
		MessageType: rawMessage.MessageType,
		Parameters:  values,
	}
}
