package config

import (
	"container/list"
	"fmt"
	"genx-go/core/configuration"
	"genx-go/core/configuration/request"
	"genx-go/core/configuration/task"
	"genx-go/message"
	"genx-go/parser"
	"genx-go/test/mock"
	"testing"
)

var factory = message.CounstructRawMessageFactory()

func TestConfigurationLogic(t *testing.T) {
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
	configuration := configuration.NewConfiguration(device, &mock.FacadeClient{})
	configuration.NewRequest(req)
	for i := 0; i < 3; i++ {
		packet := []byte(commandAcks[i])
		rm := factory.BuildRawMessage(packet)
		p := parser.ConstructAckMesageParser()
		ackMessage := p.Parse(rm)
		device.MessageArrived(ackMessage)
	}
	ct := configuration.CurrentTask()
	currentTask := ct.(*task.ConfigTask)
	if currentTask.SentCommands.Len() != 3 {
		t.Error("Unexpected sent commands count")
	}
	if currentTask.Commands.Len() != 7 {
		t.Error("Unexpected commands count")
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
	configuration.NewRequest(newReq)
	ct = configuration.CurrentTask()
	currentTask = ct.(*task.ConfigTask)
	if currentTask.SentCommands.Len() != 3 {
		t.Error("Unexpected sent commands count")
	}
	if currentTask.Commands.Len() != 7 {
		t.Error("Unexpected commands count")
	}
	shouldSentCommands := []string{
		"1=One;",
		"2=Two;",
		"3=Three;",
	}
	shouldNotSentCommands := []string{
		"4=FFFour;",
		"5=FFFive;",
		"6=SSSix;",
		"7=SSSeven;",
		"8=EEEight;",
		"9=NNNine;",
		"10=TTTen;",
	}
	checkCommands(shouldSentCommands, currentTask.SentCommands, t)
	checkCommands(shouldNotSentCommands, currentTask.Commands, t)

	commandAcks = []string{
		"000003870006 ACK <SETPARAMVERIFY;4=FFFour;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;5=FFFive;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;6=SSSix;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;7=SSSeven;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;8=EEEight;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;9=NNNine;ENDPARAM;BACKUPNVRAM;>",
		"000003870006 ACK <SETPARAMVERIFY;10=TTTen;ENDPARAM;BACKUPNVRAM;>",
	}
	for i := 0; i < len(commandAcks); i++ {
		packet := []byte(commandAcks[i])
		rm := factory.BuildRawMessage(packet)
		p := parser.ConstructAckMesageParser()
		ackMessage := p.Parse(rm)
		device.MessageArrived(ackMessage)
	}
	if configuration.CurrentTask() != nil {
		t.Error("Current task sjould be nil")
	}
	var doneTaskCount int
	var canseledtaskCount int
	for _task := configuration.Tasks().Front(); _task != nil; _task = _task.Next() {
		if _, s := _task.Value.(*task.DoneConfigTask); s {
			doneTaskCount++
		} else if _, s = _task.Value.(*task.CanceledConfigTask); s {
			canseledtaskCount++
		} else {
			t.Error("Unexpected task type")
		}
	}
	if doneTaskCount != 1 || canseledtaskCount != 1 {
		t.Error("Unexpected done task count ", doneTaskCount, " or canceled task count ", canseledtaskCount)
	}
}

func checkCommands(commands []string, taskCommands *list.List, t *testing.T) {
	item := taskCommands.Front()
	if len(commands) != taskCommands.Len() {
		t.Error("Task commands len and expected commands len are not equal")
	}
	for i := 0; i < len(commands); i++ {
		var sItem string
		if str, success := item.Value.(string); !success {
			sItem, _ = item.Value.(*list.Element).Value.(string)
		} else {
			sItem = str
		}
		if sItem != commands[i] {
			t.Error(fmt.Sprintf("Task sendes command %v doesnt equal to %v", sItem, commands[i]))
		}
		item = item.Next()
	}
}
