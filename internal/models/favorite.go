package models

import (
	"gorm.io/gorm"
)

type FavoriteModel struct{}

type Favorite struct {
	BaseModel
	UserId  int `column:"user_id" json:"userId"`
	WorksId int `column:"works_id" json:"worksId"`
}

// 自定义表名
func (Favorite) TableName() string {
	return "favorite"
}

// 获取关注
func (*FavoriteModel) GetFavorite(maps interface{}) ([]Favorite, error) {
	var (
		favorite []Favorite
		err      error
	)
	err = dbHandle.Where(maps).Find(&favorite).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return favorite, nil
}

// 是否关注
func (*FavoriteModel) IsFavorite(userId int, worksId int) bool {
	var favorite Favorite
	err := dbHandle.Where("user_id = ? AND works_id = ?", userId, worksId).First(&favorite).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}

	return favorite.UserId > 0
}

func (*FavoriteModel) AddFavorite(favorite *Favorite) error {
	if err := dbHandle.Select(
		"UserId",
		"WorksId",
	).Create(&favorite).Error; err != nil {
		return err
	}

	return nil
}

// 获取总记录数
func (*FavoriteModel) GetFavoriteTotal(maps interface{}) (int64, error) {
	var count int64
	if err := dbHandle.Model(&Favorite{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (*FavoriteModel) DeleteFavorite(userId int, worksId int) error {
	if err := dbHandle.Where("user_id = ? AND works_id = ?", userId, worksId).Delete(Favorite{}).Error; err != nil {
		return err
	}

	return nil
}
