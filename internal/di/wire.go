//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	// externalinterface
	"github.com/higuruchi/participant-app/internal/externalinterface/server"
	"github.com/higuruchi/participant-app/internal/externalinterface/database"
	// interfaceadapter
	"github.com/higuruchi/participant-app/internal/interfaceadapter/controller"
	interfaceadapterRepository "github.com/higuruchi/participant-app/internal/interfaceadapter/repository"
	"github.com/higuruchi/participant-app/internal/interfaceadapter/repository/worker"
	// usecase
	"github.com/higuruchi/participant-app/internal/usecase"
	usecaseRepository "github.com/higuruchi/participant-app/internal/usecase/repository"
)

func InitializeServer() (server.Server, func()) {
	panic(
		wire.Build(
			server.NewServer,
			controller.NewParticipantsController,
			controller.NewUserController,
			usecase.NewParticipantsUsecase,
			interfaceadapterRepository.NewUserRepository,
			usecase.NewUserUsecase,
			interfaceadapterRepository.NewParticipantsRepository,
			database.NewDBHandler,

			wire.Bind(
				new(usecaseRepository.ParticipantsRepository),
				new(*interfaceadapterRepository.ParticipantsRepository),
			),
			wire.Bind(
				new(worker.DatabaseHandler),
				new(*database.DatabaseHandler),
			),
			wire.Bind(
				new(usecaseRepository.UserRepository),
				new(*interfaceadapterRepository.UserRepository),
			),
		),
	)
}