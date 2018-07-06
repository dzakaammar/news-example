package service

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type getAllNewsRequest struct {
	FilterBy map[string][]string
}

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func makeGetAllNewsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAllNewsRequest)
		data, err := s.GetAllNews(req.FilterBy)
		return response{
			Code:    http.StatusOK,
			Message: "Success",
			Data:    data,
		}, err
		//return data, err
	}
}

type createNewsRequest struct {
	Title   string
	Status  string
	Content string
	Topics  []string
}

func makeCreateNewsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createNewsRequest)
		err := s.CreateNews(req.Title, req.Content, req.Status, req.Topics)
		return response{
			Code:    http.StatusOK,
			Message: "Success",
		}, err
	}
}

type removeNewsRequest struct {
	ID uint
}

func makeRemoveNewsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(removeNewsRequest)
		err := s.RemoveNews(req.ID)
		return response{
			Code:    http.StatusOK,
			Message: "Success",
		}, err
	}
}

type getAllTopicsRequest struct{}

func makeGetAllTopicsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(getAllTopicsRequest)
		data, err := s.GetAllTopic()
		return response{
			Code:    http.StatusOK,
			Message: "Success",
			Data:    data,
		}, err
	}
}

type removeTopicRequest struct {
	ID uint
}

func makeRemoveTopicEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(removeTopicRequest)
		err := s.RemoveTopic(req.ID)
		return response{
			Code:    http.StatusOK,
			Message: "Success",
		}, err
	}
}
