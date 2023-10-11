package entities

import (
	"errors"
	"goapi/pkg/baseentity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	baseentity.Base `bson:", inline"`
	Name            string     `bson:"name"`
	Description     string     `bson:"description"`
	DateTime        *time.Time `bson:"date_time"`
}

func (b Event) Value() primitive.M {
	byte, _ := bson.Marshal(b)

	var updated bson.M
	bson.Unmarshal(byte, &updated)

	return updated
}

func (b Event) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return bson.Unmarshal(j, &b)
}
