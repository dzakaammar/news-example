// Code generated by mockery v1.0.0
package mocks

import domain "github.com/dzakaammar/news-example/domain"
import mock "github.com/stretchr/testify/mock"

// NewsRepository is an autogenerated mock type for the NewsRepository type
type NewsRepository struct {
	mock.Mock
}

// Find provides a mock function with given fields: ID
func (_m *NewsRepository) Find(ID uint) (*domain.News, error) {
	ret := _m.Called(ID)

	var r0 *domain.News
	if rf, ok := ret.Get(0).(func(uint) *domain.News); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.News)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: FilterBy
func (_m *NewsRepository) List(FilterBy map[string][]string) ([]*domain.News, error) {
	ret := _m.Called(FilterBy)

	var r0 []*domain.News
	if rf, ok := ret.Get(0).(func(map[string][]string) []*domain.News); ok {
		r0 = rf(FilterBy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.News)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string][]string) error); ok {
		r1 = rf(FilterBy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: news
func (_m *NewsRepository) Remove(news *domain.News) error {
	ret := _m.Called(news)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.News) error); ok {
		r0 = rf(news)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store provides a mock function with given fields: news, Topics
func (_m *NewsRepository) Store(news *domain.News, Topics []string) error {
	ret := _m.Called(news, Topics)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.News, []string) error); ok {
		r0 = rf(news, Topics)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}