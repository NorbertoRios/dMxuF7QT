package logger

import (
	"fmt"
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

var staticLogger *Logger

//BuildLogger returns new logger
func BuildLogger() {
	log.SetFlags(log.Ldate | log.Ltime)
	errorOut := &lumberjack.Logger{
		Filename:   "logs/error.log",
		MaxSize:    500, // megabytes
		MaxBackups: 6,
		MaxAge:     7,    //days
		Compress:   true, // disabled by default
	}
	infoOut := &lumberjack.Logger{
		Filename:   "logs/info.log",
		MaxSize:    500, // megabytes
		MaxBackups: 6,
		MaxAge:     7,    //days
		Compress:   true, // disabled by default
	}
	fatalOut := &lumberjack.Logger{
		Filename:   "logs/fatal.log",
		MaxSize:    500, // megabytes
		MaxBackups: 6,
		MaxAge:     7,    //days
		Compress:   true, // disabled by default
	}
	staticLogger = &Logger{
		ErrorOutput: errorOut,
		InfoOutput:  infoOut,
		FatalOutput: fatalOut,
	}
}

//Logger logger
type Logger struct {
	ErrorOutput *lumberjack.Logger
	InfoOutput  *lumberjack.Logger
	FatalOutput *lumberjack.Logger
}

//Error for log errors
func Error(content ...interface{}) {
	logContent := fmt.Sprintln(content...)
	log.SetOutput(staticLogger.ErrorOutput)
	log.Println(logContent)
}

//Info for log info
func Info(content ...interface{}) {
	logContent := fmt.Sprintln(content...)
	log.SetOutput(staticLogger.InfoOutput)
	log.Println(logContent)
}

//Fatal for log fatal
func Fatal(content ...interface{}) {
	logContent := fmt.Sprintln(content...)
	log.SetOutput(staticLogger.FatalOutput)
	log.Fatalf(logContent)
}
