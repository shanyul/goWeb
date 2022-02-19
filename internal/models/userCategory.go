package models

import (
	"gorm.io/gorm"
	"time"
)

type UserCategoryModel struct{}

type UserCategory struct {
	UcatId          int    `gorm:"primary_key" column:"ucat_id" json:"ucatId"`
	UcatName        string `column:"ucat_name" json:"ucatName"`
	UserId          int    `column:"user_id" json:"userId"`
	Username        string `column:"user_name" json:"username"`
	CreateTime      string `column:"create_time" json:"createTime"`
	UpdateTime      string `column:"update_time" json:"updateTime"`
	DeleteTimestamp int    `column:"delete_timestamp" json:"deleteTimestamp"`
}

// 自定义表名
func (UserCategory) TableName() string {
	return "user_category"
}

// 获取分类
func (*UserCategoryModel) GetCategoryList(maps interface{}) ([]UserCategory, error) {
	var (
		category []UserCategory
		err      error
	)
	err = dbHandle.Where(maps).Find(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return category, nil
}

func (*UserCategoryModel) GetCategory(id int, userId int) (UserCategory, error) {
	var category UserCategory
	err := dbHandle.Where("ucat_id = ? AND user_id = ?", id, userId).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return category, err
	}

	return category, nil
}

// 通过名称判断是否存在
func (*UserCategoryModel) IsExist(name string, userId int) (bool, error) {
	var category UserCategory
	err := dbHandle.Where("ucat_name = ? AND user_id = ?", name, userId).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return category.UcatId > 0, nil
}

func (*UserCategoryModel) AddCategory(category *UserCategory) error {
	if err := dbHandle.Select(
		"UcatName",
		"UserId",
		"Username",
	).Create(&category).Error; err != nil {
		return err
	}

	return nil
}

func (*UserCategoryModel) EditCategory(id int, category *UserCategory) error {
	if err := dbHandle.Model(&UserCategory{}).Where("ucat_id = ?", id).Updates(category).Error; err != nil {
		return err
	}

	return nil
}

func (*UserCategoryModel) DeleteCategory(id int, userId int) error {
	maps := make(map[string]interface{})
	maps["delete_timestamp"] = time.Now().Unix()
	if err := dbHandle.Model(&UserCategory{}).Select("delete_timestamp").Where("ucat_id = ? AND user_id = ?", id, userId).Updates(maps).Error; err != nil {
		return err
	}

	return nil
}

// 获取总记录数
func (*UserCategoryModel) GetCategoryTotal(maps interface{}) (int64, error) {
	var count int64
	if err := dbHandle.Model(&UserCategory{}).Where(maps).Count(&count).Error; err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	return count, nil
}
