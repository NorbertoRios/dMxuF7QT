package test

import (
	"genx-go/logger"
	"testing"
)

func TestLoggers(t *testing.T) {
	
	for i := 0; i < 100; i++ {
		logger.Logger().WriteToLog(logger.Error, "ERROR LOG ", i)
		logger.Logger().WriteToLog(logger.Info, "INFO LOG ", i*2)
	}
}
