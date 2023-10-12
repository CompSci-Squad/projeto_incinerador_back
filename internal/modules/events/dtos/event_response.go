package dtos

import (
	"goapi/internal/modules/events/entities"
	"goapi/pkg/pagination"
)

type EventIndexResponse struct {
	*pagination.Paginate[entities.Event]
}

type EventShowResponse struct {
	*entities.Event `json:",inline"`
}

type EventStoreResponse struct {
	*entities.Event `json:",inline"`
}

type EventUpdateResponse struct {
	*entities.Event `json:",inline"`
}
