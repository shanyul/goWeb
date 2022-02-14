package models

import (
	"gorm.io/gorm"
)

type Topic struct {
	TopicId         int    `gorm:"primary_key" column:"topic_id" json:"topicId"`
	UserId          int    `column:"user_id" json:"userId"`
	UserName        string `column:"user_name" json:"userName"`
	WorksId         int    `column:"works_id" json:"worksId"`
	Content         string `column:"content" json:"content"`
	RelationId      int    `column:"relation_id" json:"relationId"`
	CreateTime      string `column:"create_time" json:"createTime"`
	DeleteTimestamp int    `column:"delete_timestamp" json:"deleteTimestamp"`
}

// 自定义表名
func (Topic) TableName() string {
	return "topic"
}

// 获取作品
func GetTopic(pageNum int, pageSize int, maps interface{}) ([]Topic, error) {
	var (
		topic []Topic
		err   error
	)

	if pageSize > 0 && pageNum > 0 {
		err = dbHandle.Where(maps).Offset(pageNum).Limit(pageSize).Find(&topic).Error
	} else {
		err = dbHandle.Where(maps).Find(&topic).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return topic, nil
}

func AddTopic(data map[string]interface{}) error {
	topic := Topic{
		UserId:     data["userId"].(int),
		UserName:   data["userName"].(string),
		WorksId:    data["worksId"].(int),
		Content:    data["content"].(string),
		RelationId: data["relationId"].(int),
	}

	if err := dbHandle.Select(
		"UserId",
		"UserName",
		"WorksId",
		"Content",
		"RelationId",
	).Create(&topic).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTopic(id int) error {
	if err := dbHandle.Where("topic_id = ?", id).Delete(Topic{}).Error; err != nil {
		return err
	}

	return nil
}
