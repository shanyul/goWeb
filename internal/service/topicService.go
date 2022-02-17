package service

import "designer-api/internal/models"

type Topic struct {
	models.Topic

	PageNum  int
	PageSize int
}

type TopicService struct {
	TopicModel models.TopicModel
	WorksModel models.WorksModel
}

func (service *TopicService) Add(topic *Topic) error {
	topicData := models.Topic{}
	topicData.UserId = topic.UserId
	topicData.Username = topic.Username
	topicData.WorksId = topic.WorksId
	topicData.Content = topic.Content
	topicData.RelationId = topic.RelationId

	if err := service.TopicModel.AddTopic(&topicData); err != nil {
		return err
	}

	field := "topic_num"
	_ = service.WorksModel.Increment(topic.WorksId, field)

	return nil
}

func (service *TopicService) GetAll(topic *Topic) ([]models.Topic, error) {
	data, err := service.TopicModel.GetTopic(topic.PageNum, topic.PageSize, service.getMaps(topic))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *TopicService) Get(id int) (*models.Topic, error) {
	data, err := service.TopicModel.GetOneTopic(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *TopicService) Delete(topicId int, userId int) error {
	topicData, err := service.Get(topicId)
	if err != nil {
		return err
	}

	if topicData.DeleteTimestamp == 0 {
		if err := service.TopicModel.DeleteTopic(topicId, userId); err != nil {
			return err
		}

		// 扣减评论数
		field := "topic_num"
		_ = service.WorksModel.Decrement(topicData.WorksId, field)
	}

	return nil
}

func (service *TopicService) Count(topic *Topic) (int64, error) {
	return service.TopicModel.GetTopicTotal(service.getMaps(topic))
}

func (service *TopicService) getMaps(topic *Topic) map[string]interface{} {
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
