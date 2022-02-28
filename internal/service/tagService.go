package service

import "designer-api/internal/models"

type TagService struct {
	TagsModel     models.TagsModel
	WorksTagModel models.WorksTagModel
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
	_ = service.WorksTagModel.UpdateTag(tag.TagId, tag.TagName)

	return nil
}

func (service *TagService) GetAll(tag *Tags) ([]models.Tags, error) {
	tagList, err := service.TagsModel.GetTagList(tag.OrderBy, service.getMaps(tag))
	if err != nil {
		return nil, err
	}

	return tagList, nil
}

func (service *TagService) GetByTag(tag *Tags) ([]models.Tags, error) {
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
	err := service.TagsModel.DeleteTag(tagId, userId)
	if err != nil {
		return err
	}
	err = service.WorksTagModel.DeleteByTag(tagId)
	if err != nil {
		return err
	}

	return nil
}

func (service *TagService) Count(tag *Tags) (int64, error) {
	return service.TagsModel.GetTagTotal(service.getMaps(tag))
}

func (service *TagService) getMaps(tag *Tags) map[string]interface{} {
	maps := make(map[string]interface{})
	if tag.Delete > 0 {
		maps["tags.delete_timestamp"] = tag.Delete
	} else {
		maps["tags.delete_timestamp"] = 0
	}
	if tag.TagId != 0 {
		maps["tags.tag_id"] = tag.TagId
	}
	if tag.UserId != 0 {
		maps["tags.user_id"] = tag.UserId
	}
	if tag.TagName != "" {
		maps["tags.tag_name"] = tag.TagName
	}

	return maps
}
