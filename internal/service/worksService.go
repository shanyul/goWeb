package service

import (
	"designer-api/internal/models"
)

type Works struct {
	WorksId          int
	WorksName        string
	UserId           int
	Username         string
	State            int
	CatId            int
	WorksLink        string
	WorksType        int
	WorksDescription string
	Remark           string
	IsOpen           int
	FavoriteNum      int
	ViewNum          int
	TopicNum         int
	CreateTime       string
	UpdateTime       string
	DeleteTimestamp  int

	PageNum  int
	PageSize int
}

func (w *Works) ExistByName() (bool, error) {
	return models.ExistWorksByName(w.WorksName)
}

func (w *Works) Add() error {
	works := map[string]interface{}{
		"worksName":        w.WorksName,
		"userId":           w.UserId,
		"username":         w.Username,
		"state":            w.State,
		"catId":            w.CatId,
		"worksLink":        w.WorksLink,
		"worksType":        w.WorksType,
		"worksDescription": w.WorksDescription,
		"remark":           w.Remark,
	}

	if err := models.AddWorks(works); err != nil {
		return err
	}

	return nil
}

func (w *Works) Edit() error {
	works := map[string]interface{}{
		"worksName":        w.WorksName,
		"state":            w.State,
		"catId":            w.CatId,
		"worksLink":        w.WorksLink,
		"worksType":        w.WorksType,
		"worksDescription": w.WorksDescription,
		"remark":           w.Remark,
	}

	if err := models.EditWorks(w.WorksId, works); err != nil {
		return err
	}

	return nil
}

func (w *Works) GetAll(orderBy string) ([]models.Works, error) {
	works, err := models.GetWorks(w.PageNum, w.PageSize, w.getMaps(), orderBy)
	if err != nil {
		return nil, err
	}

	return works, nil
}

func (w *Works) Get() (*models.Works, error) {
	works, err := models.GetOneWorks(w.WorksId)
	if err != nil {
		return nil, err
	}

	return works, nil
}

func (w *Works) Delete() error {
	return models.DeleteWorks(w.WorksId, w.UserId)
}

func (w *Works) ExistByID() (bool, error) {
	return models.ExistWorksById(w.WorksId)
}

func (w *Works) Count() (int64, error) {
	return models.GetWorksTotal(w.getMaps())
}

func (w *Works) Increment(field string) error {
	return models.Increment(w.WorksId, field)
}

func (w *Works) Decrement(field string) error {
	return models.Decrement(w.WorksId, field)
}

func (w *Works) getMaps() map[string]interface{} {
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
