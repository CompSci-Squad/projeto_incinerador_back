package repositories

import (
	"context"
	"fmt"
	"goapi/pkg/errors"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *EventsRepository) Delete(id uuid.UUID) error {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	filter := bson.D{{Key: "id", Value: id}}

	_, err := repo.eventsCollections.DeleteOne(ctxTimeout, filter)
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
