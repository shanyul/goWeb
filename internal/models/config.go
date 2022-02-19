package models

import (
	"gorm.io/gorm"
)

type ConfigModel struct{}

type Config struct {
	ConfigId   int    `gorm:"primary_key" column:"config_id" json:"configId"`
	Key        string `column:"key" json:"key"`
	Value      string `column:"value" json:"value"`
	CreateTime string `column:"create_time" json:"createTime"`
	UpdateTime string `column:"update_time" json:"updateTime"`
}

// 自定义表名
func (Config) TableName() string {
	return "config"
}

// 获取配置
func (*ConfigModel) GetConfigList(maps interface{}) ([]Config, error) {
	var (
		config []Config
		err    error
	)
	err = dbHandle.Where(maps).Find(&config).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return config, nil
}

func (*ConfigModel) GetConfig(key string) (Config, error) {
	var config Config
	err := dbHandle.Where("key = ?", key).First(&config).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return config, err
	}

	return config, nil
}

// 通过名称判断是否存在
func (*ConfigModel) IsExist(key string) (bool, error) {
	var config Config
	err := dbHandle.Where("key = ?", key).First(&config).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return config.ConfigId > 0, nil
}

func (*ConfigModel) AddConfig(config *Config) error {
	if err := dbHandle.Select(
		"Key",
		"Value",
	).Create(&config).Error; err != nil {
		return err
	}

	return nil
}

func (*ConfigModel) EditConfig(key string, config *Config) error {
	if err := dbHandle.Model(&Config{}).Where("key = ?", key).Updates(config).Error; err != nil {
		return err
	}

	return nil
}

func (*ConfigModel) DeleteConfig(id int) error {
	if err := dbHandle.Where("config_id = ?", id).Delete(&Config{}).Error; err != nil {
		return err
	}

	return nil
}

// 获取总记录数
func (*ConfigModel) GetConfigTotal(maps interface{}) (int64, error) {
	var count int64
	if err := dbHandle.Model(&Config{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
