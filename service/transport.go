package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/dzakaammar/news-example/repository"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

var (
	ErrBadRequest = errors.New("Bad Request")
)

func MakeHandler(s Service, logger log.Logger) http.Handler {

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getAllNewsHandler := kithttp.NewServer(
		makeGetAllNewsEndpoint(s),
		decodeGetAllNews,
		encodeResponse,
		opts...,
	)

	createNewsHandler := kithttp.NewServer(
		makeCreateNewsEndpoint(s),
		decodeCreateNews,
		encodeResponse,
		opts...,
	)

	removeNewsHandler := kithttp.NewServer(
		makeRemoveNewsEndpoint(s),
		decodeRemoveNews,
		encodeResponse,
		opts...,
	)

	getAllTopicsHandler := kithttp.NewServer(
		makeGetAllTopicsEndpoint(s),
		decodeGetAllTopicsNews,
		encodeResponse,
		opts...,
	)

	removeTopicHandler := kithttp.NewServer(
		makeRemoveTopicEndpoint(s),
		decodeRemoveTopic,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/news/", getAllNewsHandler).Methods("GET")
	r.Handle("/news/", createNewsHandler).Methods("POST")
	r.Handle("/news/{id}", removeNewsHandler).Methods("DELETE")
	r.Handle("/topics/", getAllTopicsHandler).Methods("GET")
	r.Handle("/topics/{id}", removeTopicHandler).Methods("DELETE")

	return r

}

func decodeGetAllNews(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		FilterBy map[string][]string `json:"filter_by,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, ErrBadRequest
	}

	return getAllNewsRequest{
		FilterBy: body.FilterBy,
	}, nil
}

func decodeCreateNews(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Title   string   `json:"title"`
		Status  string   `json:"status"`
		Content string   `json:"content"`
		Topics  []string `json:"topics,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, ErrBadRequest
	}

	return createNewsRequest{
		Title:   body.Title,
		Status:  body.Status,
		Content: body.Content,
		Topics:  body.Topics,
	}, nil
}

func decodeRemoveNews(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRequest
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, ErrBadRequest
	}

	return removeNewsRequest{
		ID: uint(intID),
	}, nil
}

func decodeGetAllTopicsNews(_ context.Context, r *http.Request) (interface{}, error) {
	return getAllTopicsRequest{}, nil
}

func decodeRemoveTopic(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRequest
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, ErrBadRequest
	}

	return removeTopicRequest{
		ID: uint(intID),
	}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	code := http.StatusOK
	switch err {
	case repository.NotFoundErr:
		w.WriteHeader(http.StatusNotFound)
		code = http.StatusNotFound
	case repository.CannotRetrieveErr:
		w.WriteHeader(http.StatusInternalServerError)
		code = http.StatusInternalServerError
	case repository.CannotStoreErr:
		w.WriteHeader(http.StatusInternalServerError)
		code = http.StatusInternalServerError
	case ErrBadRequest:
		w.WriteHeader(http.StatusBadRequest)
		code = http.StatusBadRequest
	default:
		w.WriteHeader(http.StatusBadRequest)
		code = http.StatusBadRequest
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":  code,
		"error": err.Error(),
	})
}
