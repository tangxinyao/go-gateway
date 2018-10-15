package service

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"
)

const saltSize = 16

func GenerateSalt(secret []byte) (string, error) {
	buf := make([]byte, saltSize, saltSize+sha1.Size)
	_, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		return "", err
	}

	hash := sha1.New()
	hash.Write(buf)
	hash.Write(secret)
	return fmt.Sprintf("%x", hash.Sum(buf)), nil
}

func HashPassword(password string, salt string) string {
	// generate password + salt hash to store into database
	combination := string(salt) + string(password)
	passwordHash := sha1.New()
	io.WriteString(passwordHash, combination)
	return fmt.Sprintf("%x", passwordHash.Sum(nil))
}
