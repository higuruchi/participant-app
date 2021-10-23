package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/higuruchi/participant-app/internal/usecase"
)

type participantsController struct {
	participantsUsecase usecase.ParticipantsUsecase
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

	// 一時的に文字列型にして返却
	return c.String(http.StatusOK, fmt.Sprintf("%v", participants.GetParticipants("B3")[0]))
}