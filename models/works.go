package models

import (
	"time"

	"gorm.io/gorm"
)

type Works struct {
	CatId    int      `json:"cat_id"`
	Category Category `gorm:"foreignKey:cat_id"`

	WorksId          int    `gorm:"primary_key" json:"works_id"`
	UserId           int    `json:"user_id"`
	WorksName        string `json:"works_name"`
	WorksLink        string `json:"works_link"`
	WorksType        string `json:"works_type"`
	WorksDescription string `json:"works_description"`
	FavoriteNum      int    `json:"favorite_num"`
	ViewNum          int    `json:"view_num"`
	TopicNum         int    `json:"topic_num"`
	IsOpen           int    `json:"is_open"`
	State            int    `json:"state"`
	Remark           string `json:"remark"`
	CreateTime       string `json:"create_time"`
	UpdateTime       string `json:"update_time"`
	DeleteTimestamp  int    `json:"delete_timestamp"`
}

// 获取作品
func GetWorks(pageNum int, pageSize int, maps interface{}) (works []Works) {
	dbHandle.Preload("Category").Where(maps).Offset(pageNum).Limit(pageSize).Find(&works)
	return
}

func GetOneWorks(id int) (works Works) {
	dbHandle.Where("works_id = ?", id).First(&works)
	dbHandle.Model(&works).Association("Category").Find(&works.Category)

	return
}

// 获取总记录数
func GetWorksTotal(maps interface{}) (count int64) {
	dbHandle.Model(&Works{}).Where(maps).Count(&count)
	return
}

// 通过名称判断是否存在
func ExistWorksByName(name string) bool {
	var works Works
	dbHandle.Select("works_id").Where("works_name = ?", name).First(&works)

	return works.WorksId > 0
}

// 通过ID判断是否存在
func ExistWorksById(id int) bool {
	var works Works
	dbHandle.Select("works_id").Where("works_id = ?", id).First(&works)

	return works.WorksId > 0
}

// 新增数据
func AddWorks(data map[string]interface{}) bool {
	works := Works{
		CatId:            data["cat_id"].(int),
		WorksName:        data["works_name"].(string),
		UserId:           data["user_id"].(int),
		State:            data["state"].(int),
		WorksLink:        data["works_link"].(string),
		WorksType:        data["works_type"].(string),
		WorksDescription: data["works_description"].(string),
		Remark:           data["remark"].(string),
	}

	dbHandle.Create(&works)

	return true
}

func DeleteWorks(id int) bool {
	dbHandle.Where("works_id = ?", id).Delete(&Works{})

	return true
}

func EditWorks(id int, data interface{}) bool {
	dbHandle.Model(&Works{}).Where("works_id = ?", id).Updates(data)

	return true
}

// 创建回调函数
func (works *Works) BeforeCreate(scope *gorm.DB) error {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	works.CreateTime = nowTime
	works.UpdateTime = nowTime
	return nil
}

// 更新回调函数
func (works *Works) BeforeUpdate(scope *gorm.DB) error {
	works.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
