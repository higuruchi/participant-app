package usecase

import (
	"fmt"
	"net"
	"github.com/higuruchi/participant-app/internal/usecase/repository"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

type UserUsecase interface {
	CreateUser(string, string, net.HardwareAddr) error
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (userUsecase *userUsecase) CreateUser(
	id string,
	name string,
	macaddress net.HardwareAddr,
) error {
	if len(id) > 8 {
		return fmt.Errorf("invalid input data")
	}

	if len(id) > 20 {
		return fmt.Errorf("invalid input data")
	}

	err := userUsecase.userRepository.CreateUser(id, name, macaddress)
	if err != nil {
		return fmt.Errorf("calling userUsecase: %w", err)
	}

	return nil
}