package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"go-gateway/domain"
	"go-gateway/global"
	"go-gateway/middleware"
	"go-gateway/router"
	"go-gateway/service"
	"os"
)

func main() {
	// Get the configuration file, initialize the logger, redis client, and mysql client, and start the web server
	SetLogger()
	config := GetConfig()
	SetProxy(config)
	SetDB(config)
	StartServer(config)
}

// router table
func StartServer(config service.Config) {
	e := echo.New()
	e.Use(middleware.Logger)
	e.Use(middleware.CORS)
	router.SetRouter(e)
	serverAddress := fmt.Sprintf("%s:%s", config.Server.Addr, config.Server.Port)
	e.Logger.Fatal(e.Start(serverAddress))
}

func SetProxy(config service.Config) {
	global.RedirectUrl = config.Redirect.Addr
}

func SetDB(config service.Config) {
	var err error
	global.RedisClient, err = service.GetRedis(config.Redis)
	if err != nil {
		fmt.Println("Something unexcepted happened when connecting redis")
		os.Exit(0)
	}
	global.MySQLClient, err = service.GetMySQL(config.Mysql)

	global.MySQLClient.Model(&domain.User{}).Related(&domain.Role{}, "roles")
	global.MySQLClient.Model(&domain.Role{}).Related(&domain.Permission{}, "permissions")
	global.MySQLClient.AutoMigrate(&domain.User{}, &domain.Role{}, &domain.Permission{})

	if err != nil {
		fmt.Println("Something unexcepted happened when connecting mysql")
		os.Exit(0)
	}
}

func GetConfig() service.Config {
	// Parse flag
	filename := flag.String("config.location", "./application.yml", "Location of configuration")
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
