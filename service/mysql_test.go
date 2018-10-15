package service

import (
	"testing"
)

func TestGetMySQL(t *testing.T) {
	config := MysqlConfig{
		Addr:     "tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local",
		Username: "root",
		Password: "",
	}
	db, err := GetMySQL(config)
	if err != nil {
		t.Error(err)
	}
	db.Close()
}
