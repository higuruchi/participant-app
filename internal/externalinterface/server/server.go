package server

import(
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)

type server struct {
	echoImplement *echo.Echo
}

type Server interface {
	Run() error
}

func NewServer() Server {
	return &server{
		echoImplement: echo.New(),
	}
}

func (server *server) Run() error {
	server.echoImplement.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	err := server.echoImplement.Start(":1323")
	if err != nil {
		return fmt.Errorf("calling Start: %w", err)
	}
	return nil
}