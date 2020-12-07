package lock

import (
	"genx-go/core/device/interfaces"
	"genx-go/core/filter"
	lockRequest "genx-go/core/lock/request"
	"genx-go/core/lock/task"
	bRequest "genx-go/core/request"
	"genx-go/core/usecase"
	"genx-go/test/mock"
	"testing"
	"time"
)

func TestLockLogic(t *testing.T) {
	exT := time.Now().UTC().Add(1 * time.Minute)
	req := &lockRequest.UnlockRequest{
		ExpirationTime: exT.Format("2006-01-02T15:04:05Z"),
		TimeToPulse:    3,
	}
	req.Port = "OUT0"
	req.Identity = "genx_000003870006"
	device := mock.NewDevice()
	usecase.NewLockUseCase(device, req).Launch()
	packet := []byte("000003870006 ACK < SETRELAYDRIVE1X3FFFFFFF SERIALFILTER 000003870006>")
	usecase.NewMessageArrivedUseCase(device, packet).Launch()
	out := bRequest.OutputNumber{Data: req.Port}
	lock := device.ElectricLock(out.Index())
	if lock.CurrentTask() != nil {
		t.Error("CurrentTask should be nil")
	}
	_, f := lock.Tasks().Front().Value.(*task.DoneElectricLockTask)
	if !f {
		t.Error("Last task should be done task")
	}
	for tsk := lock.Tasks().Front(); tsk != nil; tsk = tsk.Next() {
		oFilter := filter.NewObserversFilter(device.Observable())
		taskObservers := oFilter.Extract(tsk.Value.(interfaces.ITask))
		if len(taskObservers) > 0 {
			t.Error("Task didnt detach all own observers\n ", taskObservers)
		}
	}
}
