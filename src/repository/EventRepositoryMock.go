package repository

import (
	"events-ms/src/model"

	"github.com/stretchr/testify/mock"
)

type EventRepositoryMock struct {
	mock.Mock
}

func (repo *EventRepositoryMock) Add(e model.Event) (model.Event, error) {
	args := repo.Called(e)
	if args.Get(1) == nil {
		return args.Get(0).(model.Event), nil
	}
	return args.Get(0).(model.Event), args.Get(1).(error)
}

func (repo *EventRepositoryMock) GetAll() ([]*model.Event, error) {
	args := repo.Called()
	if args.Get(1) == nil {
		return args.Get(0).([]*model.Event), nil
	}
	return args.Get(0).([]*model.Event), args.Get(1).(error)
}
