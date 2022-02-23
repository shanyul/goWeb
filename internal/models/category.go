package models

import (
	"gorm.io/gorm"
)

type CategoryModel struct{}

type Category struct {
	BaseModel
	CatId           int    `gorm:"primary_key" column:"cat_id" json:"catId"`
	CatName         string `column:"cat_name" json:"catName"`
	ParentId        int    `column:"parent_id" json:"parentId"`
	IsDirectory     int    `column:"is_directory" json:"isDirectory"`
	Level           int    `column:"level" json:"level"`
	Path            string `column:"path" json:"path"`
	DeleteTimestamp int    `column:"delete_timestamp" json:"deleteTimestamp"`
}

// 自定义表名
func (Category) TableName() string {
	return "category"
}

// 获取作品
func (*CategoryModel) GetCategory(pageNum int, pageSize int, maps interface{}) ([]Category, error) {
	var (
		category []Category
		err      error
	)

	if pageSize > 0 && pageNum > 0 {
		err = dbHandle.Where(maps).Offset(pageNum).Limit(pageSize).Find(&category).Error
	} else {
		err = dbHandle.Where(maps).Find(&category).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return category, nil
}

// 获取总记录数
func (*CategoryModel) GetCategoryTotal(maps interface{}) (int64, error) {
	var count int64
	if err := dbHandle.Model(&Category{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// 通过名称判断是否存在
func (*CategoryModel) ExistCategoryByName(name string) (bool, error) {
	var category Category
	err := dbHandle.Select("cat_id").Where("cat_name = ? AND delete_timestamp = ?", name, 0).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return category.CatId > 0, nil
}

// 通过ID判断是否存在
func (*CategoryModel) ExistCategoryById(id int) (bool, error) {
	var category Category
	err := dbHandle.Select("cat_id").Where("cat_id = ? AND delete_timestamp = ?", id, 0).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return category.CatId > 0, nil
}

// 新增数据
func (*CategoryModel) AddCategory(data *Category) error {

	if err := dbHandle.Select("cat_name", "parent_id").Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (*CategoryModel) DeleteCategory(id int) error {
	if err := dbHandle.Where("cat_id = ?", id).Delete(Category{}).Error; err != nil {
		return err
	}

	return nil
}

func (*CategoryModel) EditCategory(id int, category Category) error {
	if err := dbHandle.Model(&Category{}).Where("cat_id = ?", id).Updates(category).Error; err != nil {
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
