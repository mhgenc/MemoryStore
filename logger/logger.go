package logger

import (
	"MemoryStore/config"
	"io"
	"log"
	"os"
)

var LogFile *os.File

//Info log with info message
func Info(args ...interface{}) {
	log.Printf("Info : %s", args...)
}

//Error log with error message
func Error(args ...interface{}) {
	log.Printf("Error : %s", args...)
}

//Fatal log with fatal message
func Fatal(args ...interface{}) {
	log.Printf("Fatal : %s", args...)
}

// ConfigureLogger If the log file is defined, it provides logging to the file in addition to stdout.
func ConfigureLogger() {

	//When working config.txt, logfile would be nil
	if len(config.LOG_FILE) > 0 {
		logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		mw := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(mw)
	}
}
