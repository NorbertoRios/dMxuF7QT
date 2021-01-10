package test

import (
	"fmt"
	"genx-go/core/sensors"
	"genx-go/message"
	"genx-go/parser"
	"testing"
)

// 1WIRE:10/Present
// 1:10B8993E01080099 TS-65.300003
// 2:102E5A080000006A TS-56.299999
// 3:109F6108000000AE TS-60.799999
// 4:288D19BD01000096 TS-66.199997
// 5:01AB16430F00004E ID
// 000003912835 3912835

// 1WIRE:10/NoDevice
// 1:0000000000000000
// 2:0000000000000000
// 3:0000000000000000
// 4:0000000000000000
// 5:0000000000000000
// 000003912835 3912835

// 1WIRE:10/Present
// 1:10B8993E01080099 TS-65.300003
// 2:102E5A080000006A TS-56.299999
// 3:109F6108000000AE TS-60.799999
// 4:0000000000000000
// 5:01AB16430F00004E ID
// 000003912835 3912835

func TestPresent1WireMessage3TempSensors(t *testing.T) {
	packet := []byte("1WIRE:10/Present\n1:10B8993E01080099 TS-65.300003\n2:102E5A080000006A TS-56.299999\n3:109F6108000000AE TS-60.799999\n4:0000000000000000\n5:01AB16430F00004E ID\n000003912835 3912835")
	rm := factory.BuildRawMessage(packet)
	parser := parser.BuildOneWireMessageParser()
	msg := parser.Parse(rm)
	if len(msg.(*message.Message).Sensors) != 4 {
		t.Error("Error in count of sensors")
	}
	checkTempSensors(msg.(*message.Message).Sensors[0], byte(1), "10B8993E01080099", float32(65.300003), t)
	checkTempSensors(msg.(*message.Message).Sensors[1], byte(2), "102E5A080000006A", float32(56.299999), t)
	checkTempSensors(msg.(*message.Message).Sensors[2], byte(3), "109F6108000000AE", float32(60.799999), t)
	checkIButtonSensor(msg.(*message.Message).Sensors[3], int32(0xF00004E), t)
}

func TestNoDevice1WireMessageParsing(t *testing.T) {
	packet := []byte("1WIRE:10/NoDevice\n1:0000000000000000\n2:0000000000000000\n3:0000000000000000\n4:0000000000000000\n5:0000000000000000\n000003912835 3912835")
	rm := factory.BuildRawMessage(packet)
	parser := parser.BuildOneWireMessageParser()
	msg := parser.Parse(rm)
	if msg.(*message.Message).Sensors != nil {
		t.Error("Error. No device")
	}
}

func TestPresent1WireMessageParsing(t *testing.T) {
	packet := []byte("1WIRE:10/Present\n1:10B8993E01080099 TS-65.300003\n2:102E5A080000006A TS-56.299999\n3:109F6108000000AE TS-60.799999\n4:288D19BD01000096 TS-66.199997\n5:01AB16430F00004E ID\n000003912835 3912835")
	rm := factory.BuildRawMessage(packet)
	parser := parser.BuildOneWireMessageParser()
	msg := parser.Parse(rm)
	if len(msg.(*message.Message).Sensors) != 5 {
		t.Error("Error in count of sensors")
	}
	checkTempSensors(msg.(*message.Message).Sensors[0], byte(1), "10B8993E01080099", float32(65.300003), t)
	checkTempSensors(msg.(*message.Message).Sensors[1], byte(2), "102E5A080000006A", float32(56.299999), t)
	checkTempSensors(msg.(*message.Message).Sensors[2], byte(3), "109F6108000000AE", float32(60.799999), t)
	checkTempSensors(msg.(*message.Message).Sensors[3], byte(4), "288D19BD01000096", float32(66.199997), t)
	checkIButtonSensor(msg.(*message.Message).Sensors[4], int32(0xF00004E), t)
}

func checkIButtonSensor(s sensors.ISensor, driverID int32, t *testing.T) {
	sensor, f := s.(*sensors.IButton)
	if !f {
		t.Error("Cant cast sensor to IButtonSensor")
	}
	if sensor.BtnID != driverID {
		t.Error(fmt.Sprintf("DriverID=%v doesnt equal %v", sensor.BtnID, driverID))
	}
}

func checkTempSensors(s sensors.ISensor, id byte, imei string, value float32, t *testing.T) {
	sensor, f := s.(*sensors.TemperatureSensor)
	if !f {
		t.Error("Cant cast sensor to TemperatureSensor")
	}
	if sensor.ID != id {
		t.Error(fmt.Sprintf("ID=%v doesnt equal %v", sensor.ID, id))
	}
	if sensor.Imei != imei {
		t.Error(fmt.Sprintf("IMEI=%v doesnt equal %v", sensor.Imei, imei))
	}
	if sensor.Value != value {
		t.Error(fmt.Sprintf("Value=%v doesnt equal %v", sensor.Value, value))
	}
}
