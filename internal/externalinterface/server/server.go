package server

import(
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
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
	server.echoImplement.PUT("/macaddr/:id", server.userCtrl.UpdateUserMacaddr)

	err := server.echoImplement.Start(fmt.Sprintf(":%d", server.port))
	if err != nil {
		return fmt.Errorf("calling Start: %w", err)
	}
	return nil
}