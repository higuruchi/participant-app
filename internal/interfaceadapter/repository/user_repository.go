package repository

import(
	"fmt"
	"net"
	"github.com/higuruchi/participant-app/internal/interfaceadapter/repository/worker"
)

type UserRepository struct {
	databaseHandler worker.DatabaseHandler
}

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
	if len(id) <= 0 || 8 < len(id) {
		return fmt.Errorf("invalid input data")
	} 

	if len(name) <= 0 || 20 < len(name) {
		return fmt.Errorf("invalid input data")
	}

	sql := `
	INSERT INTO users
	(id, name, mac_address)
	VALUES
	(?, ?, ?)
	`

	_, err := userRepository.databaseHandler.Execute(sql, id, name, macaddress.String())
	if err != nil {
		return fmt.Errorf("calling userRepository.databaseHandler.Execute: %w", err)
	}

	return nil
}

func (userRepository *UserRepository) UpdateUserMacaddr(
	id string,
	macaddress net.HardwareAddr,
) error {
	if len(id) != 6 {
		return fmt.Errorf("invalid input data")
	}

	sql := `
	UPDATE users
	SET mac_address=?
	WHERE id=?
	`

	_, err := userRepository.databaseHandler.Execute(sql, macaddress.String(), id)
	if err != nil {
		return fmt.Errorf("calling userRepository.databaseHandler.Execute: %w", err)
	}

	return nil
}
