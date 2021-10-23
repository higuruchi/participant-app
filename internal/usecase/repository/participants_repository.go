package repository

import (
	"github.com/higuruchi/participant-app/internal/usecase/model"
)

type ParticipantsRepository interface {
	GetParticipants(int, int, int) ([]model.Participant, int, error)
}