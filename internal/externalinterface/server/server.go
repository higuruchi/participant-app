package server

import(
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	// "github.com/go-playground/validator"
	// "github.com/labstack/echo/v4/middleware"
	"github.com/higuruchi/participant-app/internal/config"
	"github.com/higuruchi/participant-app/internal/interfaceadapter/controller"
)

type (
	server struct {
		port int
		echoImplement *echo.Echo
		participantsCtrl controller.ParticipantsController
		userCtrl controller.UserController
	}
	
	// customValidator struct {
	// 	validator *validator.Validate
	// }

	Server interface {
		Run() error
	}
)

func NewServer(
	participantsCtrl controller.ParticipantsController,
	userCtrl controller.UserController,
	config *config.Config,
) Server {
	e := echo.New();
	// e.Validator = &CustomValidator{validator: validetor.New()}
	return &server{
		port: config.Server.Port,
		echoImplement: e,
		participantsCtrl: participantsCtrl,
		userCtrl: userCtrl,
	}
}

func (server *server) Run() error {
	server.echoImplement.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	server.echoImplement.GET("/participants/:year/:month/:date", server.participantsCtrl.GetParticipants)
	server.echoImplement.POST("/participants", server.participantsCtrl.SaveParticipants)
	server.echoImplement.POST("/user", server.userCtrl.CreateUser)

	err := server.echoImplement.Start(fmt.Sprintf(":%d", server.port))
	if err != nil {
		return fmt.Errorf("calling Start: %w", err)
	}
	return nil
}

// func (cv *CustomValidator) Validate(i interface{}) error {
// 	if err != cv.validator.Struct(i); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return nil
// }