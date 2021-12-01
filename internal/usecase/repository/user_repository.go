package repository

import (
	"net"
)

type UserRepository interface {
	CreateUser(string, string, net.HardwareAddr) error
	UpdateUserMacaddr(string, net.HardwareAddr) error
}