package usecase

import (
	"fmt"
	"time"
	"errors"
	"strconv"
	"github.com/higuruchi/participant-app/internal/entity"
	// "github.com/higuruchi/participant-app/internal/usecase/model"
	"github.com/higuruchi/participant-app/internal/usecase/repository"
)

type participantsUsecase struct {
	participantsRepository repository.ParticipantsRepository
}

type ParticipantsUsecase interface {
	GetParticipants(int, int, int) (entity.ParticipantsEntity, error)
}

func NewParticipantsUsecase(
	participantsRepository repository.ParticipantsRepository,
) ParticipantsUsecase {
	return &participantsUsecase{
		participantsRepository: participantsRepository,
	}
}

func (participantsUsecase *participantsUsecase) GetParticipants(
	year int,
	month int,
	date int,
) (entity.ParticipantsEntity, error) {

	if year < 0 {
		return nil, errors.New("syntax error: not enough arguments or value is out of range")
	}

	if month < 1 || 12 < month {
		return nil, errors.New("syntax error; not enough arguments or value is out of range")
	}

	if date < 1 || 31 < date {
		return nil, errors.New("syntax error; not enough arguments or value is out of range")
	}

	participants, num, err := participantsUsecase.participantsRepository.GetParticipants(year, month, date)
	if err != nil {
		return nil, fmt.Errorf("calling participantsUsecase.participantsRepository.GetParticipants: %w", err)
	}

	participantsEntity := entity.NewParticipahtsEntity()

	for i := 0; i < num; i++ {
		if participants[i].ID == "unknown" {
			continue
		}

		participantEntity, err := entity.NewParticipantEntity(participants[i].ID, participants[i].Name)
		if err != nil {
			return nil, fmt.Errorf("calling NewParticipantEntity: %w", err)
		}

		grade, err := distinguishGrade(participantEntity.GetID())
		if err != nil {
			return nil, fmt.Errorf("calling distinguishGrade: %w", err)
		}

		participantsEntity.SetParticipants(participantEntity, grade)
	}

	return participantsEntity, nil	
}

// この部分汚いw
func distinguishGrade(id string) (string, error) {
	t := time.Now()
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