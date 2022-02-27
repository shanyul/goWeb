package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type WorksModel struct{}

type Works struct {
	BaseModel
	Tags []WorksTag `gorm:"foreignKey:works_id" json:"tags"`

	WorksId          int    `gorm:"primary_key" column:"works_id" json:"worksId"`
	UserId           int    `column:"user_id" json:"userId"`
	Username         string `column:"username" json:"username"`
	WorksName        string `column:"works_name" json:"worksName"`
	WorksLink        string `column:"works_link" json:"worksLink"`
	WorksType        int    `column:"works_type" json:"worksType"`
	WorksDescription string `column:"works_description" json:"worksDescription"`
	FavoriteNum      int    `column:"favorite_num" json:"favoriteNum"`
	ViewNum          int    `column:"view_num" json:"viewNum"`
	TopicNum         int    `column:"topic_num" json:"topicNum"`
	IsOpen           int    `column:"is_open" json:"isOpen"`
	State            int    `column:"state" json:"state"`
	Remark           string `column:"remark" json:"remark"`
	DeleteTimestamp  int    `column:"delete_timestamp" json:"deleteTimestamp"`
}

// 自定义表名
func (Works) TableName() string {
	return "works"
}

// GetWorks 获取作品
func (*WorksModel) GetWorks(pageNum int, pageSize int, maps map[string]interface{}, orderBy string) ([]Works, error) {
	var works []Works
	var err error
	query := dbHandle.Preload("Tags")
	if maps["delete_timestamp"] == 1 {
		delete(maps, "delete_timestamp")
		query = query.Where("delete_timestamp > ?", 100)
	}
	query = query.Where(maps)
	if orderBy != "" {
		query = query.Order(orderBy)
	}
	err = query.Limit(pageSize).Offset(pageNum).Find(&works).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return works, nil
}

func (*WorksModel) GetOneWorks(id int) (Works, error) {
	works := Works{}
	err := dbHandle.Preload("Tags").Where("works_id = ?", id).First(&works).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return works, err
	}

	return works, nil
}

func (*WorksModel) GetWorksTotal(maps map[string]interface{}) (int64, error) {
	var count int64
	var err error
	if maps["delete_timestamp"] == 1 {
		delete(maps, "delete_timestamp")
		err = dbHandle.Model(&Works{}).Where("delete_timestamp > ?", 100).Where(maps).Count(&count).Error
	} else {
		err = dbHandle.Model(&Works{}).Where(maps).Count(&count).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return count, nil
}

func (*WorksModel) Increment(id int, field string) error {
	opString := fmt.Sprintf("%s + ?", field)
	if err := dbHandle.Model(&Works{}).Where("works_id = ?", id).UpdateColumn(field, gorm.Expr(opString, 1)).Error; err != nil {
		return err
	}
	return nil
}

func (*WorksModel) Decrement(id int, field string) error {
	opString := fmt.Sprintf("%s - ?", field)
	if err := dbHandle.Model(&Works{}).Where(field+" > 0 AND works_id = ?", id).UpdateColumn(field, gorm.Expr(opString, 1)).Error; err != nil {
		return err
	}
	return nil
}

// ExistWorksByName 通过名称判断是否存在
func (*WorksModel) ExistWorksByName(name string) (bool, error) {
	var works Works
	err := dbHandle.Select("works_id").Where("works_name = ? AND deleted_timestamp = ? ", name, 0).First(&works).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return works.WorksId > 0, nil
}

func (*WorksModel) ExistWorksById(id int) (bool, error) {
	var works Works
	err := dbHandle.Select("works_id").Where("works_id = ?", id).First(&works).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return works.WorksId > 0, nil
}

// 新增数据
func (*WorksModel) AddWorks(works *Works) (int, error) {
	if err := dbHandle.Select(
		"cat_id", "works_name", "user_id", "is_open", "state", "works_link", "works_type", "works_description", "remark",
	).Create(&works).Error; err != nil {
		return 0, err
	}

	return works.WorksId, nil
}

func (*WorksModel) DeleteWorks(id int, userId int) error {
	maps := make(map[string]interface{})
	maps["delete_timestamp"] = time.Now().Unix()
	if err := dbHandle.Model(&Works{}).Select("delete_timestamp").Where("works_id = ? AND user_id = ?", id, userId).Updates(maps).Error; err != nil {
		return err
	}

	return nil
}

func (*WorksModel) Delete(id int, userId int) error {
	maps := make(map[string]interface{})
	maps["delete_timestamp"] = -1
	if err := dbHandle.Model(&Works{}).Select("delete_timestamp").Where("works_id = ? AND user_id = ?", id, userId).Updates(maps).Error; err != nil {
		return err
	}

	return nil
}

func (*WorksModel) EditWorks(id int, works Works) error {
	if err := dbHandle.Model(&Works{}).Where("works_id = ?", id).Updates(works).Error; err != nil {
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
