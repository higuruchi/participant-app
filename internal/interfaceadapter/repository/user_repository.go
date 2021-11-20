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

func (userRepository *UserRepository)CreateUser(id string, name string, macaddress net.HardwareAddr) error {
	fmt.Println("fugahoge")
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