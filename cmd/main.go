package main

import (
	"fmt"
	"log"
	"github.com/higuruchi/participant-app/internal/di"
	"github.com/higuruchi/participant-app/internal/config"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalf("calling cmd.Execute: %v", err)
	}

	fmt.Printf("Start participant-app on %d!", config.Server.Port)
	server, f := di.InitializeServer(config)
	server.Run()
	f()
}