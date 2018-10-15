package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"log"
)

var MyLogger *log.Logger
var RedisClient *redis.Client
var MySQLClient *gorm.DB
var RedirectUrl string
