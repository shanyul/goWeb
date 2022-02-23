package models

import (
	"gorm.io/gorm"
)

type ViewerModel struct{}

type Viewer struct {
	BaseModel
	UserId  int `column:"user_id" json:"userId"`
	WorksId int `column:"works_id" json:"worksId"`
}

// 自定义表名
func (Viewer) TableName() string {
	return "viewer"
}

// 获取关注
func (*ViewerModel) GetView(maps interface{}) ([]Viewer, error) {
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
func (*ViewerModel) IsView(userId int, worksId int) bool {
	var view Viewer
	err := dbHandle.Where("user_id = ? AND works_id = ?", userId, worksId).First(&view).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}

	return view.UserId > 0
}

func (*ViewerModel) AddView(view *Viewer) error {
	if err := dbHandle.Select(
		"UserId",
		"WorksId",
	).Create(&view).Error; err != nil {
		return err
	}

	return nil
}

// 获取总记录数
func (*ViewerModel) GetViewTotal(maps interface{}) (int64, error) {
	var count int64
	if err := dbHandle.Model(&Viewer{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
