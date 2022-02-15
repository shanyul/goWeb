package service

import "designer-api/internal/models"

type Topic struct {
	TopicId         int
	UserId          int
	Username        string
	WorksId         int
	Content         string
	RelationId      int
	CreateTime      string
	DeleteTimestamp int

	PageNum  int
	PageSize int
}

func (topic *Topic) Add() error {
	topicData := map[string]interface{}{
		"userId":     topic.UserId,
		"username":   topic.Username,
		"worksId":    topic.WorksId,
		"content":    topic.Content,
		"relationId": topic.RelationId,
	}

	if err := models.AddTopic(topicData); err != nil {
		return err
	}

	worksService := Works{
		WorksId: topic.WorksId,
	}

	field := "topic_num"
	_ = worksService.Increment(field)

	return nil
}

func (topic *Topic) GetAll() ([]models.Topic, error) {
	data, err := models.GetTopic(topic.PageNum, topic.PageSize, topic.getMaps())
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (topic *Topic) Get() (*models.Topic, error) {
	data, err := models.GetOneTopic(topic.TopicId)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (topic *Topic) Delete() error {
	topicData, err := topic.Get()
	if err != nil {
		return err
	}

	if topicData.DeleteTimestamp == 0 {
		if err := models.DeleteTopic(topic.TopicId, topic.UserId); err != nil {
			return err
		}

		// 扣减评论数
		worksService := Works{
			WorksId: topicData.WorksId,
		}
		field := "topic_num"
		_ = worksService.Decrement(field)
	}

	return nil
}

func (topic *Topic) Count() (int64, error) {
	return models.GetTopicTotal(topic.getMaps())
}

func (topic *Topic) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["delete_timestamp"] = 0
	if topic.WorksId > 0 {
		maps["works_id"] = topic.WorksId
	}
	if topic.UserId > 0 {
		maps["user_id"] = topic.UserId
	}

	return maps
}
