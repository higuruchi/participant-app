package usecase

import (
	"fmt"
	"net"
	"errors"
	"github.com/higuruchi/participant-app/internal/entity"
	"github.com/higuruchi/participant-app/internal/usecase/repository"
)

type participantsUsecase struct {
	participantsRepository repository.ParticipantsRepository
}

type ParticipantsUsecase interface {
	GetParticipants(int, int, int) (entity.ParticipantsEntity, error)
	SaveParticipant(int, int, int, int, int, int, net.HardwareAddr) error
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

		grade, err := participantEntity.DistinguishGrade()
		if err != nil {
			return nil, fmt.Errorf("calling distinguishGrade: %w", err)
		}

		participantsEntity.SetParticipants(participantEntity, grade)
	}

	return participantsEntity, nil	
}

func (participantsUsecase *participantsUsecase) SaveParticipant(
	year int,
	month int,
	date int,
	hour int,
	minute int,
	second int,
	macaddress net.HardwareAddr,
) error {	
	err := participantsUsecase.participantsRepository.SaveParticipant(year, month, date, hour, minute, second, macaddress)
	if err != nil {
		return fmt.Errorf("calling participantsUsecase.participantsRepository.SaveParticipant %w", err)
	}

	return nil
}