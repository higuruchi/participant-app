package repository

import (
	"net"
	"github.com/higuruchi/participant-app/internal/usecase/model"
)

type ParticipantsRepository interface {
	GetParticipants(int, int, int) ([]model.Participant, int, error)
	SaveParticipant(int, int, int, int, int, int, net.HardwareAddr) error
}