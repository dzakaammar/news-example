package service

import (
	"github.com/dzakaammar/news-example/domain"
)

type Service interface {
	GetAllNews(FilterBy map[string][]string) ([]*domain.News, error)
	CreateNews(Title, Content, Status string, Topics []string) error
	RemoveNews(ID uint) error

	GetAllTopic() ([]*domain.Topic, error)
	RemoveTopic(ID uint) error
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type service struct {
	newsRepository  domain.NewsRepository
	topicRepository domain.TopicRepository
}

func NewService(newsRepository domain.NewsRepository, topicRepository domain.TopicRepository) Service {
	return &service{newsRepository, topicRepository}
}

var filter = []string{"status", "topics"}

func (s *service) GetAllNews(filterBy map[string][]string) ([]*domain.News, error) {
	if len(filterBy) > 0 {
		for keys, _ := range filterBy {
			if b := stringInSlice(keys, filter); !b {
				return nil, ErrBadRequest
			}
		}
	}

	news, err := s.newsRepository.List(filterBy)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (s *service) CreateNews(title, content, status string, topics []string) error {
	news := domain.NewNews(title, content, status)

	return s.newsRepository.Store(news, topics)
}

func (s *service) RemoveNews(id uint) error {
	news, err := s.newsRepository.Find(id)
	if err != nil {
		return err
	}

	return s.newsRepository.Remove(news)
}

func (s *service) GetAllTopic() ([]*domain.Topic, error) {
	topics, err := s.topicRepository.List()
	if err != nil {
		return nil, err
	}

	return topics, nil
}

func (s *service) RemoveTopic(id uint) error {
	topic, err := s.topicRepository.Find(id)
	if err != nil {
		return err
	}

	return s.topicRepository.Remove(topic)
}
