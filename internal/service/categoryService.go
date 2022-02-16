package service

import "designer-api/internal/models"

type CategoryService struct {
	CategoryModel models.CategoryModel
}

type Category struct {
	models.Category

	PageNum  int
	PageSize int
}

func (service *CategoryService) ExistByName(name string) (bool, error) {
	return service.CategoryModel.ExistCategoryByName(name)
}

func (service *CategoryService) Add(cat *Category) error {
	categoryData := models.Category{}
	categoryData.CatName = cat.CatName
	categoryData.ParentId = cat.ParentId

	if err := service.CategoryModel.AddCategory(&categoryData); err != nil {
		return err
	}

	return nil
}

func (service *CategoryService) Edit(cat *Category) error {
	categoryData := models.Category{}
	categoryData.CatName = cat.CatName
	categoryData.ParentId = cat.ParentId

	if err := service.CategoryModel.EditCategory(cat.CatId, categoryData); err != nil {
		return err
	}

	return nil
}

func (service *CategoryService) GetAll(cat *Category) ([]models.Category, error) {
	category, err := service.CategoryModel.GetCategory(cat.PageNum, cat.PageSize, service.getMaps(cat))
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (service *CategoryService) Delete(catId int) error {
	return service.CategoryModel.DeleteCategory(catId)
}

func (service *CategoryService) ExistByID(catId int) (bool, error) {
	return service.CategoryModel.ExistCategoryById(catId)
}

func (service *CategoryService) Count(cat *Category) (int64, error) {
	return service.CategoryModel.GetCategoryTotal(service.getMaps(cat))
}

func (service *CategoryService) getMaps(cat *Category) map[string]interface{} {
	maps := make(map[string]interface{})
	maps["delete_timestamp"] = 0
	if cat.CatName != "" {
		maps["cat_name"] = cat.CatName
	}
	if cat.ParentId != -1 {
		maps["parent_id"] = cat.ParentId
	}

	return maps
}
