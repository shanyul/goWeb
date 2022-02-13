package models

import (
	"gorm.io/gorm"
)

type Category struct {
	CatId           int    `gorm:"primary_key" column:"cat_id" json:"catId"`
	CatName         string `column:"cat_name" json:"catName"`
	ParentId        int    `column:"parent_id" json:"parentId"`
	IsDirectory     int    `column:"is_directory" json:"isDirectory"`
	Level           int    `column:"level" json:"level"`
	Path            string `column:"path" json:"path"`
	CreateTime      string `column:"create_time" json:"createTime"`
	UpdateTime      string `column:"update_time" json:"updateTime"`
	DeleteTimestamp int    `column:"delete_timestamp" json:"deleteTimestamp"`
}

// 自定义表名
func (Category) TableName() string {
	return "category"
}

// 获取作品
func GetCategory(pageNum int, pageSize int, maps interface{}) ([]Category, error) {
	var (
		category []Category
		err      error
	)

	if pageSize > 0 && pageNum > 0 {
		err = dbHandle.Where(maps).Find(&category).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = dbHandle.Where(maps).Find(&category).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return category, nil
}

// 获取总记录数
func GetCategoryTotal(maps interface{}) (int64, error) {
	var count int64
	if err := dbHandle.Model(&Category{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// 通过名称判断是否存在
func ExistCategoryByName(name string) (bool, error) {
	var category Category
	err := dbHandle.Select("cat_id").Where("cat_name = ? AND delete_timestamp = ?", name, 0).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return category.CatId > 0, nil
}

// 通过ID判断是否存在
func ExistCategoryById(id int) (bool, error) {
	var category Category
	err := dbHandle.Select("cat_id").Where("cat_id = ? AND delete_timestamp = ?", id, 0).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return category.CatId > 0, nil
}

// 新增数据
func AddCategory(data map[string]interface{}) error {
	category := Category{
		CatName:  data["catName"].(string),
		ParentId: data["parentId"].(int),
	}

	if err := dbHandle.Select("cat_name", "parent_id").Create(&category).Error; err != nil {
		return err
	}

	return nil
}

func DeleteCategory(id int) error {
	if err := dbHandle.Where("cat_id = ?", id).Delete(Category{}).Error; err != nil {
		return err
	}

	return nil
}

func EditCategory(id int, data interface{}) error {
	if err := dbHandle.Model(&Category{}).Where("cat_id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

/*// 创建回调函数
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
}*/
