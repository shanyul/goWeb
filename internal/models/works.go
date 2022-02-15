package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Works struct {
	CatId    int      `json:"cat_id"`
	Category Category `gorm:"foreignKey:cat_id" json:"category"`

	WorksId          int    `gorm:"primary_key" column:"works_id" json:"worksId"`
	UserId           int    `column:"user_id" json:"userId"`
	Username         string `column:"username" json:"username"`
	WorksName        string `column:"works_name" json:"worksName"`
	WorksLink        string `column:"works_link" json:"worksLink"`
	WorksType        string `column:"works_type" json:"worksType"`
	WorksDescription string `column:"works_description" json:"worksDescription"`
	FavoriteNum      int    `column:"favorite_num" json:"favoriteNum"`
	ViewNum          int    `column:"view_num" json:"viewNum"`
	TopicNum         int    `column:"topic_num" json:"topicNum"`
	IsOpen           int    `column:"is_open" json:"isOpen"`
	State            int    `column:"state" json:"state"`
	Remark           string `column:"remark" json:"remark"`
	CreateTime       string `column:"create_time" json:"createTime"`
	UpdateTime       string `column:"update_time" json:"updateTime"`
	DeleteTimestamp  int    `column:"delete_timestamp" json:"deleteTimestamp"`
}

// GetWorks 获取作品
func GetWorks(pageNum int, pageSize int, maps interface{}, orderBy string) ([]Works, error) {
	var works []Works
	var err error
	if orderBy != "" {
		err = dbHandle.Preload("Category").Where(maps).Order(orderBy).Limit(pageSize).Offset(pageNum).Find(&works).Error
	} else {
		err = dbHandle.Preload("Category").Where(maps).Limit(pageSize).Offset(pageNum).Find(&works).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return works, nil
}

func GetOneWorks(id int) (*Works, error) {
	var works Works
	err := dbHandle.Where("works_id = ?", id).First(&works).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	err = dbHandle.Model(&works).Association("Category").Find(&works.Category)
	if err != nil {
		return nil, err
	}

	return &works, nil
}

func GetWorksTotal(maps interface{}) (int64, error) {
	var count int64
	if err := dbHandle.Model(&Works{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func Increment(id int, field string) error {
	opString := fmt.Sprintf("%s + ?", field)
	if err := dbHandle.Model(&Works{}).Where("works_id = ?", id).UpdateColumn(field, gorm.Expr(opString, 1)).Error; err != nil {
		return err
	}
	return nil
}

func Decrement(id int, field string) error {
	opString := fmt.Sprintf("%s - ?", field)
	if err := dbHandle.Model(&Works{}).Where(field+" > 0 AND works_id = ?", id).UpdateColumn(field, gorm.Expr(opString, 1)).Error; err != nil {
		return err
	}
	return nil
}

// ExistWorksByName 通过名称判断是否存在
func ExistWorksByName(name string) (bool, error) {
	var works Works
	err := dbHandle.Select("works_id").Where("works_name = ? AND deleted_timestamp = ? ", name, 0).First(&works).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return works.WorksId > 0, nil
}

func ExistWorksById(id int) (bool, error) {
	var works Works
	err := dbHandle.Select("works_id").Where("works_id = ?", id).First(&works).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return works.WorksId > 0, nil
}

// 新增数据
func AddWorks(data map[string]interface{}) error {
	works := Works{
		CatId:            data["catId"].(int),
		WorksName:        data["worksName"].(string),
		UserId:           data["userId"].(int),
		Username:         data["username"].(string),
		State:            data["state"].(int),
		WorksLink:        data["worksLink"].(string),
		WorksType:        data["worksType"].(string),
		WorksDescription: data["worksDescription"].(string),
		Remark:           data["remark"].(string),
	}

	if err := dbHandle.Select(
		"cat_id", "works_name", "user_id", "state", "works_link", "works_type", "works_description", "remark",
	).Create(&works).Error; err != nil {
		return err
	}

	return nil
}

func DeleteWorks(id int, userId int) error {
	maps := make(map[string]interface{})
	maps["delete_timestamp"] = time.Now().Unix()
	if err := dbHandle.Model(&Works{}).Select("delete_timestamp").Where("works_id = ? AND user_id = ?", id, userId).Updates(maps).Error; err != nil {
		return err
	}

	return nil
}

func EditWorks(id int, data interface{}) error {
	if err := dbHandle.Model(&Works{}).Where("works_id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

/*// 创建回调函数
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
}*/
