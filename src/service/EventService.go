package service

import (
	"events-ms/src/dto"
	"events-ms/src/mapper"
	"events-ms/src/repository"
	"fmt"

	"github.com/sirupsen/logrus"
)

type EventService struct {
	EventRepo repository.IEventRepository
	Logger    *logrus.Entry
}

type IEventService interface {
	Add(*dto.EventRequestDTO) (*dto.EventResponseDTO, error)
	GetAll() ([]*dto.EventResponseDTO, error)
}

func NewEventService(EventRepository repository.IEventRepository, logger *logrus.Entry) IEventService {
	return &EventService{
		EventRepository,
		logger,
	}
}

func (service *EventService) Add(dto *dto.EventRequestDTO) (*dto.EventResponseDTO, error) {
	err := dto.Validate()
	if err != nil {
		service.Logger.Debug(err.Error())
		return nil, err
	}

	entity := mapper.EventRequestDTOToEvent(dto)

	service.Logger.Info("Adding new system event in database")

	addedEntity, err := service.EventRepo.Add(*entity)
	if err != nil {
		service.Logger.Debug(err.Error())
		return nil, err
	}

	service.Logger.Info(fmt.Sprintf("Successfully added new system event in database with id %d", addedEntity.ID))
	return mapper.EventToEventResponseDTO(&addedEntity), nil
}

func (service *EventService) GetAll() ([]*dto.EventResponseDTO, error) {
	service.Logger.Info("Getting system events from database for company ")
	offers, err := service.EventRepo.GetAll()

	if err != nil {
		service.Logger.Debug(err.Error())
		return nil, err
	}

	res := make([]*dto.EventResponseDTO, len(offers))
	for i := 0; i < len(offers); i++ {
		res[i] = mapper.EventToEventResponseDTO(offers[i])
	}

	service.Logger.Info("Successfully got system events from database")
	return res, nil
}
