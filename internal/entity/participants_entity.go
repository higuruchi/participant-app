package entity

import (
	"errors"
)

type participantsEntity struct {
	M2 []ParticipantEntity
	M1 []ParticipantEntity
	B4 []ParticipantEntity
	B3 []ParticipantEntity
	B2 []ParticipantEntity
	B1 []ParticipantEntity
}

type ParticipantsEntity interface {
	SetParticipants(ParticipantEntity, grade) (ParticipantsEntity, error)
	GetParticipants(string) []ParticipantEntity
}

func NewParticipahtsEntity() ParticipantsEntity {
	return &participantsEntity{}
}

func (participantsEntity *participantsEntity) SetParticipants(
	participant ParticipantEntity,
	grade grade,
) (ParticipantsEntity, error) {
	switch grade {
	case M2:
		participantsEntity.M2 = append(participantsEntity.M2, participant)
	case M1:
		participantsEntity.M1 = append(participantsEntity.M1, participant)
	case B4:
		participantsEntity.B4 = append(participantsEntity.B4, participant)
	case B3:
		participantsEntity.B3 = append(participantsEntity.B3, participant)
	case B2:
		participantsEntity.B2 = append(participantsEntity.B2, participant)
	case B1:
		participantsEntity.B1 = append(participantsEntity.B1, participant)
	default:
		return participantsEntity, errors.New("grade is invalid")
	}
	return participantsEntity, nil
}

func (participantsEntity *participantsEntity) GetParticipants(grade string) []ParticipantEntity {
	switch grade {
	case "M2":
		return participantsEntity.M2
	case "M1":
		return participantsEntity.M1
	case "B4":
		return participantsEntity.B4
	case "B3":
		return participantsEntity.B3
	case "B2":
		return participantsEntity.B2
	case "B1":
		return participantsEntity.B1
	}
	return nil
}
