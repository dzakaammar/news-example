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

	actualResult, _ := service.GetAllNews(filterBy)

	assert.Equal(t, excpectedResult, actualResult)
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

	actualResult, _ := service.GetAllNews(filterBy)

	assert.Equal(t, excpectedResult, actualResult)
}

func TestCreateNews(t *testing.T) {
	newsRepository := new(mocks.NewsRepository)
	topicRepository := new(mocks.TopicRepository)

	news := &domain.News{
		Title:   "Contoh",
		Status:  "draft",
		Content: "Contoh", /*
			Topics: []*domain.Topic{
				&domain.Topic{
					ID:   1,
					Name: "politik",
				},
				&domain.Topic{
					ID:   1,
					Name: "ekonomi",
				},
			}, */
	}
	newsRepository.On("Store", news, []string{"politik", "ekonomi"}).Return(nil)

	service := NewService(newsRepository, topicRepository)

	err := service.CreateNews("Contoh", "draft", "Contoh", []string{"politik", "ekonomi"})

	assert.Equal(t, nil, err)
}
