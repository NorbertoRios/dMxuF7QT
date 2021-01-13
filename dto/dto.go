package dto

import (
	"encoding/json"
	"genx-go/core/sensors"
)

//IMessage ...
type IMessage interface {
	SetValue(key string, value interface{})
	AppendRange(data map[string]interface{})
}

//NewMessage returns new struct of  message
func NewMessage() *Message {
	return &Message{Data: make(map[string]interface{})}
}

//Message struct for parsed messages
type Message struct {
	SID                uint64            `json:"sid"`
	TemperatureSensors []sensors.ISensor `json:"ts,omitempty"`
	Data               map[string]interface{}
}

//SetValue to Data field
func (m *Message) SetValue(key string, value interface{}) {
	m.Data[key] = value
}

//AppendRange append data fields to current Data
func (m *Message) AppendRange(data map[string]interface{}) {
	for k, v := range data {
		m.SetValue(k, v)
	}
}

//UnMarshalMessage given string to Message struct
func UnMarshalMessage(str string) (*Message, error) {
	message := &Message{}
	err := json.Unmarshal([]byte(str), message)
	if err != nil {
		return &Message{}, err
	}
	return message, err
}
