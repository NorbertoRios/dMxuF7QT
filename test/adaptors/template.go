package adaptors

type ITemplate interface {
}

type IgnitionTemplate struct {
	State byte
}

type GPSTemplate struct {
	Latitude    float32
	Longitude   float32
	Speed       float32
	Heading     float32
	Satellites  byte
	GpsValidity byte
}

type GPIOTemplate struct {
	Switch1 byte
	Switch2 byte
	Switch3 byte
	Switch4 byte
}

type RelaysTemplate struct {
	Relay1 byte
	Relay2 byte
	Relay3 byte
	Relay4 byte
}

type PowerTemplate struct {
	Power int32
	State string
}

type LocIDTemplate struct {
	LockID uint32
}

type TripTemplate struct {
	Odometer int32
}

type NetworkTemplate struct {
	RSSI int8
	CSID int32
}

type TimeTemplate struct {
	TimeStamp string
	EventTime string
}

type FirmwareTemplate struct {
	Version string
}

func NewTemplate(_templates map[string]ITemplate) *Template {
	return &Template{
		Templates: _templates,
	}
}

type Template struct {
	Templates map[string]ITemplate
}

func (t *Template) Template(name string) ITemplate {
	if t, f := t.Templates[name]; f {
		return t
	}
	return nil
}
