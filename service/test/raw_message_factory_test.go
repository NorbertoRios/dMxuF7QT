package test

import (
	"fmt"
	"genx-go/message"
	"testing"
)

var factory = message.CounstructRawMessageFactory()

func TestBinaryReporPreParsing(t *testing.T) {
	packet := []byte{0x33, 0x34, 0x36, 0x31, 0x34, 0x37, 0x30, 0x39, 0x00, 0x00, 0x02, 0x78, 0x1C, 0x00, 0x29, 0x5E, 0xC2, 0x83, 0xC3, 0x99, 0xC2, 0x8F, 0x03, 0x28, 0xC2, 0x94, 0xC2, 0x9B, 0x19, 0xC3, 0xB7, 0xC2, 0x82, 0xC3, 0xA5, 0x01, 0x19, 0x61, 0xC2, 0xBB, 0x00, 0x3C, 0x00, 0xC3, 0xB2, 0x01, 0x00, 0x0C, 0xC2, 0xA7, 0x00, 0x00, 0xC3, 0x89, 0x2F, 0x37, 0x77, 0x00}
	res := factory.BuildRawMessage(packet)
	checkRawMessage("TestBinaryReporPreParsing", res, "34614709", "breport", t)
}

func TestAckPreParsing(t *testing.T) {
	packet := []byte("000003912835 ACK <SETPARAM 133=59; ENDPARAM;BACKUPNVRAM;>")
	res := factory.BuildRawMessage(packet)
	checkRawMessage("TestAckPreParsing", res, "000003912835", "ack", t)
}

func TestAckPreParsing2(t *testing.T) {
	packet := []byte("000003912835 ACK <SETPARAM 133=59; ENDPARAM;BACKUPNVRAM;>")
	res := factory.BuildRawMessage(packet)
	checkRawMessage("TestAckPreParsing2", res, "000003912835", "ack", t)
}

func TestNackPreParsing(t *testing.T) {
	packet := []byte("000003912835 NAK- ERROR <ARAM BLA=dfgdfg; END>")
	res := factory.BuildRawMessage(packet)
	checkRawMessage("TestNackPreParsing", res, "000003912835", "nack", t)
}

func TestAllParameters(t *testing.T) {
	packet := []byte("ALL-PARAMETERS\n500=GFM121232;\n501=259200;\n503=100;\n505=0;\n2=300;\n4=14400.0;\n5=2;\n6=1;\n7=35.197.10.57.0.0;\n8=300;\n9=3912835;\n10=;\n11=;\n12=;\n13=;\n14=0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0;\n15=;\n16=0.0.0.0;\n17=;\n18=WWW.KYIVSTAR.NET;\n19=;\n20=;\n24=1.7.13.36.3.4.23.65.10.17.11.79.46.44.43.82.152.41.48.0.0.0.0.0.0.0.0.0.0.0;\n25=0;\n26=2;\n27=1;\n28=;\n29=;\n30=34.83.8.55.37.22;\n31=;\n32=;\n33=maps.google.com/maps?q=%%LAT--%%+%%LON--%%;\n34=0.0.0.0.0.0.0.0;\n35=;\n36=0;\n37=75.10.252.41.99.109.207.5;\n38=;\n39=1.1;\n")
	res := factory.BuildRawMessage(packet)
	checkRawMessage("TestAllParameters", res, "3912835", "param", t)
}

func TestPoll(t *testing.T) {
	packet := []byte("3912835,11/25 12:32 GMT,POLL\nIdling for 339min \nhttp://maps.google.com/maps?q=48.746658+37.590587\nIgnition ON,Odo:40.3km,12131mV,SN:000003912835\n")
	res := factory.BuildRawMessage(packet)
	checkRawMessage("TestPoll", res, "3912835", "poll", t)
}

func TestDiagHardwareBrief(t *testing.T) {
	packet := []byte("3912835:FW:G699.06.78kX\nHW:656, HWOPTID:0016\nOn:431:50:26(48)\nIgn-ON,Volt-12131,Switch-0000,Relay-0000,A2D-4151\n\n000003912835 3912835\n")
	res := factory.BuildRawMessage(packet)
	checkRawMessage("TestDiagHardwareBrief", res, "3912835", "diag_hardware_brief", t)
}

func TestDiag1Wire(t *testing.T) {
	packet := []byte("1WIRE:10/NoDevice\n1:0000000000000000\n2:0000000000000000\n3:0000000000000000\n4:0000000000000000\n5:0000000000000000\n\n000003912835 3912835\n")
	res := factory.BuildRawMessage(packet)
	checkRawMessage("TestDiag1Wire", res, "3912835", "diag_1wire", t)
}

func TestDiagParam(t *testing.T) {
	packet := []byte("PARAMETERS\n56=;\n000003912835 3912835\n")
	res := factory.BuildRawMessage(packet)
	checkRawMessage("TestDiagParam", res, "3912835", "param", t)
}

func TestDiagCan(t *testing.T) {
	packet := []byte("3912835:J1939(0):250kHz VIN:314-841204       \nCE 0,3,3 00000000 5,5,0,3,1,0,0,0\nAge:0 Spd:0.0 FL:-999.0 RPM:699.6/1145/22729 Odo:1942122.0 CT:190.4 TH:37598.5 \nJ1708 ENABLED(INTERFACE UP)\nBCS:0.0\nMIDS:128,130,137,172,136\nVIN:314-841204       \nAge:1 Spd:0.0 RPM:699.3 Odo:1942913.0 FL:-999.0/-999.0 CT:190.0 FC:846238.6 TH:37598.5\nJ1708DTC\n130 194 3 237 243 255 218 \n136 194 2 4 53 123 \n130 194 6 237 243 255 218 179 255 75 \n\n000003912835 3912835\n")
	res := factory.BuildRawMessage(packet)
	checkRawMessage("TestDiagCan", res, "3912835", "diag", t)
}

func checkRawMessage(methodName string, res *message.RawMessage, shouldSerial string, shouldType string, t *testing.T) {
	if res.MessageType != shouldType {
		t.Error(fmt.Sprintf("[%v] Error in message type, should: %v, current: %v", methodName, shouldType, res.MessageType))
	}
}
