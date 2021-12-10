package config

import (
	"fmt"
	"net"
	"errors"
	"regexp"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `yaml: server`
	DB DBConfig `yaml: db`
}

type ServerConfig struct {
	Port int `yaml: port`
}

type DBConfig struct {
	User string `yaml: user`
	Password string `yaml: password`
	IP string `yaml: ip`
	Port int `yaml: port`
	Name string `yaml: name`
}

var (
	// サーバ用エラー
	ErrInvalidServerPort = errors.New("server port number is invalid")

	// データベース用エラー
	ErrInvalidDBIP = errors.New("database ip address is invalid")
	ErrInvalidDBUser = errors.New("database user name is invalid")
	ErrInvalidDBPassword = errors.New("database password is invalid")
	ErrInvalidDBPort = errors.New("database port number is invalid")
	ErrInvalidDBName = errors.New("database name is invalid")
)

func Load() (*Config, error) {
	viper.SetConfigName("participant-app-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("calling viper.ReadInConfig: %w", err)
	}

	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("calling viper.Unmarshal %w", err)
	}

	err = config.configValidation()
	if err != nil {
		return nil, fmt.Errorf("calling config.configValidation %w", err)
	}

	return &config, nil
}


func (config *Config)configValidation() error {
	if config.Server.Port < 1 || 65535 < config.Server.Port {
		return ErrInvalidServerPort
	}

	if ip := net.ParseIP(config.DB.IP); ip == nil {
		return ErrInvalidDBIP
	}

	match, err := regexp.MatchString(".+", config.DB.User)
	if err != nil {
		return fmt.Errorf("calling regexp.MatchString: %w", err)
	}
	if !match {
		return ErrInvalidDBUser
	}


	match, err = regexp.MatchString(".+", config.DB.Password)
	if err != nil {
		return fmt.Errorf("calling regexp.MatchString: %w", err)
	}
	if !match {
		return ErrInvalidDBPassword
	}

	if config.DB.Port < 1 || 65535 < config.DB.Port {
		return ErrInvalidDBPort
	}

	match, err = regexp.MatchString(".+", config.DB.Name)
	if err != nil {
		return fmt.Errorf("calling regexp.MatchString: %w", err)
	}
	if !match {
		return ErrInvalidDBName
	}

	return nil
}