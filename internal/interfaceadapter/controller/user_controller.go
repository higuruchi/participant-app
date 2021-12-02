package controller

import (
	"net"
	"regexp"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/higuruchi/participant-app/internal/usecase"
)

type userController struct {
	userUsecase usecase.UserUsecase
}

type UserController interface {
	CreateUser(c echo.Context) error
	UpdateUserMacaddr(c echo.Context) error
}

type ReturnData struct {
	Status bool `json: "status"`
}

func NewUserController(userUsecase usecase.UserUsecase) UserController {
	return &userController{
		userUsecase: userUsecase,
	}
}

func (userController *userController) CreateUser(c echo.Context) error {
	id := c.FormValue("id")
	match, err := regexp.MatchString("[1-9]{2}(T|G)[1-9]{3}", id); 
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Internal Server Error")
	}
	if !match {
		return echo.NewHTTPError(http.StatusBadRequest, "id is required or invalid")
	} 

	name := c.FormValue("name")
	match, err = regexp.MatchString(".{1, 20}", name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Internal Server Error")
	}
	if !match {
		return echo.NewHTTPError(http.StatusBadRequest, "name is required or invalid")
	}

	macaddress := c.FormValue("macaddress")
	hw, err := net.ParseMAC(macaddress)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "macaddress is required or invalid")
	}

	err = userController.userUsecase.CreateUser(id, name, hw)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, ReturnData{Status: true})
}

func (userController *userController) UpdateUserMacaddr(c echo.Context) error {
	id := c.FormValue("id")
	match, err := regexp.MatchString("[1-9]{2}(T|G)[1-9]{3}", id); 
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Internal Server Error")
	}
	if !match {
		return echo.NewHTTPError(http.StatusBadRequest, "id is required or invalid")
	}

	macaddress := c.FormValue("macaddress")
	hw, err := net.ParseMAC(macaddress)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "macaddress is required or invalid")
	}

	err  = userController.userUsecase.UpdateUserMacaddr(id, hw)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, ReturnData{Status: true})
}