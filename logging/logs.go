package logging

import (
	"io"
	"log"
	"os"
)

var (
	ErrorLogger   *log.Logger
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
)

func init() {
	logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("could not create a log file, aborting!")
	}

	mw := io.MultiWriter(os.Stdout, logFile)

	ErrorLogger = log.New(mw, "ERROR : ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(mw, "WARN : ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(mw, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)

}
