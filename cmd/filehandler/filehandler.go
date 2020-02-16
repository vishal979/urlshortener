package filehandler

import (
	"fmt"
	"log"
	"os"
)

var LogsFile *os.File

func Open() *os.File {

	LogsFile, err := os.OpenFile("logs/server.logs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(LogsFile)
	WriteLogs("logs file successfully opened")
	return LogsFile
}

func WriteLogs(logs string) {
	// LogsFile.Write([]byte(logs))
	log.Println(logs)
}

func ErrorHandling(custom string, err error) {
	error := fmt.Sprintf("", err)
	logs := "error while serving" + error
	WriteLogs(logs)
}

func CloseFile() {
	LogsFile.Close()
}
