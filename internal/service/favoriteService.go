package service

import "designer-api/internal/models"

type FavoriteService struct {
	FavoriteModel models.FavoriteModel
	WorksService  WorksService
}

type Favorite struct {
	models.Favorite
}

func (service *FavoriteService) Add(favorite *Favorite) error {
	if isFavorite := service.FavoriteModel.IsFavorite(favorite.UserId, favorite.WorksId); !isFavorite {

		favoriteData := models.Favorite{}
		favoriteData.UserId = favorite.UserId
		favoriteData.WorksId = favorite.WorksId

		if err := service.FavoriteModel.AddFavorite(&favoriteData); err != nil {
			return err
		}

		field := "favorite_num"
		_ = service.WorksService.Increment(favorite.WorksId, field)
	}

	return nil
}

func (service *FavoriteService) GetAll(favorite *Favorite) ([]models.Favorite, error) {
	data, err := service.FavoriteModel.GetFavorite(service.getMaps(favorite))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *FavoriteService) IsFavorite(favorite *Favorite) bool {
	return service.FavoriteModel.IsFavorite(favorite.UserId, favorite.WorksId)
}

func (service *FavoriteService) Delete(favorite *Favorite) error {
	count, err := service.FavoriteModel.GetFavoriteTotal(service.getMaps(favorite))
	if err != nil {
		return err
	}

	if count > 0 {
		if err := service.FavoriteModel.DeleteFavorite(favorite.UserId, favorite.WorksId); err != nil {
			return err
		}
		// 扣减关注数
		field := "favorite_num"
		_ = service.WorksService.Decrement(favorite.WorksId, field)
	}

	return nil
}

func (service *FavoriteService) Count(favorite *Favorite) (int64, error) {
	return service.FavoriteModel.GetFavoriteTotal(service.getMaps(favorite))
}

func (service *FavoriteService) getMaps(favorite *Favorite) map[string]interface{} {
	maps := make(map[string]interface{})
	if favorite.WorksId > 0 {
		maps["works_id"] = favorite.WorksId
	}
	if favorite.UserId > 0 {
		maps["user_id"] = favorite.UserId
	}

	return maps
}
