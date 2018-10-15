package service

import (
	"fmt"
	"testing"
)

func TestFindNoRobotPrefix(t *testing.T) {
	result, err := FindNoRobotPrefix("/robot/hello")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestFindNoRedirectPrefix(t *testing.T) {
	result, err := FindNoRedirectPrefix("/redirect/hello")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
