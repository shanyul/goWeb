package models

import (
	"gorm.io/gorm"
	"time"
)

type TagsModel struct{}

type Tags struct {
	BaseModel
	Works []WorksTag `gorm:"ForeignKey:tag_id" json:"works"`

	TagId           int    `gorm:"primary_key" column:"tag_id" json:"tagId"`
	TagName         string `column:"tag_name" json:"tagName"`
	UserId          int    `column:"user_id" json:"userId"`
	Username        string `column:"username" json:"username"`
	Count           int    `column:"count" json:"count"`
	DeleteTimestamp int    `column:"delete_timestamp" json:"deleteTimestamp"`
}

// 自定义表名
func (Tags) TableName() string {
	return "tags"
}

// 获取分类
func (*TagsModel) GetTagList(orderBy string, maps interface{}) ([]Tags, error) {
	var (
		tags []Tags
		err  error
	)

	if orderBy != "" {
		err = dbHandle.Preload("Works").Where(maps).Order(orderBy).Find(&tags).Error
	} else {
		err = dbHandle.Preload("Works").Where(maps).Find(&tags).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

func (*TagsModel) GetTag(id int) (Tags, error) {
	var tag Tags
	err := dbHandle.Preload("Works").Where("tag_id = ?", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}

	return tag, nil
}

// 通过名称判断是否存在
func (*TagsModel) IsExist(name string, userId int) (bool, error) {
	var tag Tags
	err := dbHandle.Where("tag_name = ? AND user_id = ?", name, userId).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return tag.TagId > 0, nil
}

func (*TagsModel) AddTag(tag *Tags) error {
	if err := dbHandle.Select(
		"TagName",
		"UserId",
		"Username",
	).Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

func (*TagsModel) EditTag(id int, userId int, tag *Tags) error {
	if err := dbHandle.Model(&Tags{}).Where("tag_id = ? AND user_id = ?", id).Updates(tag).Error; err != nil {
		return err
	}

	return nil
}

func (*TagsModel) DeleteTag(id int, userId int) error {
	maps := make(map[string]interface{})
	maps["delete_timestamp"] = time.Now().Unix()
	if err := dbHandle.Model(&Tags{}).Select("delete_timestamp").Where("tag_id = ? AND user_id = ?", id, userId).Updates(maps).Error; err != nil {
		return err
	}

	return nil
}

// 获取总记录数
func (*TagsModel) GetTagTotal(maps interface{}) (int64, error) {
	var count int64
	if err := dbHandle.Model(&Tags{}).Where(maps).Count(&count).Error; err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	return count, nil
}