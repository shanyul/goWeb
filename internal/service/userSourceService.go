package service

import (
	"designer-api/internal/models"
)

type UserSourceService struct {
	SourceModel models.UserSourceModel
}

type UserSource struct {
	models.Source

	PageNum  int
	PageSize int
}

func (service *UserSourceService) Add(source *UserSource) error {

	sourceData := models.Source{}
	sourceData.UcatId = source.UcatId
	sourceData.UserId = source.UserId
	sourceData.Username = source.Username
	sourceData.UcatName = source.UcatName
	sourceData.Description = source.Description
	sourceData.Title = source.Title
	sourceData.Link = source.Link

	if err := service.SourceModel.AddSource(&sourceData); err != nil {
		return err
	}

	return nil
}

func (service *UserSourceService) Edit(source *UserSource) error {

	sourceData := models.Source{}
	sourceData.UcatId = source.UcatId
	sourceData.UcatName = source.UcatName
	sourceData.Description = source.Description
	sourceData.Title = source.Title

	if err := service.SourceModel.EditSource(source.SourceId, &sourceData); err != nil {
		return err
	}

	return nil
}

func (service *UserSourceService) GetAll(source *UserSource) ([]models.Source, error) {
	data, err := service.SourceModel.GetSourceList(source.PageNum, source.PageSize, service.getMaps(source))
	if err != nil {
		return data, err
	}

	return data, nil
}

func (service *UserSourceService) Get(sourceId int) (models.Source, error) {

	category, err := service.SourceModel.GetSource(sourceId)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (service *UserSourceService) Delete(sourceId int, userId int) error {
	return service.SourceModel.DeleteSource(sourceId, userId)
}

func (service *UserSourceService) Count(source *UserSource) (int64, error) {
	return service.SourceModel.GetSourceTotal(service.getMaps(source))
}

func (service *UserSourceService) getMaps(source *UserSource) map[string]interface{} {
	maps := make(map[string]interface{})
	maps["delete_timestamp"] = 0
	if source.UserId > 0 {
		maps["user_id"] = source.UserId
	}
	if source.UcatName != "" {
		maps["ucat_name"] = source.UcatName
	}
	if source.UcatId > 0 {
		maps["ucat_id"] = source.UcatId
	}
	if source.Title != "" {
		maps["title"] = source.Title
	}

	return maps
}
