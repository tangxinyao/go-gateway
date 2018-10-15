package service

import (
	"regexp"
)

func FindNoRedirectPrefix(url string) (string, error) {
	regex, err := regexp.Compile(`(\/redirect)(.*)`)
	if err != nil {
		return "", err
	}
	data := regex.FindSubmatch([]byte(url))
	return string(data[2]), nil
}

func FindNoRobotPrefix(url string) (string, error) {
	regex, err := regexp.Compile(`(\/robot)(.*)`)
	if err != nil {
		return "", err
	}
	data := regex.FindSubmatch([]byte(url))
	return string(data[2]), nil
}
