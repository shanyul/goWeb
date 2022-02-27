package service

import "designer-api/internal/models"

type ViewService struct {
	ViewerModel models.ViewerModel
	WorksModel  models.WorksModel
}

type Viewer struct {
	models.Viewer
}

func (service *ViewService) Add(view *Viewer) error {
	/*if isView := service.ViewerModel.IsView(view.UserId, view.WorksId); !isView {

		viewData := models.Viewer{}
		viewData.UserId = view.UserId
		viewData.WorksId = view.WorksId

		if err := service.ViewerModel.AddView(&viewData); err != nil {
			return err
		}
	}*/

	field := "view_num"
	_ = service.WorksModel.Increment(view.WorksId, field)

	return nil
}

func (service *ViewService) GetAll(view *Viewer) ([]models.Viewer, error) {
	data, err := service.ViewerModel.GetView(service.getMaps(view))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *ViewService) isView(userId int, worksId int) bool {
	return service.ViewerModel.IsView(userId, worksId)
}

func (service *ViewService) Count(view *Viewer) (int64, error) {
	return service.ViewerModel.GetViewTotal(service.getMaps(view))
}

func (service *ViewService) getMaps(view *Viewer) map[string]interface{} {
	maps := make(map[string]interface{})
	if view.WorksId > 0 {
		maps["works_id"] = view.WorksId
	}
	if view.UserId > 0 {
		maps["user_id"] = view.UserId
	}

	return maps
}
