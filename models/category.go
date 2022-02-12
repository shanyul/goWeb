package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	CatId           int    `gorm:"primary_key" json:"cat_id"`
	CatName         string `json:"cat_name"`
	ParentId        int    `json:"parent_id"`
	IsDirectory     int    `json:"is_directory"`
	Level           int    `json:"level"`
	Path            string `json:"path"`
	CreateTime      string `json:"create_time"`
	UpdateTime      string `json:"update_time"`
	DeleteTimestamp int    `json:"delete_timestamp"`
}

// 自定义表名
func (Category) TableName() string {
	return "category"
}

// 获取作品
func GetCategory(pageNum int, pageSize int, maps interface{}) (category []Category) {
	dbHandle.Where(maps).Offset(pageNum).Limit(pageSize).Find(&category)
	return
}

// 获取总记录数
func GetCategoryTotal(maps interface{}) (count int64) {
	dbHandle.Model(&Category{}).Where(maps).Count(&count)
	return
}

// 通过名称判断是否存在
func ExistCategoryByName(name string) bool {
	var category Category
	dbHandle.Select("cat_id").Where("cat_name = ?", name).First(&category)

	return category.CatId > 0
}

// 通过ID判断是否存在
func ExistCategoryById(id int) bool {
	var category Category
	dbHandle.Select("cat_id").Where("cat_id = ?", id).First(&category)

	return category.CatId > 0
}

// 新增数据
func AddCategory(name string, parentId int) bool {
	dbHandle.Create(&Category{
		CatName:  name,
		ParentId: parentId,
	})

	return true
}

func DeleteCategory(id int) bool {
	dbHandle.Where("cat_id = ?", id).Delete(&Category{})

	return true
}

func EditCategory(id int, data interface{}) bool {
	dbHandle.Model(&Category{}).Where("cat_id = ?", id).Updates(data)

	return true
}

// 创建回调函数
func (category *Category) BeforeCreate(scope *gorm.DB) error {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	category.CreateTime = nowTime
	category.UpdateTime = nowTime
	return nil
}

// 更新回调函数
func (category *Category) BeforeUpdate(scope *gorm.DB) error {
	category.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
