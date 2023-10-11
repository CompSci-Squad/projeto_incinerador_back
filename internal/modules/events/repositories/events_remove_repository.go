package repositories

import (
	"context"
	"fmt"
	"goapi/pkg/errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *EventsRepository) Delete(id string) error {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	mongoId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: mongoId}}
	opts := options.Delete().SetHint(bson.D{{Key: "_id", Value: 1}})

	_, err := repo.eventsCollections.DeleteOne(ctxTimeout, filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.NotFound(errors.Message{
				"error": true,
				"msg":   fmt.Sprintf("events with ID: %s is not found", id),
			})
		}

		return err
	}

	return nil
}
