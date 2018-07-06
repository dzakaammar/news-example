package service

import (
	"time"

	"github.com/dzakaammar/news-example/domain"

	"github.com/go-kit/kit/metrics"
)

type instrumentService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (i *instrumentService) GetAllNews(filterBy map[string][]string) ([]*domain.News, error) {
	defer func(begin time.Time) {
		i.requestCount.With("method", "get_all_news").Add(1)
		i.requestLatency.With("method", "get_all_news").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.Service.GetAllNews(filterBy)
}

func (i *instrumentService) CreateNews(title, content, status string, topics []string) error {
	defer func(begin time.Time) {
		i.requestCount.With("method", "create_news").Add(1)
		i.requestLatency.With("method", "create_news").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.Service.CreateNews(title, content, status, topics)
}

func (i *instrumentService) RemoveNews(id uint) error {
	defer func(begin time.Time) {
		i.requestCount.With("method", "remove_news").Add(1)
		i.requestLatency.With("method", "remove_news").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.Service.RemoveNews(id)
}

func (i *instrumentService) GetAllTopic() ([]*domain.Topic, error) {
	defer func(begin time.Time) {
		i.requestCount.With("method", "get_all_topic").Add(1)
		i.requestLatency.With("method", "get_all_topic").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.Service.GetAllTopic()
}

func (i *instrumentService) RemoveTopic(id uint) error {
	defer func(begin time.Time) {
		i.requestCount.With("method", "remove_topic").Add(1)
		i.requestLatency.With("method", "remove_topic").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.Service.RemoveTopic(id)
}
