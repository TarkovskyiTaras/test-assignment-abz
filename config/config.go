package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Currencies    []string
	FetchInterval string
}

func ReadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	log.Printf("Using config: %s\n", viper.ConfigFileUsed())

	config := &Config{
		Currencies:    viper.GetStringSlice("currencies"),
		FetchInterval: viper.GetString("fetch_interval"),
	}

	return config, nil
}
