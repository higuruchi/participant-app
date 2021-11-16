package entity

import (
	"fmt"
	"time"
	"strconv"
	"errors"
)

type participantEntity struct {
	id string
	name string
}

type ParticipantEntity interface {
	GetID() string
	GetName() string
	DistinguishGrade() (string, error)
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

// この部分汚いw
func (participantEntity *participantEntity) DistinguishGrade() (string, error) {
	t := time.Now()
	id := participantEntity.id;
	nowYear := t.Year()
	admissionYear, err := strconv.Atoi(fmt.Sprintf("%d%s", 20, id[0:2]))
	if err != nil {
		return "", fmt.Errorf("calling strconv.Atoi: %w", err)
	}
	// 現在Gの学籍番号には対応していない
	// if id[2] == "G" {
	// 	switch nowYear-admissionYear {
	// 	case 0:
	// 		return "M1", nil
	// 	case 1:
	// 		return "M2", nil
	// 	default:
	// 		return "", errors.New("invalid year")

	// 	}
	// }

	switch nowYear-admissionYear {
	case 0:
		return "B1", nil
	case 1:
		return "B2", nil
	case 2:
		return "B3", nil
	case 3:
		return "B4", nil
	case 4:
		return "M1", nil
	case 5:
		return "M2", nil
	default:
		return "", errors.New("invalid year")
	}
}