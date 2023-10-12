package baseentity

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        *uuid.UUID `bson:"id,omitempty" json:"id,omitempty" validate:"required,uuid"`
	CreatedAt *time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt *time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}
