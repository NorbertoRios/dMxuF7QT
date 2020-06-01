package genxmessage

import "regexp"

//ReportMap report type
type ReportMap struct {
	Type string
	Reg  *regexp.Regexp
}
