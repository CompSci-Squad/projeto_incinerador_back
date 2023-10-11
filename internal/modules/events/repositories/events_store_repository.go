package repositories

import (
	"context"
	"goapi/internal/modules/events/entities"
	"goapi/pkg/errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *EventsRepository) Store(data *entities.Event) (*entities.Event, error) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	value := data.Value()

	res, err := repo.eventsCollections.InsertOne(ctxTimeout, value)
	if err != nil {
		return nil, errors.BadRequest(errors.Message{
			"error": err.Error(),
			"msg":   "Can't create event",
		})
	}

	createdEvent := entities.Event{}
	if err := repo.eventsCollections.FindOne(ctxTimeout, bson.M{"_id": res.InsertedID}).Decode(&createdEvent); err != nil {
		return nil, errors.BadRequest(errors.Message{
			"error": err.Error(),
			"msg":   "Can't create event",
		})
	}

	return &createdEvent, nil
}
