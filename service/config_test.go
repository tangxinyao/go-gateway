package service

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	result, err := readFile("../application.yml")
	if err != nil {
		t.Error()
	}
	fmt.Println(string(result))
}

func TestSetConfig(t *testing.T) {
	type Config struct {
		Mysql struct {
			Addr     string
			Username string
			Password string
		}
		Redis struct {
			Addr     string
			Password string
		}
	}
	var config Config
	SetConfig(&config, "./application.yml")
	d, e := json.Marshal(config)
	if e != nil {
		t.Fatal()
	}
	fmt.Println(string(d))
}

func TestXX(t *testing.T) {
	body := "{a:\"a\"}"
	var data map[string]string
	json.Unmarshal([]byte(body), data)
	d, _ := json.Marshal(data)
	fmt.Println(string(d))
}
