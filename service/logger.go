package service

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"os"
	"time"
)

func GetLogger(spec string) (*log.Logger, error) {
	file, err := makeFile(getCurrentTime())
	if err != nil {
		return nil, err
	}
	// saving the file pointer to close it in the further
	fileTemp := file
	debugLogger := setLogger(file)
	setTask(spec, func() {
		file, err := makeFile(getCurrentTime())
		if err != nil {
			debugLogger.Println("create new file error")
		} else {
			debugLogger = setLogger(file)
			// refresh the file pointer
			fileTemp.Close()
			fileTemp = file
		}
	})
	return debugLogger, nil
}

// make a file for saving the logs
func makeFile(filename string) (*os.File, error) {
	fileName := fmt.Sprintf("./%s.txt", filename)
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// save the logs with the designated file
func setLogger(logFile *os.File) *log.Logger {
	logger := log.New(logFile, "[lian_you]", log.Llongfile)
	return logger
}

// set the timed task to refresh the file every day
func setTask(spec string, task func()) {
	c := cron.New()
	c.AddFunc(spec, task)
	c.Start()
}

func getCurrentTime() string {
	return time.Now().Format("2006-01-02")
}
