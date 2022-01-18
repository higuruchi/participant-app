package repository

import (
	"fmt"
	"net"
	"github.com/higuruchi/participant-app/internal/usecase/model"
	"github.com/higuruchi/participant-app/internal/interfaceadapter/repository/worker"
)

type ParticipantsRepository struct {
	participantsGetter worker.DatabaseHandler
}

func NewParticipantsRepository(
	participantsGetter worker.DatabaseHandler,
) *ParticipantsRepository {
	return &ParticipantsRepository{participantsGetter: participantsGetter}
}

func (participantsRepository *ParticipantsRepository) GetParticipants(
	year int,
	month int,
	date int,
) ([]model.Participant, int, error) {
	var participants []model.Participant

	sql := `
	SELECT
		CASE 
			WHEN users.id IS NULL THEN "unknown"
			ELSE users.id
		END AS id,
		CASE
			WHEN users.name IS NULL THEN "unknown"
			ELSE users.name
		END AS name
	FROM packet_logs
	LEFT OUTER JOIN users
	ON packet_logs.mac_address = users.mac_address
	WHERE packet_logs.transit_time BETWEEN ? AND ?
	`

	from := fmt.Sprintf("%d-%d-%d 00:00:00", year, month, date)
	end := fmt.Sprintf("%d-%d-%d 23:59:59", year, month, date)

	rows, err := participantsRepository.participantsGetter.Query(sql, from, end)
	defer rows.Close()
	if err != nil {
		return nil, -1, fmt.Errorf("calling participantsRepository.participantsGetter.Query: %w", err)
	}

	for rows.Next() {
		var participant model.Participant

		err = rows.Scan(&participant.ID, &participant.Name)
		if err != nil {
			return nil, -1, fmt.Errorf("calling rows.Scan: %w", err)
		}
		participants = append(participants, participant)
	}

	return participants, len(participants), nil
}

func (participantsRepository *ParticipantsRepository) SaveParticipant(
	year int,
	month int,
	date int,
	hour int,
	minute int,
	second int,
	macaddress net.HardwareAddr,
) error {
	timestamp := fmt.Sprintf("%d-%d-%d %d:%d:%d", year, month, date, hour, minute, second)
	sql := `
	INSERT INTO packet_logs
	(transit_time, mac_address)
	VALUES (?, ?);
	`
	_, err := participantsRepository.participantsGetter.Execute(sql, timestamp, macaddress.String())
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("calling participantsRepository.participantsGetter.Execute: %w", err)
	}

	return nil
}