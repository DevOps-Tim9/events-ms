package service

import (
	"events-ms/src/dto"
	"events-ms/src/model"
	"events-ms/src/repository"
	"events-ms/src/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EventServiceUnitTestsSuite struct {
	suite.Suite
	eventRepositoryMock *repository.EventRepositoryMock
	service             IEventService
}

func TestEventServiceUnitTestsSuite(t *testing.T) {
	suite.Run(t, new(EventServiceUnitTestsSuite))
}

func (suite *EventServiceUnitTestsSuite) SetupSuite() {
	suite.eventRepositoryMock = new(repository.EventRepositoryMock)
	suite.service = NewEventService(suite.eventRepositoryMock, utils.Logger())
}

func (suite *EventServiceUnitTestsSuite) TestNewEventService() {
	assert.NotNil(suite.T(), suite.service, "Service is nil")
}

func (suite *EventServiceUnitTestsSuite) TestEventService_Add_ValidDataProvided() {
	dto := dto.EventRequestDTO{
		Message:   "New system event test",
		Timestamp: "2017-07-04 00:47:20",
	}

	entity := model.Event{
		Message:   "New system event test",
		Timestamp: "2017-07-04 00:47:20",
	}

	savedEntity := model.Event{
		Message:   "New system event test",
		Timestamp: "2017-07-04 00:47:20",
		ID:        1,
	}

	suite.eventRepositoryMock.On("Add", entity).Return(savedEntity, nil).Once()

	returned, err := suite.service.Add(&dto)

	assert.Equal(suite.T(), dto.Message, returned.Message)
	assert.Equal(suite.T(), dto.Timestamp, returned.Timestamp)
	assert.Equal(suite.T(), nil, err)
}

func (suite *EventServiceUnitTestsSuite) TestEventService_Add_MissingRequiredField() {
	dto := dto.EventRequestDTO{
		Message: "New system event test",
	}

	returned, err := suite.service.Add(&dto)

	assert.Nil(suite.T(), returned)
	assert.NotNil(suite.T(), err)
}

func (suite *EventServiceUnitTestsSuite) TestEventService_GetAll_NoEventsReturnsEmpty() {
	suite.eventRepositoryMock.On("GetAll").Return([]*model.Event{}, nil).Once()

	events, err := suite.service.GetAll()

	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), 0, len(events))
}

func (suite *EventServiceUnitTestsSuite) TestEventService_Search_EventsExist() {
	event := model.Event{
		Message:   "New system event test",
		Timestamp: "2017-07-04 00:47:20",
		ID:        1,
	}
	var list []*model.Event
	list = append(list, &event)

	suite.eventRepositoryMock.On("GetAll").Return(list, nil).Once()

	events, err := suite.service.GetAll()

	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), len(list), len(events))
	for i := 0; i < len(events); i++ {
		assert.Equal(suite.T(), list[i].ID, events[i].ID)
		assert.Equal(suite.T(), list[i].Message, events[i].Message)
		assert.Equal(suite.T(), list[i].Timestamp, events[i].Timestamp)
	}
}
