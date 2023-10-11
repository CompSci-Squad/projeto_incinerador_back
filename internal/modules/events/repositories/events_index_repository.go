package repositories

import (
	"context"
	"goapi/internal/modules/events/entities"
	"goapi/pkg/pagination"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *EventsRepository) Index(limit, page int64, filter bson.D) (*[]entities.Event, int64, error) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	count, err := repo.eventsCollections.CountDocuments(ctxTimeout, filter)
	if err != nil {
		return nil, 0, err
	}

	paginate := pagination.NewMongoPaginate(limit, page, count)
	options := paginate.Options().
		SetSort(bson.D{{Key: "createdAt", Value: -1}}).
		SetProjection(bson.M{"_id": 1, "createdAt": 1, "updatedAt": 1, "name": 1, "description": 1, "date_time": 1})

	cursor, err := repo.eventsCollections.Find(ctxTimeout, filter, options)
	if err != nil {
		return nil, 0, err
	}

	events := []entities.Event{}
	if err := cursor.All(ctxTimeout, &events); err != nil {
		return nil, 0, err
	}

	return &events, count, nil
}
