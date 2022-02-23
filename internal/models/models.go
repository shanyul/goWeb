package models

import (
	"database/sql/driver"
	"designer-api/pkg/setting"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbHandle *gorm.DB

type BaseModel struct {
	CreateTime XTime `column:"create_time" json:"createTime"`
	UpdateTime XTime `column:"update_time" json:"updateTime"`
}

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

type XTime struct {
	time.Time
}

// 2. 为 Xtime 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t XTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

// 3. 为 Xtime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 4. 为 Xtime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
