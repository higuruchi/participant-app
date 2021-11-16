package server

import(
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/higuruchi/participant-app/internal/interfaceadapter/controller"
)

type server struct {
	echoImplement *echo.Echo
	participantsCtrl controller.ParticipantsController
}

type Server interface {
	Run() error
}

func NewServer(participantsCtrl controller.ParticipantsController) Server {
	return &server{
		echoImplement: echo.New(),
		participantsCtrl: participantsCtrl,
	}
}

func (server *server) Run() error {
	server.echoImplement.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	server.echoImplement.GET("/participants/:year/:month/:date", server.participantsCtrl.GetParticipants)

	err := server.echoImplement.Start(":1323")
	if err != nil {
		return fmt.Errorf("calling Start: %w", err)
	}
	return nil
}