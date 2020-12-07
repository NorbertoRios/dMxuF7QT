package test

import (
	"genx-go/configuration"
	"genx-go/core/filter"
	"genx-go/core/location"
	"genx-go/core/request"
	"genx-go/message"
	"genx-go/parser"
	"genx-go/test/mock"
	"strings"
	"testing"
)

var factory = message.CounstructRawMessageFactory()

func locationMessage(t *testing.T) *message.Message {
	param24 := "24=1.7.13.36.3.4.23.65.10.17.11.79.46.44.43.82.152.41.48.56.70.77.93.130;"
	param24 = strings.ReplaceAll(strings.Split(param24, "=")[1], ";", "")
	param24Columns := strings.Split(param24, ".")
	file := &mock.File{FilePath: "../ReportConfiguration.xml"}
	xmlProvider := configuration.ConstructXMLProvider(file)
	config := configuration.ConstructReportConfiguration(xmlProvider)
	fields := config.GetFieldsByIds(param24Columns)
	parser := &parser.GenxBinaryReportParser{
		ReportFields: fields,
	}
	if parser == nil {
		t.Error("Parser is nil")
	}
	packet := []byte{0x33, 0x36, 0x30, 0x30, 0x32, 0x39, 0x39, 0x36, 0x00, 0x00, 0x00, 0x57, 0xc6, 0x00, 0x18, 0x5e, 0xc6, 0x4a, 0x48, 0x0a, 0x7b, 0x57, 0x16, 0x08, 0x11, 0xc3, 0xac, 0x00, 0x04, 0xf5, 0x5c, 0x06, 0x06, 0x01, 0x39, 0x01, 0x02, 0x00, 0x99, 0x00, 0x00, 0x63, 0x9f, 0x2a, 0x73, 0xf5, 0x01, 0x80, 0x2b, 0xc7, 0x38, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0C, 0x2A, 0x0D, 0x2A, 0x0F, 0x23, 0x0F, 0xE3}
	rm := factory.BuildRawMessage(packet)
	if err != nil {
		t.Error("[TestAckPreParsing] Error in construct new raw message")
	}
	result, _ := parser.Parse(rm)
	if result == nil {
		t.Error("Parsed result is null")
	}
	return result[0]
}

func TestLocationRequestLogic(t *testing.T) {
	req := &request.BaseRequest{
		FacadeCallbackID: "test_callback",
		Identity:         "genx_000003870006",
		TTL:              300,
	}
	device := mock.NewDevice()
	locationRequest := location.New(device)
	locationRequest.NewRequest(req)
	task := locationRequest.Task()
	f := filter.NewObserversFilter(device.Observable())
	if len(f.Extract(task)) == 0 {
		t.Error("Observers count unexpected 0")
	}
	deviceMessage := locationMessage(t)
	device.MessageArrived(deviceMessage)
	task = locationRequest.Task()
	f = filter.NewObserversFilter(device.Observable())
	if len(f.Extract(task)) != 0 {
		t.Error("Observers count more than 0")
	}
}
