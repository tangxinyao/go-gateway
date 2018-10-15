package service

import (
	"gopkg.in/yaml.v1"
	"io/ioutil"
)

type MysqlConfig struct {
	Addr     string
	Username string
	Password string
}

type RedisConfig struct {
	Addr     string
	Password string
}

type ServerConfig struct {
	Addr string
	Port string
}

type RedirectConfig struct {
	Addr      string
}

type Config struct {
	Mysql  MysqlConfig  `yaml:"mysql"`
	Redis  RedisConfig  `yaml:"redis"`
	Server ServerConfig `yaml:"server"`
	Redirect  RedirectConfig  `yaml:"wxbot"`
}

func SetConfig(category interface{}, filename string) (interface{}, error) {
	in, err := readFile(filename)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(in, category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func readFile(filename string) ([]byte, error) {
	configuration, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return configuration, nil
}
