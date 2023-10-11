package repositories

import (
	"context"
	"fmt"
	"goapi/internal/modules/events/entities"
	"goapi/pkg/errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *EventsRepository) Show(id string) (*entities.Event, error) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}

	data := entities.Event{}
	err := repo.eventsCollections.FindOne(ctxTimeout, filter).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NotFound(errors.Message{
				"error": true,
				"msg":   fmt.Sprintf("documents with the given ID: %s is not found", id),
			})
		}

		return nil, err
	}

	return &data, nil
}
