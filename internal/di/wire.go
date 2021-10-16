//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/higuruchi/participant-app/internal/externalinterface/server"
)

func InitializeServerz() (server.Server) {
	panic(
		wire.Build(
			server.NewServer,
		),
	)
}