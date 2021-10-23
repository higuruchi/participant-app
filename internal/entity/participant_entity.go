package entity

import (
	"errors"
)

type participantEntity struct {
	id string
	name string
}

type ParticipantEntity interface {
	GetID() string
	GetName() string
}

func NewParticipantEntity(id string, name string) (ParticipantEntity ,error) {
	if id == "" {
		return nil, errors.New("id is required or invalid")
	}

	if (name == "") {
		return nil, errors.New("name is required or invalid")
	}

	return &participantEntity{id: id, name: name}, nil
}

func (participantEntity *participantEntity) GetID() string {
	return participantEntity.id
}

func (participantEntity *participantEntity) GetName() string {
	return participantEntity.name
}