package services

import "github.com/google/uuid"

func (serv *EventService) Delete(id *uuid.UUID) error {
	err := serv.EventRepository.Delete(*id)

	if err != nil {
		return err
	}

	return nil
}
