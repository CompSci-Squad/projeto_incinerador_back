package services

import (
	"goapi/internal/modules/events/dtos"
	"goapi/pkg/convert"

	"github.com/google/uuid"
)

func (serv *EventService) Show(id uuid.UUID) (*dtos.EventShowResponse, error) {
	event, err := serv.EventRepository.Show(id)

	if err != nil {
		return nil, err
	}

	res := dtos.EventShowResponse{}

	err = convert.ToStruct(*event, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
