package services

import (
	"goapi/internal/modules/events/dtos"
	"goapi/internal/modules/events/repositories"

	"github.com/google/uuid"
)

type EventServicesImpl interface {
	Store(eventId *uuid.UUID, body *dtos.EventStoreRequest) (*dtos.EventStoreResponse, error)
	Index(payload *dtos.EventsIndexRequest) (*dtos.EventIndexResponse, error)
	Show(id uuid.UUID) (*dtos.EventShowResponse, error)
	Update(id *uuid.UUID, payload *dtos.EventUpdateRequest) (*dtos.EventUpdateResponse, error)
	Remove(id *uuid.UUID) error
}

type EventService struct {
	EventRepository repositories.EventsRepositoriesImpl
}

func NewEventSerive(eventRepository repositories.EventsRepository) {
	return &EventService{&eventRepository}
}
