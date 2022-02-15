package service

import "designer-api/internal/models"

type Viewer struct {
	UserId     int
	WorksId    int
	CreateTime string
}

func (view *Viewer) Add() error {
	if isView := models.IsFavorite(view.UserId, view.WorksId); !isView {
		viewData := map[string]interface{}{
			"userId":  view.UserId,
			"worksId": view.WorksId,
		}

		if err := models.AddView(viewData); err != nil {
			return err
		}
	}

	worksService := Works{
		WorksId: view.WorksId,
	}
	field := "view_num"
	_ = worksService.Increment(field)

	return nil
}

func (view *Viewer) GetAll() ([]models.Viewer, error) {
	data, err := models.GetView(view.getMaps())
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (view *Viewer) isView() bool {
	return models.IsFavorite(view.UserId, view.WorksId)
}

func (view *Viewer) Count() (int64, error) {
	return models.GetFavoriteTotal(view.getMaps())
}

func (view *Viewer) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if view.WorksId > 0 {
		maps["works_id"] = view.WorksId
	}
	if view.UserId > 0 {
		maps["user_id"] = view.UserId
	}

	return maps
}
