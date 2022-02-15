package models

import (
	"gorm.io/gorm"
)

type Viewer struct {
	UserId     int    `column:"user_id" json:"userId"`
	WorksId    int    `column:"works_id" json:"worksId"`
	CreateTime string `column:"create_time" json:"createTime"`
}

// 自定义表名
func (Viewer) TableName() string {
	return "viewer"
}

// 获取关注
func GetView(maps interface{}) ([]Viewer, error) {
	var (
		view []Viewer
		err  error
	)
	err = dbHandle.Where(maps).Find(&view).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return view, nil
}

// 是否查看
func IsView(userId int, worksId int) bool {
	var view Viewer
	err := dbHandle.Where("user_id = ? AND works_id = ?", userId, worksId).First(&view).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}

	return view.UserId > 0
}

func AddView(data map[string]interface{}) error {
	view := Viewer{
		UserId:  data["userId"].(int),
		WorksId: data["worksId"].(int),
	}

	if err := dbHandle.Select(
		"UserId",
		"WorksId",
	).Create(&view).Error; err != nil {
		return err
	}

	return nil
}

// 获取总记录数
func GetViewTotal(maps interface{}) (int64, error) {
	var count int64
	if err := dbHandle.Model(&Viewer{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
