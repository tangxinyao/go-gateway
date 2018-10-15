package domain

import (
	"encoding/json"
	"flag"
	"fmt"
	"go-gateway/global"
	"go-gateway/service"
	"os"
	"testing"
)

func TestCreateUser(t *testing.T) {
	SetLogger()
	config := GetConfig()
	SetDB(config)

	var roles []string
	user, err := CreateUser("root", "test123456", roles)
	result, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(result))
}

func SetDB(config service.Config) {
	var err error
	global.RedisClient, err = service.GetRedis(config.Redis)
	if err != nil {
		fmt.Println("Something unexcepted happened when connecting redis")
		os.Exit(0)
	}
	global.MySQLClient, err = service.GetMySQL(config.Mysql)

	global.MySQLClient.Model(User{}).Related(Role{}, "roles")
	global.MySQLClient.Model(Role{}).Related(Permission{}, "permissions")
	global.MySQLClient.AutoMigrate(User{}, Role{}, Permission{})

	if err != nil {
		fmt.Println("Something unexcepted happened when connecting mysql")
		os.Exit(0)
	}
}

func GetConfig() service.Config {
	// Parse flag
	filename := flag.String("config.location", "../application.yml", "Location of configuration")
	flag.Parse()
	var config service.Config
	service.SetConfig(&config, *filename)
	global.MyLogger.Println(config)
	return config
}

func SetLogger() {
	var err error
	global.MyLogger, err = service.GetLogger("0 0 0 * * ?")
	if err != nil {
		fmt.Println("Something unexcepted happened when creating logger")
		os.Exit(0)
	}
}
