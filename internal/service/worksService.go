package service

import (
	"designer-api/internal/models"
)

type WorksService struct {
	WorksModel models.WorksModel
}

type Works struct {
	models.Works

	OrderBy  string
	PageNum  int
	PageSize int
}

func (service *WorksService) ExistByName(worksName string) (bool, error) {
	return service.WorksModel.ExistWorksByName(worksName)
}

func (service *WorksService) Add(w *Works) error {
	worksData := models.Works{}
	worksData.WorksName = w.WorksName
	worksData.UserId = w.UserId
	worksData.Username = w.Username
	worksData.State = w.State
	worksData.CatId = w.CatId
	worksData.WorksLink = w.WorksLink
	worksData.WorksType = w.WorksType
	worksData.WorksDescription = w.WorksDescription
	worksData.Remark = w.Remark

	if err := service.WorksModel.AddWorks(&worksData); err != nil {
		return err
	}

	return nil
}

func (service *WorksService) Edit(w *Works) error {

	worksData := models.Works{}
	worksData.WorksName = w.WorksName
	worksData.State = w.State
	worksData.CatId = w.CatId
	worksData.WorksLink = w.WorksLink
	worksData.WorksType = w.WorksType
	worksData.WorksDescription = w.WorksDescription
	worksData.Remark = w.Remark

	if err := service.WorksModel.EditWorks(w.WorksId, worksData); err != nil {
		return err
	}

	return nil
}

func (service *WorksService) GetAll(w *Works) ([]models.Works, error) {
	data, err := service.WorksModel.GetWorks(w.PageNum, w.PageSize, service.getMaps(w), w.OrderBy)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *WorksService) Get(worksId int) (*models.Works, error) {
	works, err := service.WorksModel.GetOneWorks(worksId)
	if err != nil {
		return nil, err
	}

	return works, nil
}

func (service *WorksService) Delete(worksId int, userId int) error {
	return service.WorksModel.DeleteWorks(worksId, userId)
}

func (service *WorksService) ExistByID(worksId int) (bool, error) {
	return service.WorksModel.ExistWorksById(worksId)
}

func (service *WorksService) Count(works *Works) (int64, error) {
	return service.WorksModel.GetWorksTotal(service.getMaps(works))
}

func (service *WorksService) Increment(worksId int, field string) error {
	return service.WorksModel.Increment(worksId, field)
}

func (service *WorksService) Decrement(worksId int, field string) error {
	return service.WorksModel.Decrement(worksId, field)
}

func (service *WorksService) getMaps(w *Works) map[string]interface{} {
	maps := make(map[string]interface{})
	maps["delete_timestamp"] = 0
	maps["is_open"] = 1
	if w.Username != "" {
		maps["username"] = w.Username
	}
	if w.WorksName != "" {
		maps["works_name"] = w.WorksName
	}
	if w.CatId > 0 {
		maps["cat_id"] = w.CatId
	}

	return maps
}
