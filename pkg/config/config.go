package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DB_HOST     string
	DB_NAME     string
	DB_USER     string
	DB_PORT     string
	DB_PASSWORD string
	BASE_URL    string
}

var envs = []string{"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "BASE_URL"}

func LoadConfig() (Config, error) {
	var config Config
	viper.AddConfigPath("/")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	fmt.Println("config ...", config)
	return config, nil
}
