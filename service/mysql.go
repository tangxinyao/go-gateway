package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func GetMySQL(config MysqlConfig) (*gorm.DB, error) {
	mysqlUrl := fmt.Sprintf("%s:%s@%s", config.Username, config.Password, config.Addr)
	var err error
	dbSQL, err := sql.Open("mysql", mysqlUrl)
	if err != nil {
		return nil, err
	}
	dbSQL.SetConnMaxLifetime(time.Second * 10)

	db, err := gorm.Open("mysql", dbSQL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
