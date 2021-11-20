package repository

import (
	"net"
)

type UserRepository interface {
	CreateUser(string, string, net.HardwareAddr) error
}