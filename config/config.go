package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Host	string	`mapstructure:"host"`
	Port 	int 	`mapstructure:"port"`
}

type DBConfig struct {
	Provider string `mapstructure:"provider"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type Config struct {
	App 	AppConfig	`mapstructure:"app"`
	DB    	DBConfig    `mapstructure:"db"`
}

func GetConfig() *Config {

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(err)
	}
	
	return config
}