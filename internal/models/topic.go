package models

import (
	"gorm.io/gorm"
	"time"
)

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
		Username:   data["username"].(string),
		WorksId:    data["worksId"].(int),
		Content:    data["content"].(string),
		RelationId: data["relationId"].(int),
	}

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
func GetTopicTotal(maps interface{}) (int64, error) {
	var count int64
	if err := dbHandle.Model(&Topic{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetOneTopic(id int) (*Topic, error) {
	var topic Topic
	if err := dbHandle.Where("topic_id = ?", id).First(&topic).Error; err != nil {
		return nil, err
	}

	return &topic, nil
}

func DeleteTopic(id int, userId int) error {
	maps := make(map[string]interface{})
	maps["delete_timestamp"] = time.Now().Unix()
	if err := dbHandle.Model(&Topic{}).Select("delete_timestamp").Where("topic_id = ? AND user_id = ?", id, userId).Updates(maps).Error; err != nil {
		return err
	}

	return nil
}
