package main

import (
	"genx-go/logger"
)

func main() {
	service := NewGenxService()
	logger.Logger().WriteToLog(logger.Info, "[Main] Genx service created")
	service.Run()
}
