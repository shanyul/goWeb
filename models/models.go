package models

import (
	"designer-api/pkg/setting"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbHandle *gorm.DB

// 连接数据库
func init() {
	var (
		err                          error
		dbName, user, password, host string
		connTimeout                  int
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	connTimeout = sec.Key("TIMEOUT").MustInt(3)

	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%ds",
		user, password, host, dbName, connTimeout)

	dbHandle, err = gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	db, err := dbHandle.DB()
	if err != nil {
		log.Println(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(50)
	db.SetConnMaxIdleTime(time.Duration(10) * time.Second)
}

// 关闭数据库
func CloseDB() {
	if db, err := dbHandle.DB(); err != nil {
		_ = db.Close()
	}
	dbHandle = nil
}
