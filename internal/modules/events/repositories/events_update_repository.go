package repositories

import (
	"context"
	"fmt"
	"goapi/internal/modules/events/entities"
	"goapi/pkg/errors"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *EventsRepository) Update(id *uuid.UUID, payload *entities.Event) (*entities.Event, error) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	updateAt := time.Now()
	payload.UpdatedAt = &updateAt
	newEvent := entities.Event{}

	filter := bson.D{{Key: "id", Value: id}}
	update := bson.M{"$set": (*payload).Value()}

	err := repo.eventsCollections.FindOneAndUpdate(ctxTimeout, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&newEvent)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NotFound(errors.Message{
				"error": true,
				"msg":   fmt.Sprintf("Event with ID: %s is not found", id),
			})
		}

		return nil, err
	}

	return &newEvent, nil
}
