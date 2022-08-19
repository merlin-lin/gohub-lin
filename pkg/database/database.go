package database

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQLDB *sql.DB

func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	var err error
	// 使用 gorm.Open 连接数据库
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层的 sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())

	}
}
