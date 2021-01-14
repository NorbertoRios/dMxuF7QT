package config

import (
	"genx-go/core/configuration/request"
	"genx-go/core/configuration/task"
	"genx-go/core/usecase"
	"genx-go/message"
	"genx-go/test/mock"
	"testing"
)

func TestConfigurationLogic(t *testing.T) {
	factory := message.Factory()
	req := &request.ConfigurationRequest{
		Config: []string{
			"1=One;",
			"2=Two;",
			"3=Three;",
			"4=Four;",
			"5=Five;",
			"6=Six;",
			"7=Seven;",
			"8=Eight;",
			"9=Nine;",
			"10=Ten;",
		},
	}
	commandAcks := []string{
		"000003870006 ACK <SETPARAMVERIFY;1=One;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;2=Two;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;3=Three;ENDPARAM;BACKUPNVRAM;>",
	}
	req.Identity = "genx_000003870006"
	req.FacadeCallbackID = "testCallback"
	device := mock.NewDevice()
	usecase.NewConfigUseCase(device, req).Launch()
	for i := 0; i < 3; i++ {
		usecase.NewMessageArrivedUseCase(device, factory.BuildRawMessage([]byte(commandAcks[i]))).Launch()
	}
	ct := device.Configuration().CurrentTask().(*task.ConfigTask)
	sentCount := 0
	notSentCount := 0
	for cmd := ct.ConfigCommands().Front(); cmd != nil; cmd = cmd.Next() {
		if cmd.Value.(*request.Command).State() {
			sentCount++
		} else {
			notSentCount++
		}
	}
	if sentCount != 3 {
		t.Error("Unexpected sent commands count. ", sentCount)
	}

	if notSentCount != 7 {
		t.Error("Unexpected not sent commands count. ", notSentCount)
	}
	newReq := &request.ConfigurationRequest{
		Config: []string{
			"1=One;",
			"2=Two;",
			"3=Three;",
			"4=FFFour;",
			"5=FFFive;",
			"6=SSSix;",
			"7=SSSeven;",
			"8=EEEight;",
			"9=NNNine;",
			"10=TTTen;",
		},
	}
	commandAcks = []string{
		"000003870006 ACK <SETPARAMVERIFY;1=One;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;2=Two;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;3=Three;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;4=FFFour;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;5=FFFive;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;6=SSSix;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;7=SSSeven;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;8=EEEight;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;9=NNNine;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;10=TTTen;ENDPARAM;BACKUPNVRAM;>",
	}
	usecase.NewConfigUseCase(device, newReq).Launch()
	for i := 0; i < len(commandAcks); i++ {

		usecase.NewMessageArrivedUseCase(device, factory.BuildRawMessage([]byte(commandAcks[i]))).Launch()
	}
	doneTasks := 0
	canceledTasks := 0
	for tsk := device.Configuration().Tasks().Front(); tsk != nil; tsk = tsk.Next() {
		switch tsk.Value.(type) {
		case *task.CanceledConfigTask:
			{
				doneTasks++
			}
		case *task.DoneConfigTask:
			{
				canceledTasks++
			}
		default:
			{
				t.Error("Unexpected task type")
			}
		}
	}
	if doneTasks != 1 || canceledTasks != 1 {
		t.Error("Unexpected done/canceled tasks count")
	}
}
