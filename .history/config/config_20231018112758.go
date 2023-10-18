package config

import "github.com/spf13/viper"

type Config struct {
	PORT        string
	DB_USER     string
	DB_PASS     string
	DB_DATABASE string
	DB_HOST     string
	DB_PORT     string
}

var ENV Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := 
}