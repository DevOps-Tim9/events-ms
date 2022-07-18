package service

import (
	"events-ms/src/dto"
	"events-ms/src/model"
	"events-ms/src/repository"
	"events-ms/src/utils"
	"fmt"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EventServiceIntegrationTestSuite struct {
	suite.Suite
	service EventService
	db      *gorm.DB
	events  []model.Event
}

func (suite *EventServiceIntegrationTestSuite) SetupSuite() {
	host := os.Getenv("DATABASE_DOMAIN")
	user := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	name := os.Getenv("DATABASE_SCHEMA")
	port := os.Getenv("DATABASE_PORT")

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		name,
		port,
	)
	db, _ := gorm.Open("postgres", connectionString)

	db.AutoMigrate(model.Event{})

	eventRepository := repository.EventRepository{Database: db}

	suite.db = db

	suite.service = EventService{
		EventRepo: &eventRepository,
		Logger:    utils.Logger(),
	}

	suite.events = []model.Event{
		{
			Message:   "New system event integration test",
			Timestamp: "2017-07-04 00:47:20",
		},
	}

	tx := suite.db.Begin()

	tx.Create(&suite.events[0])

	tx.Commit()
}

func TestEventServiceIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(EventServiceIntegrationTestSuite))
}

func (suite *EventServiceIntegrationTestSuite) TestIntegrationEventService_GetAll_EventsExist() {
	events, err := suite.service.GetAll()

	assert.NotNil(suite.T(), events)
	assert.Equal(suite.T(), 1, len(events))
	assert.Nil(suite.T(), err)
}

func (suite *EventServiceIntegrationTestSuite) TestIntegrationEventService_Add_Pass() {
	e := dto.EventRequestDTO{
		Message:   "New system event integration test",
		Timestamp: "2017-07-04 00:47:20",
	}

	responseDto, err := suite.service.Add(&e)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), responseDto)
	assert.Equal(suite.T(), e.Message, responseDto.Message)
	assert.Equal(suite.T(), e.Timestamp, responseDto.Timestamp)

	suite.service.EventRepo.Delete(responseDto.ID)
}

func (suite *EventServiceIntegrationTestSuite) TestIntegrationEventService_Add_Fail_RequiredPropertyMissing() {
	e := dto.EventRequestDTO{
		Message: "New system event integration test",
	}

	responseDto, err := suite.service.Add(&e)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), responseDto)
}
