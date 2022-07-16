package handler

import (
	"events-ms/src/dto"
	"events-ms/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type EventHandler struct {
	Service *service.EventService
	Logger  *logrus.Entry
}

func (handler *EventHandler) AddEvent(ctx *gin.Context) {
	var EventDTO dto.EventRequestDTO
	if err := ctx.ShouldBindJSON(&EventDTO); err != nil {
		handler.Logger.Debug(err.Error())
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	handler.Logger.Info("Adding new system event.")

	dto, err := handler.Service.Add(&EventDTO)
	if err != nil {
		handler.Logger.Debug(err.Error())
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto)
}

func (handler *EventHandler) GetAll(ctx *gin.Context) {
	handler.Logger.Info("Getting system events")

	offersDTO, err := handler.Service.GetAll()
	if err != nil {
		handler.Logger.Debug(err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, offersDTO)
}
