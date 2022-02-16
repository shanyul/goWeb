package models

import (
	"gorm.io/gorm"
	"time"
)

type TopicModel struct{}

type Topic struct {
	TopicId         int    `gorm:"primary_key" column:"topic_id" json:"topicId"`
	UserId          int    `column:"user_id" json:"userId"`
	Username        string `column:"username" json:"username"`
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

// 获取评论
func (*TopicModel) GetTopic(pageNum int, pageSize int, maps interface{}) ([]Topic, error) {
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

func (*TopicModel) AddTopic(topic *Topic) error {
	if err := dbHandle.Select(
		"UserId",
		"Username",
		"WorksId",
		"Content",
		"RelationId",
	).Create(&topic).Error; err != nil {
		return err
	}

	return nil
}

// 获取总记录数
func (*TopicModel) GetTopicTotal(maps interface{}) (int64, error) {
	var count int64
	if err := dbHandle.Model(&Topic{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (*TopicModel) GetOneTopic(id int) (*Topic, error) {
	var topic Topic
	if err := dbHandle.Where("topic_id = ?", id).First(&topic).Error; err != nil {
		return nil, err
	}

	return &topic, nil
}

func (*TopicModel) DeleteTopic(id int, userId int) error {
	maps := make(map[string]interface{})
	maps["delete_timestamp"] = time.Now().Unix()
	if err := dbHandle.Model(&Topic{}).Select("delete_timestamp").Where("topic_id = ? AND user_id = ?", id, userId).Updates(maps).Error; err != nil {
		return err
	}

	return nil
}
