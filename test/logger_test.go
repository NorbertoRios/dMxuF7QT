package test

import (
	"genx-go/logger"
	"testing"
)

func TestLoggers(t *testing.T) {
	logger.BuildLogger()
	for i := 0; i < 100; i++ {
		logger.Error("ERROR LOG ", i)
		logger.Info("INFO LOG ", i*2)
	}
	logger.Fatal("FATAL LOG ", "The end")
}
