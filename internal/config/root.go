package config

import (
	"fmt"
	// "net"
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

	return &config, nil
}
