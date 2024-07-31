package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB struct {
		User     string
		Password string
		Host     string
		Port     string
		Name     string
	}
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
