package service

import (
	"designer-api/internal/models"
)

type UserCategoryService struct {
	CategoryModel models.UserCategoryModel
}

type UserCategory struct {
	models.UserCategory
}

func (service *UserCategoryService) ExistByName(name string, userId int) (bool, error) {
	return service.CategoryModel.IsExist(name, userId)
}

func (service *UserCategoryService) Add(cat *UserCategory) error {

	userCategory := models.UserCategory{}
	userCategory.UcatName = cat.UcatName
	userCategory.UserId = cat.UserId
	userCategory.Username = cat.Username

	if err := service.CategoryModel.AddCategory(&userCategory); err != nil {
		return err
	}

	return nil
}

func (service *UserCategoryService) Edit(cat *UserCategory) error {

	userCategory := models.UserCategory{}
	userCategory.UcatName = cat.UcatName

	if err := service.CategoryModel.EditCategory(cat.UcatId, &userCategory); err != nil {
		return err
	}

	return nil
}

func (service *UserCategoryService) GetAll(cat *UserCategory) ([]models.UserCategory, error) {
	data, err := service.CategoryModel.GetCategoryList(service.getMaps(cat))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *UserCategoryService) Get(ucatId int, userId int) (models.UserCategory, error) {

	category, err := service.CategoryModel.GetCategory(ucatId, userId)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (service *UserCategoryService) Delete(ucatId int, userId int) error {
	return service.CategoryModel.DeleteCategory(ucatId, userId)
}

func (service *UserCategoryService) Count(category *UserCategory) (int64, error) {
	return service.CategoryModel.GetCategoryTotal(service.getMaps(category))
}

func (service *UserCategoryService) getMaps(cat *UserCategory) map[string]interface{} {
	maps := make(map[string]interface{})
	maps["delete_timestamp"] = 0
	if cat.UserId > 0 {
		maps["user_id"] = cat.UserId
	}
	if cat.UcatName != "" {
		maps["ucat_name"] = cat.UcatName
	}
	if cat.UcatId > 0 {
		maps["ucat_id"] = cat.UcatId
	}

	return maps
}
