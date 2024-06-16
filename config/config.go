package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBSource string `mapstructure:"DB_SOURCE"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("[CONFIG] error read config :%v", err)
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Printf("[CONFIG] error unmarshall config file :%v", err)
		return nil, err
	}

	return &config, nil
}
