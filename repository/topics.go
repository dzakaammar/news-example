package repository

import (
	"github.com/dzakaammar/news-example/domain"
	"github.com/jinzhu/gorm"
)

type topicRepository struct {
	*gorm.DB
}

func NewTopicRepository(db *gorm.DB) domain.TopicRepository {
	return &topicRepository{db}
}

func (t *topicRepository) List() ([]*domain.Topic, error) {
	var topics []*domain.Topic

	if err := t.DB.Preload("News").Find(&topics).Error; err != nil {
		return nil, CannotRetrieveErr
	}

	return topics, nil
}

func (t *topicRepository) Find(ID uint) (*domain.Topic, error) {
	topic := &domain.Topic{}

	if notFound := t.DB.Where("id = ?", ID).First(&topic).RecordNotFound(); notFound {
		return nil, NotFoundErr
	}

	return topic, nil
}

func (t *topicRepository) Store(topic *domain.Topic) error {
	if err := t.DB.Create(&topic).Error; err != nil {
		return CannotStoreErr
	}

	return nil
}

func (t *topicRepository) Remove(topic *domain.Topic) error {
	if err := t.DB.Delete(&topic).Error; err != nil {
		return CannotRemoveErr
	}
	return nil
}
