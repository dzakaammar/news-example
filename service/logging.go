package service

import (
	"time"

	"github.com/dzakaammar/news-example/domain"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (l *loggingService) GetAllNews(filterBy map[string][]string) (resp []*domain.News, err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"method", "get_all_news",
			"filterBy", filterBy,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return l.Service.GetAllNews(filterBy)
}

func (l *loggingService) CreateNews(title, content, status string, topics []string) (err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"method", "create_news",
			"title", title,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return l.Service.CreateNews(title, content, status, topics)
}

func (l *loggingService) RemoveNews(id uint) (err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"method", "remove_news",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return l.Service.RemoveNews(id)
}

func (l *loggingService) GetAllTopic() (resp []*domain.Topic, err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"method", "get_all_topic",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return l.Service.GetAllTopic()
}

func (l *loggingService) RemoveTopic(id uint) (err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"method", "remove_topic",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return l.Service.RemoveTopic(id)
}
