package main

import (
	"github.com/higuruchi/participant-app/internal/di"
)

func main() {
	server := di.InitializeServerz()
	server.Run()
}