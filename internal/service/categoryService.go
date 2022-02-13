package service

import "designer-api/internal/models"

type Category struct {
	CatId           int
	CatName         string
	ParentId        int
	IsDirectory     int
	Level           int
	Path            string
	CreateTime      string
	UpdateTime      string
	DeleteTimestamp int

	PageNum  int
	PageSize int
}

func (cat *Category) ExistByName() (bool, error) {
	return models.ExistCategoryByName(cat.CatName)
}

func (cat *Category) Add() error {
	category := map[string]interface{}{
		"catName":  cat.CatName,
		"parentId": cat.ParentId,
	}

	if err := models.AddCategory(category); err != nil {
		return err
	}

	return nil
}

func (cat *Category) Edit() error {
	category := map[string]interface{}{
		"catName":  cat.CatName,
		"parentId": cat.ParentId,
	}

	if err := models.EditCategory(cat.CatId, category); err != nil {
		return err
	}

	return nil
}

func (cat *Category) GetAll() ([]models.Category, error) {
	category, err := models.GetCategory(cat.PageNum, cat.PageSize, cat.getMaps())
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (cat *Category) Delete() error {
	return models.DeleteWorks(cat.CatId)
}

func (cat *Category) ExistByID() (bool, error) {
	return models.ExistCategoryById(cat.CatId)
}

func (cat *Category) Count() (int64, error) {
	return models.GetCategoryTotal(cat.getMaps())
}

func (cat *Category) getMaps() map[string]interface{} {
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
