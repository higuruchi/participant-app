package main

import (
	"github.com/higuruchi/participant-app/internal/di"
)

func main() {
	server, f := di.InitializeServer()
	server.Run()
	f()
}