package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dzakaammar/news-example/domain"

	"github.com/dzakaammar/news-example/mocks"
)

func TestGetAllNewsByStatus(t *testing.T) {
	newsRepository := new(mocks.NewsRepository)
	topicRepository := new(mocks.TopicRepository)

	filterBy := map[string][]string{
		"status": []string{"draft"},
	}

	news1 := domain.NewNews("Contoh", "draft", "Contoh")

	newsRepository.On("List", filterBy).Return([]*domain.News{news1}, nil)

	service := NewService(newsRepository, topicRepository)

	excpectedResult := []*domain.News{news1}

	actualResult, err := service.GetAllNews(filterBy)

	assert.Equal(t, excpectedResult, actualResult)
	assert.Empty(t, err)
}

func TestGetAllNewsByTopics(t *testing.T) {
	newsRepository := new(mocks.NewsRepository)
	topicRepository := new(mocks.TopicRepository)

	filterBy := map[string][]string{
		"topics": []string{"politik"},
	}

	news1 := &domain.News{Title: "Contoh", Status: "published", Content: "Contoh", Topics: []*domain.Topic{&domain.Topic{ID: 1, Name: "politik"}}}

	newsRepository.On("List", filterBy).Return([]*domain.News{news1}, nil)

	service := NewService(newsRepository, topicRepository)

	excpectedResult := []*domain.News{news1}

	actualResult, err := service.GetAllNews(filterBy)

	assert.Equal(t, excpectedResult, actualResult)

	assert.Empty(t, err)
}

func TestCreateNews(t *testing.T) {
	newsRepository := new(mocks.NewsRepository)
	topicRepository := new(mocks.TopicRepository)

	news := &domain.News{
		Title:   "Contoh",
		Status:  "draft",
		Content: "Contoh",
	}
	newsRepository.On("Store", news, []string{"politik", "ekonomi"}).Return(nil)

	service := NewService(newsRepository, topicRepository)

	err := service.CreateNews("Contoh", "draft", "Contoh", []string{"politik", "ekonomi"})

	assert.Empty(t, err)
}

func TestRemoveNews(t *testing.T) {
	newsRepository := new(mocks.NewsRepository)
	topicRepository := new(mocks.TopicRepository)

	news := &domain.News{
		ID:      1,
		Title:   "Contoh",
		Status:  "draft",
		Content: "Contoh",
	}

	newsRepository.On("Find", uint(1)).Return(news, nil)
	newsRepository.On("Remove", news).Return(nil)

	service := NewService(newsRepository, topicRepository)

	err := service.RemoveNews(1)
	assert.Empty(t, err)
}

func TestGetAllTopics(t *testing.T) {
	newsRepository := new(mocks.NewsRepository)
	topicRepository := new(mocks.TopicRepository)

	service := NewService(newsRepository, topicRepository)
	topics := []*domain.Topic{
		&domain.Topic{
			ID:   1,
			Name: "politik",
		},
		&domain.Topic{
			ID:   2,
			Name: "pendidikan",
		},
		&domain.Topic{
			ID:   3,
			Name: "budaya",
		},
	}
	topicRepository.On("List").Return(topics, nil)

	actual, err := service.GetAllTopic()

	assert.Equal(t, topics, actual)

	assert.Empty(t, err)

}

func TestRemoveTopic(t *testing.T) {
	newsRepository := new(mocks.NewsRepository)
	topicRepository := new(mocks.TopicRepository)

	service := NewService(newsRepository, topicRepository)

	topic := &domain.Topic{
		ID:   1,
		Name: "politik",
	}

	topicRepository.On("Find", uint(1)).Return(topic, nil)
	topicRepository.On("Remove", topic).Return(nil)

	err := service.RemoveTopic(1)

	assert.Empty(t, err)
}
