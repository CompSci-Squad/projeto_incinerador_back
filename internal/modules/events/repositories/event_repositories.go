package repositories

import (
	"goapi/infrastructure/database"
	"goapi/internal/modules/events/entities"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventsRepositoriesImpl interface {
	Index(limit, page int64, filter bson.D) (*[]entities.Event, int64, error)
	Store(b *entities.Event) (*entities.Event, error)
	Show(id uuid.UUID) (*entities.Event, error)
	Update(id *uuid.UUID, payload *entities.Event) (*entities.Event, error)
	Delete(id uuid.UUID) error
}

type EventsRepository struct {
	eventsCollections *mongo.Collection
}

func NewEventsRepository() (EventsRepositoriesImpl, error) {
	db, err := database.OpenDBConnection("mongodb")
	if err != nil {
		// Return status 500 and database connection error.
		return nil, err
	}

	database := db.Mongo.Database("app")
	eventsCollection := database.Collection("events")

	return &EventsRepository{eventsCollection}, nil
}
