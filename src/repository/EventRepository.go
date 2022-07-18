package repository

import (
	"errors"
	"events-ms/src/model"

	"github.com/jinzhu/gorm"
)

type IEventRepository interface {
	Add(model.Event) (model.Event, error)
	GetAll() ([]*model.Event, error)
	Delete(int) error
}

func NewEventRepository(database *gorm.DB) IEventRepository {
	return &EventRepository{
		database,
	}
}

type EventRepository struct {
	Database *gorm.DB
}

func (repo *EventRepository) Add(event model.Event) (model.Event, error) {
	err := repo.Database.Save(&event).Error

	return event, err
}

func (repo *EventRepository) GetAll() ([]*model.Event, error) {
	var events = []*model.Event{}
	if result := repo.Database.Find(&events); result.Error != nil {
		return nil, errors.New("Error happened during retrieving system events")
	}

	return events, nil
}

func (repo *EventRepository) Delete(id int) error {
	result := repo.Database.Delete(&model.Event{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
