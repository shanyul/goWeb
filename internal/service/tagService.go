package service

import "designer-api/internal/models"

type TagService struct {
	TagsModel models.TagsModel
}

type Tags struct {
	models.Tags
	Delete  int
	OrderBy string
}

func (service *TagService) ExistByName(name string, userId int) (bool, error) {
	return service.TagsModel.IsExist(name, userId)
}

func (service *TagService) Add(tag *Tags) error {
	data := models.Tags{}
	data.TagName = tag.TagName
	data.UserId = tag.UserId
	data.Username = tag.Username

	if err := service.TagsModel.AddTag(&data); err != nil {
		return err
	}

	return nil
}

func (service *TagService) Edit(tag *Tags) error {
	data := models.Tags{}
	data.TagName = tag.TagName

	if err := service.TagsModel.EditTag(tag.TagId, tag.UserId, &data); err != nil {
		return err
	}

	return nil
}

func (service *TagService) GetAll(tag *Tags) ([]models.Tags, error) {
	tagList, err := service.TagsModel.GetTagList(tag.OrderBy, service.getMaps(tag))
	if err != nil {
		return nil, err
	}

	return tagList, nil
}

func (service *TagService) Get(tagId int) (models.Tags, error) {
	tagData, err := service.TagsModel.GetTag(tagId)
	if err != nil {
		return tagData, err
	}

	return tagData, nil
}

func (service *TagService) Delete(tagId int, userId int) error {
	return service.TagsModel.DeleteTag(tagId, userId)
}

func (service *TagService) Count(tag *Tags) (int64, error) {
	return service.TagsModel.GetTagTotal(service.getMaps(tag))
}

func (service *TagService) getMaps(tag *Tags) map[string]interface{} {
	maps := make(map[string]interface{})
	if tag.Delete > 0 {
		maps["delete_timestamp"] = tag.Delete
	} else {
		maps["delete_timestamp"] = 0
	}
	if tag.TagId != 0 {
		maps["tag_id"] = tag.TagId
	}
	if tag.UserId != 0 {
		maps["user_id"] = tag.UserId
	}
	if tag.TagName != "" {
		maps["tag_name"] = tag.TagName
	}

	return maps
}
