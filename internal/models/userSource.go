package models

import (
	"gorm.io/gorm"
	"time"
)

type UserSourceModel struct{}

type Source struct {
	UcatId   int          `column:"ucat_id" json:"ucatId"`
	Category UserCategory `gorm:"foreignKey:ucat_id" json:"category"`

	SourceId        int    `gorm:"primary_key" column:"source_id" json:"sourceId"`
	UserId          int    `column:"user_id" json:"userId"`
	Username        string `column:"user_name" json:"username"`
	UcatName        string `column:"ucat_name" json:"ucatName"`
	Description     string `column:"description" json:"description"`
	Link            string `column:"link" json:"link"`
	Title           string `column:"title" json:"title"`
	CreateTime      string `column:"create_time" json:"createTime"`
	UpdateTime      string `column:"update_time" json:"updateTime"`
	DeleteTimestamp int    `column:"delete_timestamp" json:"deleteTimestamp"`
}

// 自定义表名
func (Source) TableName() string {
	return "user_source"
}

// 获取素材
func (*UserSourceModel) GetSourceList(pageNum int, pageSize int, maps interface{}) ([]Source, error) {
	var (
		source []Source
		err    error
	)
	err = dbHandle.Where(maps).Find(&source).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return source, err
	}

	err = dbHandle.Preload("Category").Where(maps).Limit(pageSize).Offset(pageNum).Find(&source).Error
	if err != nil {
		return source, err
	}

	return source, nil
}

// 通过ID获取素材
func (*UserSourceModel) GetSource(id int) (Source, error) {
	var source Source
	err := dbHandle.First(&source, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return source, err
	}

	return source, nil
}

func (*UserSourceModel) AddSource(source *Source) error {
	if err := dbHandle.Select(
		"UserId",
		"Username",
		"UcatId",
		"UcatName",
		"Description",
		"Title",
		"Link",
	).Create(&source).Error; err != nil {
		return err
	}

	return nil
}

func (*UserSourceModel) EditSource(id int, source *Source) error {
	if err := dbHandle.Model(&Source{}).Where("source_id = ?", id).Updates(source).Error; err != nil {
		return err
	}

	return nil
}

func (*UserSourceModel) DeleteSource(id int, userId int) error {
	maps := make(map[string]interface{})
	maps["delete_timestamp"] = time.Now().Unix()
	if err := dbHandle.Model(&Source{}).Select("delete_timestamp").Where("source_id = ? AND user_id = ?", id, userId).Updates(maps).Error; err != nil {
		return err
	}

	return nil
}

// 获取总记录数
func (*UserSourceModel) GetSourceTotal(maps interface{}) (int64, error) {
	var count int64
	if err := dbHandle.Model(&Source{}).Where(maps).Count(&count).Error; err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	return count, nil
}
