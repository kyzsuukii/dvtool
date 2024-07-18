package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadEnv() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
