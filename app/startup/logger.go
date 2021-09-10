package startup

import (
	"log"
	"os"
)

var (
    InfoLog    *log.Logger
    ErrorLog   *log.Logger
)

type Log struct{}

func (l *Log) InitialLog() {
	// log to custom file
	logPath := "./" + os.Getenv("LOGGING_DIR") + "/"
	logInfoFile := logPath + os.Getenv("INFO_LOG_PATH")
	logErrorFile := logPath + os.Getenv("ERROR_LOG_PATH")

	InfoLog = log.New(getLogfile(logInfoFile), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLog = log.New(getLogfile(logErrorFile), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func getLogfile(file string) *os.File{
	logFile, err := os.OpenFile(file, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// defer logFile.Close()

	return logFile
}
