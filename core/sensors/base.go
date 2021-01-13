package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
	"time"
)

//Base base sensor
type Base struct {
	createdAt  time.Time
	symbol     string
	TrigeredBy byte
}

//Trigered triger for sensor
func (b *Base) Trigered(rData map[string]interface{}, posibleReasons map[byte]byte) {
	r, f := rData[core.Reason]
	if !f {
		b.TrigeredBy = byte(0)
	}
	reason := columns.BuildReasonColumn(r)
	if reason == nil {
		b.TrigeredBy = byte(0)
	}
	if v, f := posibleReasons[reason.Code]; !f {
		b.TrigeredBy = byte(0)
	} else {
		b.TrigeredBy = v
	}
}

//Symbol ...
func (b *Base) Symbol() string {
	return b.symbol
}
