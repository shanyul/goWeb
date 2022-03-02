package models

import (
	"gorm.io/gorm"
)

type WorksTagModel struct{}

type WorksTag struct {
	TagId     int    `gorm:"primaryKey" column:"tag_id" json:"tagId"`
	WorksId   int    `gorm:"primaryKey" column:"works_id" json:"worksId"`
	TagName   string `column:"tag_name" json:"tagName"`
	WorksName string `column:"works_name" json:"worksName"`
	IsDelete  int    `column:"is_delete" json:"isDelete"`

	//Tags  Tags  `gorm:"foreignKey:tag_id" json:"tags"`
}

// 自定义表名
func (WorksTag) TableName() string {
	return "works_tag"
}

// 获取分类
func (*WorksTagModel) GetList(pageNum int, pageSize int, maps interface{}) ([]WorksTag, error) {
	var (
		worksTag []WorksTag
		err      error
	)
	err = dbHandle.Where(maps).Limit(pageSize).Offset(pageNum).Find(&worksTag).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return worksTag, nil
}

func (*WorksTagModel) GetWorkIdByTagId(pageNum int, pageSize int, id []int) (worksTag []WorksTag, err error) {
	err = dbHandle.Select("works_id").Where("tag_id IN ?", id).Limit(pageSize).Offset(pageNum).Find(&worksTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	return
}

// 通过名称判断是否存在
func (*WorksTagModel) IsExist(tagId int, worksId int) (bool, error) {
	var worksTag WorksTag
	err := dbHandle.Where("tag_id = ? AND works_id = ?", tagId, worksId).First(&worksTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return worksTag.TagId > 0, nil
}

func (*WorksTagModel) BatchAdd(worksTag []WorksTag) error {
	if err := dbHandle.Create(&worksTag).Error; err != nil {
		return err
	}

	return nil
}

func (*WorksTagModel) Delete(worksId int) error {
	if err := dbHandle.Where("works_id = ?", worksId).Delete(&WorksTag{}).Error; err != nil {
		return err
	}

	return nil
}

func (*WorksTagModel) SoftDelete(worksId int, deleteType int) error {
	if err := dbHandle.Model(&WorksTag{}).Where("works_id = ?", worksId).Update("is_delete", deleteType).Error; err != nil {
		return err
	}

	return nil
}

func (*WorksTagModel) DeleteByTag(tagId int) error {
	if err := dbHandle.Where("tag_id = ?", tagId).Delete(&WorksTag{}).Error; err != nil {
		return err
	}

	return nil
}

func (*WorksTagModel) UpdateTag(tagId int, tagName string) error {
	if err := dbHandle.Model(&WorksTag{}).Where("tag_id = ?", tagId).Update("tag_name", tagName).Error; err != nil {
		return err
	}

	return nil
}

// 获取总记录数
func (*WorksTagModel) GetTotal(maps interface{}) (int64, error) {
	var count int64
	if err := dbHandle.Model(&WorksTag{}).Where(maps).Count(&count).Error; err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	return count, nil
}
