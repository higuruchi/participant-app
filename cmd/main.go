package main

import (
	"fmt"
	"github.com/higuruchi/participant-app/internal/di"
)

func main() {
	fmt.Println("Start participant-app on 1323!")
	server, f := di.InitializeServer()
	server.Run()
	f()
}