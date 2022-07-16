package mapper

import (
	"events-ms/src/dto"
	"events-ms/src/model"
)

func EventToEventResponseDTO(event *model.Event) *dto.EventResponseDTO {
	var e dto.EventResponseDTO

	e.ID = event.ID
	e.Timestamp = event.Timestamp
	e.Message = event.Message

	return &e
}

func EventRequestDTOToEvent(e *dto.EventRequestDTO) *model.Event {
	var event model.Event

	event.Timestamp = e.Timestamp
	event.Message = e.Message

	return &event
}
