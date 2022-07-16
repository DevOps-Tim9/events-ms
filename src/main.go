package main

import (
	"events-ms/src/handler"
	"events-ms/src/model"
	"events-ms/src/repository"
	"events-ms/src/service"
	"events-ms/src/utils"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
)

var db *gorm.DB
var err error

func initDB() (*gorm.DB, error) {
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

	db, _ = gorm.Open("postgres", connectionString)

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(model.Event{})
	return db, err
}

func initEventRepo(database *gorm.DB) *repository.EventRepository {
	return &repository.EventRepository{Database: database}
}

func initEventService(repo *repository.EventRepository) *service.EventService {
	return &service.EventService{EventRepo: repo, Logger: utils.Logger()}
}

func initEventHandler(service *service.EventService) *handler.EventHandler {
	return &handler.EventHandler{Service: service, Logger: utils.Logger()}
}

func handleEventFunc(handler *handler.EventHandler, router *gin.Engine) {
	router.POST("/events", handler.AddEvent)
	router.GET("/events", handler.GetAll)
}

func main() {
	logger := utils.Logger()

	logger.Info("Connecting with DB")

	database, _ := initDB()

	port := fmt.Sprintf(":%s", "9081")

	EventRepo := initEventRepo(database)
	EventService := initEventService(EventRepo)
	EventHandler := initEventHandler(EventService)

	router := gin.Default()

	handleEventFunc(EventHandler, router)

	logger.Info(fmt.Sprintf("Starting events-ms server on port %s", os.Getenv("SERVER_PORT")))
	http.ListenAndServe(port, cors.AllowAll().Handler(router))
}
