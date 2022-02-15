package service

import "designer-api/internal/models"

type Favorite struct {
	UserId     int
	WorksId    int
	CreateTime string
}

func (favorite *Favorite) Add() error {
	if isFavorite := models.IsFavorite(favorite.UserId, favorite.WorksId); !isFavorite {
		favoriteData := map[string]interface{}{
			"userId":  favorite.UserId,
			"worksId": favorite.WorksId,
		}

		if err := models.AddFavorite(favoriteData); err != nil {
			return err
		}

		worksService := Works{
			WorksId: favorite.WorksId,
		}

		field := "favorite_num"
		_ = worksService.Increment(field)
	}

	return nil
}

func (favorite *Favorite) GetAll() ([]models.Favorite, error) {
	data, err := models.GetFavorite(favorite.getMaps())
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (favorite *Favorite) isFavorite() bool {
	return models.IsFavorite(favorite.UserId, favorite.WorksId)
}

func (favorite *Favorite) Delete() error {
	count, err := models.GetFavoriteTotal(favorite.getMaps())
	if err != nil {
		return err
	}

	if count > 0 {
		if err := models.DeleteFavorite(favorite.UserId, favorite.WorksId); err != nil {
			return err
		}
		// 扣减关注数
		worksService := Works{
			WorksId: favorite.WorksId,
		}
		field := "favorite_num"
		_ = worksService.Decrement(field)
	}

	return nil
}

func (favorite *Favorite) Count() (int64, error) {
	return models.GetFavoriteTotal(favorite.getMaps())
}

func (favorite *Favorite) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if favorite.WorksId > 0 {
		maps["works_id"] = favorite.WorksId
	}
	if favorite.UserId > 0 {
		maps["user_id"] = favorite.UserId
	}

	return maps
}
