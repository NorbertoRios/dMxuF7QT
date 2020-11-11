package parser

import "genx-go/message"

//IParser parsers interface
type IParser interface {
	Parse(*message.RawMessage) interface{}
}
