package server

import(
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	// "github.com/go-playground/validator"
	// "github.com/labstack/echo/v4/middleware"
	"github.com/higuruchi/participant-app/internal/interfaceadapter/controller"
)

type (
	server struct {
		echoImplement *echo.Echo
		participantsCtrl controller.ParticipantsController
	}
	
	// customValidator struct {
	// 	validator *validator.Validate
	// }

	Server interface {
		Run() error
	}
)

func NewServer(participantsCtrl controller.ParticipantsController) Server {
	e := echo.New();
	// e.Validator = &CustomValidator{validator: validetor.New()}
	return &server{
		echoImplement: e,
		participantsCtrl: participantsCtrl,
	}
}

func (server *server) Run() error {
	server.echoImplement.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	server.echoImplement.GET("/participants/:year/:month/:date", server.participantsCtrl.GetParticipants)
	server.echoImplement.POST("/participants", server.participantsCtrl.SaveParticipants)

	err := server.echoImplement.Start(":1323")
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