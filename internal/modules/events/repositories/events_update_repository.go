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

func (repo *EventsRepository) Update(id string, payload *entities.Event) (*entities.Event, error) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	updateAt := time.Now()
	payload.UpdatedAt = &updateAt

	mongoId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: mongoId}}
	update := bson.D{{Key: "$set", Value: payload.Value()}}

	_, err := repo.eventsCollections.UpdateOne(ctxTimeout, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NotFound(errors.Message{
				"error": true,
				"msg":   fmt.Sprintf("Documents with ID: %s is not found", id),
			})
		}

		return nil, err
	}

	err = repo.eventsCollections.FindOne(ctxTimeout, filter).Decode(payload)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NotFound(errors.Message{
				"error": true,
				"msg":   fmt.Sprintf("Documents with ID: %s is not found", id),
			})
		}

		return nil, err
	}

	return payload, nil
}
