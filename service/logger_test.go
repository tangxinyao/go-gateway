package service

import (
	"testing"
	"time"
)

func TestMakeFile(t *testing.T) {
	currentTime := time.Now().Format("2006-01-02")
	file, err := makeFile(currentTime)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(file.Name()))
}

func TestSetLog(t *testing.T) {
	currentTime := time.Now().Format("2006-01-02")
	file, err := makeFile(currentTime)
	if err != nil {
		t.Fatal(err)
	}
	logger := setLogger(file)
	logger.Println("Hello")
}

func TestSetTask(t *testing.T) {
	// "0 0 12 * * ?"
	currentTime := time.Now().Format("2006-01-02")
	file, err := makeFile(currentTime)
	if err != nil {
		t.Fatal(err)
	}
	setTask("* * * * * ?", func() {
		setLogger(file)
	})
	flag := make(chan bool)
	<-flag
}

func TestGetLogger(t *testing.T) {
	logger, err := GetLogger("0 0 0 * * ?")
	if err != nil {
		t.Fatal(err)
	}
	logger.Println("hello")
	flag := make(chan bool)
	<-flag
}
