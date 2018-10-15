package service

import (
	"fmt"
	"testing"
)

func TestGenerateSalt(t *testing.T) {
	salt, err := GenerateSalt([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Salt : %x \n", salt)
}

func TestHashPassword(t *testing.T) {
	str := "hello"
	salt, err := GenerateSalt([]byte(str))
	if err != nil {
		t.Fatal(err)
	}
	result := HashPassword(str, salt)
	fmt.Println(salt)
	fmt.Println(result)
}
