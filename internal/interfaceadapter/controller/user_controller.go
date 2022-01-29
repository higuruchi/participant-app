package controller

import (
	"net"
	"fmt"
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

type CreateUserData struct {
	ID string `json: "id"`
	Name string `json: "name"`
	Macaddress string `json: "macaddress"`
}

type UpdateMacaddrData struct {
	ID string `json: "id"`
	Macaddress string `json: "macaddress"`
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
	createUserData := new(CreateUserData)
	if err := c.Bind(createUserData); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Input data is inavlid")
	}

	match, err := regexp.MatchString("^[0-9]{2}(T|G)[0-9]{3}$", createUserData.ID); 
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Internal Server Error")
	}
	if !match {
		return echo.NewHTTPError(http.StatusBadRequest, "id is required or invalid")
	}

	match, err = regexp.MatchString(".{1,20}", createUserData.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Internal Server Error")
	}
	if !match {
		return echo.NewHTTPError(http.StatusBadRequest, "name is required or invalid")
	}

	hw, err := net.ParseMAC(createUserData.Macaddress)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "macaddress is required or invalid")
	}

	err = userController.userUsecase.CreateUser(createUserData.ID, createUserData.Name, hw)
	if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, ReturnData{Status: true})
}

func (userController *userController) UpdateUserMacaddr(c echo.Context) error {
	updateMacaddrData := new(UpdateMacaddrData)
	if err := c.Bind(updateMacaddrData); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Input data is inavlid")
	}

	match, err := regexp.MatchString("^[0-9]{2}(T|G)[0-9]{3}$", updateMacaddrData.ID); 
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Internal Server Error")
	}
	if !match {
		return echo.NewHTTPError(http.StatusBadRequest, "id is required or invalid")
	}

	hw, err := net.ParseMAC(updateMacaddrData.Macaddress)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "macaddress is required or invalid")
	}

	err  = userController.userUsecase.UpdateUserMacaddr(updateMacaddrData.ID, hw)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, ReturnData{Status: true})
}