package test

import (
	"fmt"
	"genx-go/core/immobilizer/request"
	"genx-go/core/immobilizer/task"
	"genx-go/core/usecase"
	"genx-go/message"
	"genx-go/test/mock"
	"testing"
	"time"
)

func TestImmobilizerLogic(t *testing.T) {
	factory := message.Factory()
	req := &request.ChangeImmoStateRequest{
		SafetyOption: true,
		State:        "armed",
		Trigger:      "high",
	}
	req.Port = "OUT0"
	req.Identity = "genx_000003870006"
	req.FacadeCallbackID = "testCallback_-1"
	device := mock.NewDevice()
	usecase.NewImmobilizerUseCase(device, req).Launch()
	go func() {
		for i := 0; i < 10; i++ {
			req.FacadeCallbackID = fmt.Sprintf("%v_%v", req.FacadeCallbackID, i)
			usecase.NewImmobilizerUseCase(device, req).Launch()
		}
	}()
	rMessage := factory.BuildRawMessage([]byte("000003870006 ACK < SETRELAYDRIVE1ON SERIALFILTER 000003870006;BACKUPNVRAM>"))
	usecase.NewMessageArrivedUseCase(device, rMessage).Launch()
	rMessage = factory.BuildRawMessage([]byte("MODEL:GNX-5P\nSN:000003870006\nFW:G699.06.78kX 12:59:45 May 25 2012\nHW:656, HWOPTID:0016\nIMEI:357852038210210\nMVER:07.60.00\nGVER:7.03 (45969) 00040007\nOn:40:09:45(48)\nIgn-ON,Volt-12131,Switch-1001,Relay-0001,A2D-4150\nV-12131/12133/12133 Temp-319\nMallocs 0\nCRC:6a10.6ce4.88e8.88e8.60c0.948.537a.537a.9bfd.9bfd.\n000003870006 3870006"))
	time.Sleep(100 * time.Millisecond)
	usecase.NewMessageArrivedUseCase(device, rMessage).Launch()
	i := device.Immobilizer(1, "high")
	if i == nil {
		t.Error("Unexpected nil immobilizer")
	}
	doneTask := i.Tasks().Front().Value.(*task.DoneImmoTask)
	if doneTask.Task.Request().(*request.ChangeImmoStateRequest).State != "armed" {
		t.Error("Unexpected immobilizer state. Current state is ", i.State(device))
	}
	if i.Tasks().Len() == 0 {
		t.Error("Immobilizer task length is 0 ")
	}
	if i.CurrentTask() != nil {
		t.Error("Current immobilizer task is not nil ")
	}
	doneTasks := 0
	cancelTasks := 0
	for tsk := i.Tasks().Front(); tsk != nil; tsk = tsk.Next() {
		if _, v := tsk.Value.(*task.DoneImmoTask); v {
			doneTasks++
		}
		if _, v := tsk.Value.(*task.CanceledImmoTask); v {
			cancelTasks++
		}
	}
	if cancelTasks != 10 || doneTasks != 1 {
		t.Error(fmt.Sprintf("Unexpected done or cancel tasks count. Should (done: 1, cancel: 10). Current(done: %v, cancel: %v)", doneTasks, cancelTasks))
	}
}
