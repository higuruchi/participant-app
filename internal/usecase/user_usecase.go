package usecase

import (
	"fmt"
	"net"
	"regexp"
	"errors"
	"github.com/higuruchi/participant-app/internal/usecase/repository"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

var (
	ErrInvalidInputData = errors.New("Invalid input data")
)

type UserUsecase interface {
	CreateUser(string, string, net.HardwareAddr) error
	UpdateUserMacaddr(string, net.HardwareAddr) error
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
	match, err := regexp.MatchString("^[0-9]{2}(T|G)[0-9]{3}$", id); 
	if err != nil {
		return fmt.Errorf("calling regexp.MatchString: %w", err)
	}
	if !match {
		return ErrInvalidInputData
	}

	match, err = regexp.MatchString(".{1,20}", name)
	if err != nil {
		return fmt.Errorf("calling regexp.MatchString: %w", err)
	}
	if !match {
		return ErrInvalidInputData
	}

	err = userUsecase.userRepository.CreateUser(id, name, macaddress)
	if err != nil {
		return fmt.Errorf("calling userUsecase: %w", err)
	}

	return nil
}

func (userUsecase *userUsecase) UpdateUserMacaddr(
	id string,
	macaddress net.HardwareAddr,
) error {
	match, err := regexp.MatchString("^[0-9]{2}(T|G)[0-9]{3}$", id); 
	if err != nil {
		return fmt.Errorf("calling regexp.MatchString: %w", err)
	}
	if !match {
		return ErrInvalidInputData
	}

	err = userUsecase.userRepository.UpdateUserMacaddr(id, macaddress)
	if err != nil {
		return fmt.Errorf("calling userUsecase.userRepository.UpdateUserMacaddr: %w", err)
	}

	return nil
}