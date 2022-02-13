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

// SetUp 连接数据库
func SetUp() {
	var err error

	dbHandle, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%ds",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Timeout,
	)), &gorm.Config{})

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

// CloseDB 关闭数据库
func CloseDB() {
	if db, err := dbHandle.DB(); err != nil {
		_ = db.Close()
	}
	dbHandle = nil
}
