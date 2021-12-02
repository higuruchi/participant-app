package repository

import(
	"fmt"
	"net"
	"errors"
	"regexp"
	"github.com/higuruchi/participant-app/internal/interfaceadapter/repository/worker"
)

type UserRepository struct {
	databaseHandler worker.DatabaseHandler
}

var (
	ErrInvalidInputData = errors.New("Invalid input data")
)

func NewUserRepository(databaseHandler worker.DatabaseHandler) *UserRepository {
	return &UserRepository{
		databaseHandler: databaseHandler,
	}
}

func (userRepository *UserRepository)CreateUser(
	id string,
	name string,
	macaddress net.HardwareAddr,
) error {
	match, err := regexp.MatchString("[1-9]{2}(T|G)[1-9]{3}", id); 
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

	sql := `
	INSERT INTO users
	(id, name, mac_address)
	VALUES
	(?, ?, ?)
	`

	_, err = userRepository.databaseHandler.Execute(sql, id, name, macaddress.String())
	if err != nil {
		return fmt.Errorf("calling userRepository.databaseHandler.Execute: %w", err)
	}

	return nil
}

func (userRepository *UserRepository) UpdateUserMacaddr(
	id string,
	macaddress net.HardwareAddr,
) error {
	match, err := regexp.MatchString("[1-9]{2}(T|G)[1-9]{3}", id); 
	if err != nil {
		return fmt.Errorf("calling regexp.MatchString: %w", err)
	}
	if !match {
		return ErrInvalidInputData
	}

	sql := `
	UPDATE users
	SET mac_address=?
	WHERE id=?
	`

	_, err = userRepository.databaseHandler.Execute(sql, macaddress.String(), id)
	if err != nil {
		return fmt.Errorf("calling userRepository.databaseHandler.Execute: %w", err)
	}

	return nil
}
