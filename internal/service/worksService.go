package service

import (
	"designer-api/internal/models"
	"encoding/json"
	"github.com/unknwon/com"
	"strconv"
	"strings"
)

type WorksService struct {
	WorksModel    models.WorksModel
	WorksTagModel models.WorksTagModel
	FavoriteModel models.FavoriteModel
}

type Works struct {
	models.Works

	TagId      string
	TagName    string
	OrderBy    string
	IsFavorite int
	Delete     int
	PageNum    int
	PageSize   int
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
	worksData.IsOpen = w.IsOpen
	worksData.WorksLink = w.WorksLink
	worksData.WorksType = w.WorksType
	worksData.WorksDescription = w.WorksDescription
	worksData.Remark = w.Remark

	worksId, err := service.WorksModel.AddWorks(&worksData)
	if err != nil {
		return err
	}

	var tagMap []models.WorksTag
	var nameMap []string
	idMap := strings.Split(w.TagId, ",")
	if w.TagName != "" {
		nameMap = strings.Split(w.TagName, ",")
	}

	for k, v := range idMap {
		var tag models.WorksTag
		tag.TagId = com.StrTo(v).MustInt()
		tag.TagName = ""
		if len(nameMap)-k > 0 && nameMap[k] != "" {
			tag.TagName = nameMap[k]
		}
		tag.WorksId = worksId
		tag.WorksName = w.WorksName
		tagMap = append(tagMap, tag)
	}
	_ = service.WorksTagModel.BatchAdd(tagMap)

	return nil
}

func (service *WorksService) Edit(w *Works) error {

	worksData := models.Works{}
	worksData.WorksName = w.WorksName
	worksData.State = w.State
	worksData.IsOpen = w.IsOpen
	worksData.WorksLink = w.WorksLink
	worksData.WorksType = w.WorksType
	worksData.WorksDescription = w.WorksDescription
	worksData.Remark = w.Remark

	if err := service.WorksModel.EditWorks(w.WorksId, worksData); err != nil {
		return err
	}

	if err := service.WorksTagModel.Delete(w.WorksId); err != nil {
		return err
	}

	var tagMap []models.WorksTag
	var nameMap []string
	idMap := strings.Split(w.TagId, ",")
	if w.TagName != "" {
		nameMap = strings.Split(w.TagName, ",")
	}

	for k, v := range idMap {
		var tag models.WorksTag
		tag.TagId, _ = strconv.Atoi(v)
		tag.TagName = ""
		if len(nameMap)-k > 0 && nameMap[k] != "" {
			tag.TagName = nameMap[k]
		}
		tag.WorksId = w.WorksId
		tag.WorksName = w.WorksName
		tagMap = append(tagMap, tag)
	}
	_ = service.WorksTagModel.BatchAdd(tagMap)

	return nil
}

func (service *WorksService) GetAll(w *Works) ([]models.Works, error) {
	maps := service.getMaps(w)
	tagIdMap := service.getIdMap(w.TagId)
	if len(tagIdMap) > 0 {
		maps["tagIdMap"] = tagIdMap
	}

	data, err := service.WorksModel.GetWorks(w.PageNum, w.PageSize, maps, w.OrderBy)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *WorksService) Get(worksId int, userId int) (map[string]interface{}, error) {
	var mapData map[string]interface{}
	works, err := service.WorksModel.GetOneWorks(worksId)
	if err != nil {
		return mapData, err
	}

	data, _ := json.Marshal(&works)
	_ = json.Unmarshal(data, &mapData)

	mapData["isFavorite"] = 0
	if userId > 0 && works.WorksId > 0 {
		isFavorite := service.FavoriteModel.IsFavorite(userId, worksId)
		if isFavorite {
			mapData["isFavorite"] = 1
		}
	}

	return mapData, nil
}

func (service *WorksService) Delete(worksId int, userId int) error {
	return service.WorksModel.DeleteWorks(worksId, userId)
}

func (service *WorksService) TrueDelete(worksId int, userId int) error {
	return service.WorksModel.Delete(worksId, userId)
}

func (service *WorksService) ExistByID(worksId int) (bool, error) {
	return service.WorksModel.ExistWorksById(worksId)
}

func (service *WorksService) Count(works *Works) (int64, error) {
	maps := service.getMaps(works)
	tagIdMap := service.getIdMap(works.TagId)
	if len(tagIdMap) > 0 {
		maps["tagIdMap"] = tagIdMap
	}
	return service.WorksModel.GetWorksTotal(maps)
}

func (service *WorksService) Increment(worksId int, field string) error {
	return service.WorksModel.Increment(worksId, field)
}

func (service *WorksService) Decrement(worksId int, field string) error {
	return service.WorksModel.Decrement(worksId, field)
}

func (service *WorksService) getMaps(w *Works) map[string]interface{} {
	maps := make(map[string]interface{})
	maps["works.is_open"] = w.IsOpen

	if w.Username != "" {
		maps["works.username"] = w.Username
	}
	if w.UserId != 0 {
		maps["works.user_id"] = w.UserId
	}
	if w.WorksName != "" {
		maps["works.works_name"] = w.WorksName
	}
	if w.Delete != 0 {
		maps["delete_timestamp"] = w.Delete
	} else {
		maps["delete_timestamp"] = 0
	}

	return maps
}

func (service *WorksService) getIdMap(tagId string) []int {
	var tagIdMap []int
	if tagId != "" {
		idMap := strings.Split(tagId, ",")
		for i := 0; i < len(idMap); i++ {
			tagIdMap = append(tagIdMap, com.StrTo(idMap[i]).MustInt())
		}
	}

	return tagIdMap
}
