package test

import (
	"genx-go/core/immobilizer/request"
	"genx-go/core/immobilizer/task"
	bRequest "genx-go/core/request"
	"genx-go/message"
	"genx-go/parser"
	"genx-go/test/mock"
	"testing"
)

var factory = message.CounstructRawMessageFactory()

func TestImmobilizerLogic(t *testing.T) {
	req := &request.ChangeImmoStateRequest{
		SafetyOption: true,
		State:        "armed",
		Trigger:      "high",
	}
	req.Port = "OUT0"
	req.Identity = "genx_000003870006"
	req.FacadeCallbackID = "testCallback"
	device := mock.NewDevice()
	out := &bRequest.OutputNumber{Data: req.Port}
	immo := device.Immobilizer(out.Index(), req.Trigger)
	immo.NewRequest(req)
	packet := []byte("000003870006 ACK < SETRELAYDRIVE1ON SERIALFILTER 000003870006;BACKUPNVRAM>")
	rm := factory.BuildRawMessage(packet)
	p := parser.ConstructAckMesageParser()
	ackMessage := p.Parse(rm)
	device.MessageArrived(ackMessage)
	packet = []byte("MODEL:GNX-5P\nSN:000003870006\nFW:G699.06.78kX 12:59:45 May 25 2012\nHW:656, HWOPTID:0016\nIMEI:357852038210210\nMVER:07.60.00\nGVER:7.03 (45969) 00040007\nOn:40:09:45(48)\nIgn-ON,Volt-12131,Switch-1001,Relay-0001,A2D-4150\nV-12131/12133/12133 Temp-319\nMallocs 0\nCRC:6a10.6ce4.88e8.88e8.60c0.948.537a.537a.9bfd.9bfd.\n000003870006 3870006")
	rm = factory.BuildRawMessage(packet)
	hp := parser.BuildGenxHardwareMessageParser()
	hwMessage := hp.Parse(rm)
	device.MessageArrived(hwMessage)
	i := device.Immobilizer(1, "high")
	if i == nil {
		t.Error("Unexpected nil immobilizer")
	}
	doneTask := i.Tasks().Front().Value.(*task.DoneImmoTask)
	if doneTask.Task.Request().(*request.ChangeImmoStateRequest).State != "armed" {
		t.Error("Unexpected immobilizer state. Current state is ", i.State())
	}
	if i.Tasks().Len() == 0 {
		t.Error("Immobilizer task length is 0 ")
	}
	if i.CurrentTask() != nil {
		t.Error("Current immobilizer task is not nil ")
	}
	for tsk := i.Tasks().Front(); tsk != nil; tsk = tsk.Next() {
		if _, v := tsk.Value.(*task.DoneImmoTask); !v {
			t.Error("Unexpected task type")
		}
	}
}
