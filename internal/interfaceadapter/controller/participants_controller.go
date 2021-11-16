package controller

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/higuruchi/participant-app/internal/usecase"
	"github.com/higuruchi/participant-app/internal/entity"
)

type participantsController struct {
	participantsUsecase usecase.ParticipantsUsecase
}

type ParticipantReturnData struct {
	Id string `json: "id"`
	Name string `json: "name"`
}

type ParticipantsReturnData struct {
	B1 []ParticipantReturnData `json: "B1"`
	B2 []ParticipantReturnData `json: "B2"`
	B3 []ParticipantReturnData `json: "B3"`
	B4 []ParticipantReturnData `json: "B4"`
}

type ParticipantsController interface {
	GetParticipants(echo.Context) error
}


func NewParticipantsController(
	participantsUsecase usecase.ParticipantsUsecase,
) ParticipantsController {
	return &participantsController{
		participantsUsecase: participantsUsecase,
	}
}

func NewParticipantReturnData(
	participantEntity entity.ParticipantEntity,
) ParticipantReturnData {
	return ParticipantReturnData{
		Id: participantEntity.GetID(),
		Name: participantEntity.GetName(),
	}
}

func NewParticipantsReturnData(
	participantsEntity entity.ParticipantsEntity,
) ParticipantsReturnData {
	var participantsReturnData ParticipantsReturnData
	grades := []string{"B1", "B2", "B3", "B4"}

	for _, grade := range grades {
		switch grade {
		case "B1":
			for _, participantEntity := range participantsEntity.GetParticipants(grade) {
				participantsReturnData.B1 = append(participantsReturnData.B1, NewParticipantReturnData(participantEntity))
			}
		case "B2":
			for _, participantEntity := range participantsEntity.GetParticipants(grade) {
				participantsReturnData.B1 = append(participantsReturnData.B1, NewParticipantReturnData(participantEntity))
			}
		case "B3":
			for _, participantEntity := range participantsEntity.GetParticipants(grade) {
				participantsReturnData.B1 = append(participantsReturnData.B1, NewParticipantReturnData(participantEntity))
			}
		case "B4":
			for _, participantEntity := range participantsEntity.GetParticipants(grade) {
				participantsReturnData.B1 = append(participantsReturnData.B1, NewParticipantReturnData(participantEntity))
			}
		}
	}

	return participantsReturnData
}

func (participantsCtrl *participantsController) GetParticipants(c echo.Context) error {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil || year < 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "year is required or invalid")
	}


	month, err := strconv.Atoi(c.Param("month"))
	if err != nil || month < 1 || 12 < month {
		return echo.NewHTTPError(http.StatusBadRequest, "month is required or invalid")
	}

	date, err := strconv.Atoi(c.Param("date"))
	if err != nil || date < 1 || 31 < date{
		return echo.NewHTTPError(http.StatusBadRequest, "date is required or invalid")
	}

	participants, err := participantsCtrl.participantsUsecase.GetParticipants(year, month, date)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, NewParticipantsReturnData(participants))
}